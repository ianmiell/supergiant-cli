package apictl

import (
	"errors"

	"github.com/supergiant/supergiant/client"
	"github.com/supergiant/supergiant/common"
)

// CreateContainer creates a new Container for a resource.
func CreateContainer(r *client.ReleaseResource, containerName string, image string, cmax string, cmin string, mmax string, mmin string) error {

	Container := &common.ContainerBlueprint{
		Name:  containerName,
		Image: image,
		CPU: &common.CpuAllocation{
			Max: common.CoresFromString(cmax),
			Min: common.CoresFromString(cmin),
		},
		RAM: &common.RamAllocation{
			Max: common.BytesFromString(mmax),
			Min: common.BytesFromString(mmin),
		},
	}

	r.Containers = append(r.Containers, Container)

	_, err := r.Save()
	if err != nil {
		return err
	}

	if helloWorldExist(r) {
		DeleteContainer(r, "Hello World")
	}

	return nil
}

// UpdateContainer updates a Container for a resource.
func UpdateContainer(r *client.ReleaseResource, containerName string, image string, cmax string, cmin string, mmax string, mmin string) error {
	/// Find our Container.
	if len(r.Containers) == 0 {
		return errors.New("This Component has not Containers.")
	}
	for i, Container := range r.Containers {
		if Container.Name == containerName {
			// Edit the Container.
			if image != "" {
				r.Containers[i].Image = image
			}
			if cmax != "" {
				r.Containers[i].CPU.Max = common.CoresFromString(cmax)
			}
			if cmin != "" {
				r.Containers[i].CPU.Min = common.CoresFromString(cmin)
			}
			if mmax != "" {
				r.Containers[i].RAM.Max = common.BytesFromString(mmax)
			}
			if mmin != "" {
				r.Containers[i].RAM.Min = common.BytesFromString(mmin)
			}
		} else {
			return errors.New("Container not found.")
		}
	}

	_, err := r.Save()
	if err != nil {
		return err
	}

	return nil
}

// DeleteContainer creates a new Container for a resource.
func DeleteContainer(r *client.ReleaseResource, name string) error {
	success := false
	if len(r.Containers) == 0 {
		return errors.New("This Component has not Containers.")
	}

	for i, container := range r.Containers {
		if container.Name == name {
			r.Containers = append(r.Containers[:i], r.Containers[i+1:]...)
			success = true
		}
	}

	if !success {
		return errors.New("Container not found.")
	}

	_, err := r.Save()
	if err != nil {
		return err
	}

	return nil
}

func helloWorldExist(r *client.ReleaseResource) bool {
	if len(r.Containers) == 0 {
		return false
	}

	for _, container := range r.Containers {
		if container.Name == "Hello World" {
			return true
		}
	}
	return false
}
