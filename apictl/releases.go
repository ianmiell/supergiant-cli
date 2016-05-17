package apictl

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
	"github.com/supergiant/supergiant/client"
	"github.com/supergiant/supergiant/common"
)

var (
	releaseBoiler = &client.Release{
		InstanceCount: 1,
		Volumes:       []*common.VolumeBlueprint{},
		Meta: &common.Meta{
			Tags: map[string]string{},
		},
		Containers: []*common.ContainerBlueprint{
			&common.ContainerBlueprint{
				Name:  "Hello World",
				Image: "hello-world",
				CPU: &common.CpuAllocation{
					Max: common.CoresFromString("0"),
					Min: common.CoresFromString("0"),
				},
				RAM: &common.RamAllocation{
					Max: common.BytesFromString("0"),
					Min: common.BytesFromString("0"),
				},
				Mounts:  []*common.Mount{},
				Ports:   []*common.Port{},
				Env:     []*common.EnvVar{},
				Command: []string{},
			},
		},
	}
)

func getReleaseFromFile(thefile string) (*client.Release, error) {
	// Make sure release file exists.
	_, err := os.Stat(thefile)
	// else create it.
	if os.IsNotExist(err) {
		return nil, errors.New("Unable to locate file...")
	}

	// Read data from file.
	file, err := ioutil.ReadFile(thefile)
	if err != nil {
		return nil, err
	}

	release := new(client.Release)
	//Lets try to read json first.
	jsonerr := json.Unmarshal(file, &release)
	if jsonerr != nil {
		// if that fails... lets try yaml
		yamlerr := yaml.Unmarshal(file, &release)
		if yamlerr != nil {
			return nil, errors.New("Unable to determine file format. Please ensure your file is properly formated in json, or yaml format.")
		}
	}
	return release, nil
}

//GetReleaseExample will return examples of how to create a import file to the user.
func GetReleaseExample() error {
	data, err := json.MarshalIndent(releaseBoiler, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println("Example JSON File:")
	fmt.Println(string(data))

	fmt.Println("Example YAML File:")
	yaml, err := yaml.JSONToYAML(data)
	if err != nil {
		return err
	}
	fmt.Println(string(yaml))
	return nil
}

// GetRelease gets the target release, or a boiler plate with context
func (sg *SGClient) GetRelease(appName string, compName string) (*client.ReleaseResource, error) {

	comp, err := sg.GetComponent(compName, appName)
	if err != nil {
		return nil, compGetError(compName, err)
	}

	release, err := comp.TargetRelease()
	if err != nil {
		release, err = comp.CurrentRelease()
		if err != nil {
			return nil, errors.New("This Component has no releases.")
		}
	}
	return release, nil
}
