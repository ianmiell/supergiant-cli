package sgcore

import (
	"fmt"
	"strings"

	"github.com/supergiant/guber"
	"github.com/supergiant/supergiant-cli/apictl"
	"github.com/supergiant/supergiant-cli/spacetime"
)

//This may be stupid, but for now it will aid with dubugs.

func initDash(c guber.Client, version string, k *spacetime.Kube, update bool) (string, error) {
	sg, err := apictl.NewClient("", "", "")
	if err != nil {
		return "", err
	}

	if version == "" {
		version = ":latest"
	} else {
		version = ":" + version
	}
	fmt.Println("Installing Dashboard version", version)
	sg.CreateApp("supergiant")

	if !update {
		err = sg.CreateEntryPoint(k.Name)
		if err != nil {
			fmt.Println("WARN ENTRY POINT:", err)
		}
	}

	err = sg.CreateComponent("sg-ui", "supergiant", "")
	if err != nil {
		fmt.Println("WARN COMP:", err)
	}

	release, err := sg.GetRelease("supergiant", "sg-ui")
	if err != nil {
		return "", err
	}

	err = apictl.CreateContainer(
		release,
		"dashboard",
		"supergiant/supergiant-dashboard"+version,
		"0",
		"0",
		"0",
		"0",
	)
	if err != nil {
		return "", err
	}

	err = apictl.CreatePort(
		release,
		"dashboard",
		"TCP",
		9001,
		true,
		k.Name,
		80,
	)
	if err != nil {
		return "", err
	}

	err = apictl.CreateEnv(
		release,
		"dashboard",
		"SG_API_HOST",
		"http://supergiant-api",
	)
	if err != nil {
		return "", err
	}

	err = apictl.CreateEnv(
		release,
		"dashboard",
		"SG_API_PORT",
		"80",
	)
	if err != nil {
		return "", err
	}

	err = sg.DeployComponent("sg-ui", "supergiant")
	if err != nil {
		return "", err
	}

	dash, err := sg.GetEntryURL(k.Name)
	if err != nil {
		return "", err
	}
	s := strings.Split(dash, ":")
	return s[0], nil

}

func destroyDash(k *spacetime.Kube, update bool) error {
	sg, err := apictl.NewClient("", "", "")
	if err != nil {
		return err
	}
	err = sg.DestroyComponent("sg-ui", "supergiant")
	if err != nil {
		return err
	}

	if !update {
		err = sg.DestroyEntryPoint(k.Name)
		if err != nil {
			return nil
		}
	}
	return nil
}
