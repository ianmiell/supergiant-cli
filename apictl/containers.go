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
			if cmax != "" && cmin != "" {
				r.Containers[i].CPU = &common.CpuAllocation{
					Max: common.CoresFromString(cmax),
					Min: common.CoresFromString(cmin),
				}
			} else if cmax != "" || cmin != "" {
				return errors.New("Both CPU MAX, and MIN must be defined.")
			}
			if mmax != "" && mmin != "" {
				r.Containers[i].RAM = &common.RamAllocation{
					Max: common.BytesFromString(mmax),
					Min: common.BytesFromString(mmin),
				}
			} else if mmax != "" || mmin != "" {
				return errors.New("Both RAM MAX, and MIN must be defined.")
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
