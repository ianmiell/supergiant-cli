package apictl

import (
	"errors"

	"github.com/supergiant/supergiant/client"
	"github.com/supergiant/supergiant/common"
)

// CreateEnv creates a new volume for a resource.
func CreateEnv(r *client.ReleaseResource, containerName string, name string, value string) error {
	env := &common.EnvVar{
		Name:  name,
		Value: value,
	}

	for _, container := range r.Containers {
		if container.Name == containerName {
			container.Env = append(container.Env, env)
		}
	}

	_, err := r.Save()
	if err != nil {
		return err
	}

	return nil
}

// UpdateEnv updates a Env for a container resource.
func UpdateEnv(r *client.ReleaseResource, containerName string, name string, value string) error {

	for ci, container := range r.Containers {
		if container.Name == containerName {
			if len(container.Env) == 0 {
				return errors.New("This container has no Envs.")
			}
			for ei, env := range container.Env {
				if env.Name == name {
					r.Containers[ci].Env[ei].Value = value
				} else {
					return errors.New("Env not found.")
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

// DeleteEnv deletes a Env for a container resource.
func DeleteEnv(r *client.ReleaseResource, containerName string, name string) error {

	for ci, container := range r.Containers {
		if container.Name == containerName {
			if len(container.Env) == 0 {
				return errors.New("This container has no Envs.")
			}
			for ei, env := range container.Env {
				if env.Name == name {
					r.Containers[ci].Env = append(r.Containers[ci].Env[:ei], r.Containers[ci].Env[ei+1:]...)
				} else {
					return errors.New("Env not found.")
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
