package apictl

import (
	"errors"
	"strings"

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

// GetRelease gets the target release, or a boiler plate with context
func GetRelease(appName string, compName string) (*client.ReleaseResource, error) {

	comp, err := GetComponent(compName, appName)
	if err != nil {
		return nil, compGetError(compName, err)
	}

	release, err := comp.TargetRelease()
	if strings.Contains(err.Error(), "no TargetReleaseTimestamp") {
		release, err = comp.CurrentRelease()
	}
	if err != nil {
		return nil, errors.New("This Component has no releases.")
	}
	return release, nil
}
