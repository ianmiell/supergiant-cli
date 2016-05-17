package apictl

import (
	"errors"
	"os"

	"github.com/supergiant/supergiant-cli/spacetime"
	"github.com/supergiant/supergiant/client"
)

//SGClient holds a supergiant client.
type SGClient struct {
	*client.Client
}

//NewClient gets a new supergiant api client.
func NewClient(ip string, user string, pass string) (*SGClient, error) {
	if ip == "" {
		kube, err := spacetime.GetDefaultKube()
		if err != nil {
			return nil, errors.New("No spacetime selected. Use the select action to select the Kubernetes cluster you wish to use.")
		}
		ip = kube.IP
		user = kube.User
		pass = kube.Pass
	}

	c := client.New("https://"+ip+"/api/v1/proxy/namespaces/supergiant/services/supergiant-api:frontend/v0",
		user,
		pass,
		true,
	)

	debug := os.Getenv("SG_CLI_DEBUG")

	if debug == "true" {
		client.Log.SetLevel("debug")
	}
	return &SGClient{c}, nil
}
