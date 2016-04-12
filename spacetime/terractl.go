package spacetime

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/google/go-github/github"
	"github.com/mitchellh/go-homedir"
)

// Grab assets from git.
func terraformINIT(version string, root string) error {
	client := github.NewClient(nil)
	orgs, _, err := client.Repositories.GetLatestRelease("supergiant", "terraform-assets")
	if err != nil {
		return err
	}

	err = download(*orgs.TarballURL, root)

	err = ungzip(""+root+"/terra.tar.gz", ""+root+"/terra.tar")
	if err != nil {
		return err
	}

	err = untar(""+root+"/terra.tar", root, "terraform-assets")
	if err != nil {
		return err
	}

	err = os.Remove("" + root + "/pax_global_header")
	if err != nil {
		return err
	}

	return nil
}

// Verifies a provider is good to go by creating some enviromental requirments for supergiant.
func (provider *ProviderConfig) init(verbose bool) error {
	home, err := homedir.Dir()
	if err != nil {
		return err
	}

	// Setup a state dir.
	terraStateDir := "" + home + "/.supergiant/terraform-states/" + provider.Name + ""
	sgroot := "" + home + "/.supergiant"

	// Make sure terra assets are in place.
	err = terraformINIT("latest", sgroot)
	if err != nil {
		return err
	}

	// Make sure the state dir is created.
	err = os.MkdirAll(terraStateDir, 0700)
	if err != nil {
		return err
	}

	// Copy the assets to an isolated state dir.
	err = copyDir(""+sgroot+"/terraform-assets/"+provider.Type+"/provider", terraStateDir)
	if err != nil {
		return err
	}

	// Build the command that will setup the env.
	cmd := exec.Command(
		"terraform",
		"apply",
		"-var", "aws_access_key="+provider.AccessKey+"",
		"-var", "aws_secret_key="+provider.SecretKey+"",
		"-var", "region=us-east-1",
	)

	// Log file
	outfile, err := os.Create("" + terraStateDir + "/supergiant.log")
	if err != nil {
		return err
	}
	provider.Log = "" + terraStateDir + "/supergiant.log"
	provider.update()

	logger := bufio.NewWriter(outfile)
	defer logger.Flush()

	cmd.Dir = terraStateDir
	stdOutPipe, _ := cmd.StdoutPipe()
	stdErrPipe, _ := cmd.StderrPipe()

	//	if verbose {
	//		cmd.Stdout = os.Stdout
	//		cmd.Stderr = os.Stderr
	//	}

	err = cmd.Start()
	if err != nil {
		return err
	}

	scanout := bufio.NewScanner(stdOutPipe)
	scanerr := bufio.NewScanner(stdErrPipe)

	//	go io.Copy(logger, stdOutPipe)
	//	go io.Copy(logger, stdErrPipe)

	// I am going it this way instead of sending directoy to STDOUT so I have the abioity to parse output.

	for scanout.Scan() {
		fmt.Fprintln(logger, scanout.Text())
		if verbose {
			fmt.Println(scanout.Text())
		}
	}

	var okFaultCounter int
	for scanerr.Scan() {
		fmt.Fprintln(logger, scanerr.Text())
		if verbose {
			fmt.Println(scanerr.Text())
		}
		result := strings.Contains(scanerr.Text(), "EntityAlreadyExists")
		if result {
			okFaultCounter++
		}
	}

	err = cmd.Wait()
	if err != nil {
		if okFaultCounter > 0 {
			provider.Status = "Verified"
			provider.update()
			if verbose {
				fmt.Println("Resource already exists...")
			}
			return nil
		}
		return err
	}

	provider.Status = "Verified"
	provider.update()
	if verbose {
		fmt.Println("Terraform build completed successfully...")
	}

	return nil
}

func (kube *Kube) init(verbose bool) error {
	home, err := homedir.Dir()
	if err != nil {
		kube.fail()
		return err
	}
	terraStateDir := "" + home + "/.supergiant/terraform-states/" + kube.Name + ""
	sgroot := "" + home + "/.supergiant"

	err = terraformINIT("latest", sgroot)
	if err != nil {
		kube.fail()
		return err
	}

	err = os.MkdirAll(terraStateDir, 0700)
	if err != nil {
		kube.fail()
		return err
	}

	// Fetch the current config
	providerConfigs, err := loadProviderConfig()
	if err != nil {
		kube.fail()
		return err
	}

	provider, err := providerConfigs.get(kube.Provider)
	if err != nil {
		kube.fail()
		return err
	}

	err = copyDir(""+sgroot+"/terraform-assets/"+provider.Type+"/"+kube.KubeVersion, terraStateDir)
	if err != nil {
		kube.fail()
		return err
	}
	fmt.Println("5")
	cmd := exec.Command(
		"terraform",
		"apply",
		"-var", "aws_access_key="+provider.AccessKey+"",
		"-var", "aws_secret_key="+provider.SecretKey+"",
		"-var", "cluster_id="+kube.Name+"",
		"-var", "region="+kube.Region+"",
		"-var", "availability_zone="+kube.AZ+"",
		"-var", "kube_user="+kube.User+"",
		"-var", "kube_pass="+kube.Pass+"",
	)

	outfile, err := os.Create("" + terraStateDir + "/supergiant.log")
	if err != nil {
		kube.fail()
		return err
	}
	kube.Log = "" + terraStateDir + "/supergiant.log"
	kube.Update()

	logger := bufio.NewWriter(outfile)
	defer logger.Flush()

	cmd.Dir = terraStateDir
	stdOutPipe, _ := cmd.StdoutPipe()
	stdErrPipe, _ := cmd.StderrPipe()

	err = cmd.Start()
	if err != nil && verbose {
		return err
	}

	scanout := bufio.NewScanner(stdOutPipe)
	scanerr := bufio.NewScanner(stdErrPipe)

	// I am going it this way instead of sending directoy to STDOUT so I have the abioity to parse output.
	for scanout.Scan() {
		fmt.Fprintln(logger, scanout.Text())
		if verbose {
			fmt.Println(scanout.Text())
		}
	}

	for scanerr.Scan() {
		fmt.Fprintln(logger, scanerr.Text())
		if verbose {
			fmt.Println(scanerr.Text())
		}
	}

	err = cmd.Wait()
	if err != nil && verbose {
		kube.fail()
		return err
	}

	if verbose {
		fmt.Println("Terraform build completed successfully...")
		// Fetch info from our new cluster
		fmt.Println("Loading cluster informations from terraform...")
		kube.IP, err = terraStateParser("" + terraStateDir + "/terraform.tfstate")
	}
	if err != nil {
		kube.fail()
		return err
	}
	kube.Update()

	// tesing that we are able to reach the cluster with guber.

	for i := 0; i < 5; i++ {
		if verbose {
			fmt.Println("Check for life poll...", i)
		}
		err = checkForLife(kube.IP, kube.User, kube.Pass)
		if err != nil {
			time.Sleep(200 * time.Second)
		} else {
			kube.fail()
			break
		}
	}
	if err != nil {
		kube.fail()
		return err
	}
	kube.Status = "Active"
	kube.Update()
	return nil
}
func (kube *Kube) destroy(verbose bool) error {

	home, err := homedir.Dir()
	if err != nil {
		kube.fail()
		return err
	}
	terraStateDir := "" + home + "/.supergiant/terraform-states/" + kube.Name + ""

	_, err = os.Stat(terraStateDir)
	if os.IsNotExist(err) {
		return nil
	}
	// Fetch the current config
	providerConfigs, err := loadProviderConfig()
	if err != nil {
		kube.fail()
		return err
	}
	provider, err := providerConfigs.get(kube.Provider)
	if err != nil {
		kube.fail()
		return err
	}

	cmd := exec.Command(
		"terraform",
		"destroy",
		"-force",
		"-var", "aws_access_key="+provider.AccessKey+"",
		"-var", "aws_secret_key="+provider.SecretKey+"",
		"-var", "cluster_id="+kube.Name+"",
		"-var", "region="+kube.Region+"",
		"-var", "availability_zone="+kube.AZ+"",
		"-var", "kube_user="+kube.User+"",
		"-var", "kube_pass="+kube.Pass+"",
	)

	outfile, err := os.Create("" + terraStateDir + "/supergiant.log")
	if err != nil {
		return err
	}

	logger := bufio.NewWriter(outfile)
	defer logger.Flush()

	cmd.Dir = terraStateDir
	stdOutPipe, _ := cmd.StdoutPipe()
	stdErrPipe, _ := cmd.StderrPipe()

	// TODO: We should parse this writer/stdout/stderr and check for basic repairable errors like non existant AZ
	if verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	err = cmd.Start()
	if err != nil {
		return err
	}

	go io.Copy(logger, stdOutPipe)
	go io.Copy(logger, stdErrPipe)
	err = cmd.Wait()
	if err != nil {
		return err
	}

	err = os.RemoveAll(terraStateDir)
	if err != nil {
		return err
	}
	return nil
}
