package apictl

import (
	"errors"

	"github.com/supergiant/supergiant/client"
)

// CreateCommand creates a new command for a resource.
func CreateCommand(r *client.ReleaseResource, containerName string, commands []string) error {
	for _, container := range r.Containers {
		if container.Name == containerName {
			for _, command := range commands {
				container.Command = append(container.Command, command)
			}
		}
	}

	_, err := r.Save()
	if err != nil {
		return err
	}

	return nil
}

// DeleteCommand deletes a Command for a container resource.
func DeleteCommand(r *client.ReleaseResource, containerName string, commands []string) error {

	for ci, container := range r.Containers {
		if container.Name == containerName {
			if len(container.Command) == 0 {
				return errors.New("This container has no Commands.")
			}
			for cmi, ecommand := range container.Command {
				for _, command := range commands {
					if ecommand == command {
						r.Containers[ci].Command = append(r.Containers[ci].Command[:cmi], r.Containers[ci].Command[cmi+1:]...)
					}
				}
			}
		}
	}

	_, err := r.Save()
	if err != nil {
		return err
	}

	return nil
}
