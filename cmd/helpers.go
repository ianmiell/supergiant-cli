package cmd

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/codegangsta/cli"
	"github.com/google/go-github/github"
	"github.com/supergiant/supergiant-cli/spacetime"
)

// Hack to fill in an required flags from the user.
func required(c *cli.Context, s string, msg string) string {
	var resp string
	if c.String(s) == "" {
		fmt.Print("Please enter " + msg + ": ")
		fmt.Scanln(&resp)
		if resp == "" {
			fmt.Println("ERROR: Empty response.")
			os.Exit(7)
		}
		return resp
	}
	return c.String(s)
}

func context(contextType string) (string, error) {
	switch contextType {
	case "kube":
		kube, err := spacetime.GetDefaultKube()
		if err != nil {
			return "", err
		}
		return kube.Name, nil

	}
	return "", errors.New("Context Not Set.")
}

func getApp(c *cli.Context) string {
	var appName string
	args := c.Args().Tail()
	for i, arg := range args {
		if arg == "--app" {
			appName = args[i+1]
		}
	}
	if appName == "" {
		appName = required(c, "app", "Application Name")
	}
	return appName
}
func getComp(c *cli.Context) string {
	var comp string
	args := c.Args().Tail()
	for i, arg := range args {
		if arg == "--comp" {
			comp = args[i+1]
		}
	}
	if comp == "" {
		comp = required(c, "comp", "Component Name")
	}
	return comp
}
func getFile(c *cli.Context) string {
	var fileName string
	args := c.Args().Tail()
	for i, arg := range args {
		if arg == "--file" {
			fileName = args[i+1]
		}
	}
	return fileName
}

func updateCLI(version string) error {
	out, err := exec.Command("supergiant", "--version").Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(5)
	}
	outs := strings.Fields(string(out))
	currentVersion := outs[2]
	fmt.Printf("Current CLI version is: %s\n", currentVersion)

	client := github.NewClient(nil)
	releases, _, err := client.Repositories.GetLatestRelease("supergiant", "supergiant-cli")
	//	releases, _, err := client.Repositories.ListReleases("supergiant", "supergiant-cli", nil)
	if err != nil {
		return err
	}
	latest := *releases.TagName
	latestVersion := strings.Trim(latest, "v")
	fmt.Println("Latest CLI version is:", latestVersion)

	if currentVersion != latestVersion {
		fmt.Println("An update is available!")
		fmt.Println("Searching for latest " + runtime.GOOS + " binary...")

		for _, release := range releases.Assets {
			name := *release.Name
			if strings.Contains(name, runtime.GOOS) && strings.Contains(name, runtime.GOARCH) {
				url := *release.BrowserDownloadURL
				fmt.Println("Downloading release:", url)
				s := spinner.New(spinner.CharSets[1], 100*time.Millisecond)
				s.Start()
				err := downloadBin(url, "/tmp")
				if err != nil {
					return err
				}
				s.Stop()

				err = os.Rename("/tmp/supergiant", "/usr/local/bin/supergiant")
				if err != nil {
					return errors.New("Install of supergiant bin to : /usr/local/bin/ , has failed. Try running update with root level access.")
				}

				return nil
			}
		}
	}
	fmt.Println("Your supergiant cli is up to date...")
	return nil
}

func downloadBin(url string, root string) error {
	out, err := os.Create("" + root + "/supergiant")
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	err = os.Chmod(""+root+"/supergiant", 0755)
	if err != nil {
		return err
	}

	return nil
}
