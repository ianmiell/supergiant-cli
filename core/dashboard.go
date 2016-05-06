package sgcore

import (
	"fmt"

	"github.com/supergiant/guber"
	"github.com/supergiant/supergiant-cli/apictl"
)

//This may be stupid, but for now it will aid with dubugs.

func initDash(c guber.Client, version string) (string, error) {

	if version == "" {
		version = ":latest"
	} else {
		version = ":" + version
	}
	fmt.Println("Installing Dashboard version", version)
	apictl.CreateApp("supergiant")
	err := apictl.CreateEntryPoint("supergiant")
	if err != nil {
		fmt.Println("WARN ENTRY POINT:", err)
	}

	err = apictl.CreateComponent("sg-ui", "supergiant", "")
	if err != nil {
		fmt.Println("WARN COMP:", err)
	}

	release, err := apictl.GetRelease("supergiant", "sg-ui")
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
		"supergiant",
		0,
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

	err = apictl.DeployComponent("sg-ui", "supergiant")
	if err != nil {
		return "", err
	}

	dash, err := apictl.GetEntryURL("supergiant")
	if err != nil {
		return "", err
	}
	return dash, nil

}

func destroyDash() error {
	err := apictl.DestroyComponent("sg-ui", "supergiant")
	if err != nil {
		return err
	}
	err = apictl.DestroyEntryPoint("supergiant")
	if err != nil {
		return nil
	}
	return nil
}
