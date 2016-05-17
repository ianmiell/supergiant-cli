package apictl

import (
	"errors"
	"fmt"

	"github.com/supergiant/supergiant/client"
	"github.com/supergiant/supergiant/common"
)

// CreateMount creates a new volume for a resource.
func CreateMount(r *client.ReleaseResource, containerName string, path string, volume string) error {

	if !volumeExist(r, volume) {
		return errors.New("Volume Does Not Exist...")
	}

	mount := &common.Mount{
		Path:   path,
		Volume: &volume,
	}

	fault := false
	for _, container := range r.Containers {
		if container.Name == containerName {
			container.Mounts = append(container.Mounts, mount)
			fault = false
			break
		}
		fault = true
	}

	// error if no containers found.
	if fault {
		return errors.New("Container Does Not Exist...")
	}

	_, err := r.Save()
	if err != nil {
		return err
	}

	return nil
}

// UpdateMount updates a mount for a container resource.
func UpdateMount(r *client.ReleaseResource, containerName string, path string, volume string) error {

	if !volumeExist(r, volume) {
		return errors.New("Volume Does Not Exist...")
	}

	fault := false
	for ci, container := range r.Containers {
		fmt.Println(container.Name)
		if container.Name == containerName {
			if len(container.Mounts) == 0 {
				return errors.New("This container has no mounts.")
			}
			for mi, mount := range container.Mounts {
				fmt.Println(mount.Path)
				if mount.Path == path {
					if volume != "" {
						r.Containers[ci].Mounts[mi].Volume = &volume
					}
				} else {
					return errors.New("Mount not found.")
				}
			}
			fault = false
			break
		}
		fault = true
	}

	// error if no containers found.
	if fault {
		return errors.New("Container Does Not Exist...")
	}

	_, err := r.Save()
	if err != nil {
		return err
	}

	return nil
}

// DeleteMount deletes a mount for a container resource.
func DeleteMount(r *client.ReleaseResource, containerName string, path string) error {

	fault := false
	for ci, container := range r.Containers {
		if container.Name == containerName {
			if len(container.Mounts) == 0 {
				return errors.New("This container has no mounts.")
			}
			for mi, mount := range container.Mounts {
				if mount.Path == path {
					r.Containers[ci].Mounts = append(r.Containers[ci].Mounts[:mi], r.Containers[ci].Mounts[mi+1:]...)
				} else {
					return errors.New("Mount not found.")
				}
			}
			fault = false
			break
		}
		fault = true
	}

	// error if no containers found.
	if fault {
		return errors.New("Container Does Not Exist...")
	}

	_, err := r.Save()
	if err != nil {
		return err
	}

	return nil
}
