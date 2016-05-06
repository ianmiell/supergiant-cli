package apictl

import (
	"errors"
	"os"

	"github.com/supergiant/supergiant-cli/spacetime"
	"github.com/supergiant/supergiant/client"
)

func getClient() (*client.Client, error) {
	kube, err := spacetime.GetDefaultKube()
	if err != nil {
		return nil, errors.New("No spacetime selected. Use the select action to select the Kubernetes cluster you wish to use.")
	}
	c := client.New("https://"+kube.IP+"/api/v1/proxy/namespaces/supergiant/services/supergiant-api:frontend/v0", kube.User, kube.Pass, true)
	//c := client.New("http://localhost:8080/v0", "", "", false)

	debug := os.Getenv("SG_CLI_DEBUG")

	if debug == "true" {
		client.Log.SetLevel("debug")
	}
	return c, nil
}
