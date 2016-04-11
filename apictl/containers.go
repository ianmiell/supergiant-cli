package apictl

import (
	"errors"

	"github.com/supergiant/supergiant/client"
	"github.com/supergiant/supergiant/common"
)

// CreateContainer creates a new Container for a resource.
func CreateContainer(r *client.ReleaseResource, containerName string, image string, cmax uint, cmin uint, mmax uint, mmin uint) error {

	Container := &common.ContainerBlueprint{
		Name:  containerName,
		Image: image,
		CPU: &common.ResourceAllocation{
			Max: cmax,
			Min: cmin,
		},
		RAM: &common.ResourceAllocation{
			Max: mmax,
			Min: mmin,
		},
	}

	r.Containers = append(r.Containers, Container)

	_, err := r.Save()
	if err != nil {
		return err
	}

	return nil
}

// UpdateContainer updates a Container for a resource.
func UpdateContainer(r *client.ReleaseResource, containerName string, image string, cmax uint, cmin uint, mmax uint, mmin uint) error {
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
			if cmax != 0 {
				r.Containers[i].CPU.Max = cmax
			}
			if cmin != 0 {
				r.Containers[i].CPU.Min = cmin
			}
			if mmax != 0 {
				r.Containers[i].RAM.Max = mmax
			}
			if mmin != 0 {
				r.Containers[i].RAM.Min = mmin
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

	if len(r.Containers) == 0 {
		return errors.New("This Component has not Containers.")
	}
	for i, container := range r.Containers {
		if container.Name == name {
			r.Containers = append(r.Containers[:i], r.Containers[i+1:]...)
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
