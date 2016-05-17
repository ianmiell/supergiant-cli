package apictl

import (
	"errors"

	"github.com/supergiant/supergiant/client"
	"github.com/supergiant/supergiant/common"
)

// CreatePort creates a new volume for a resource.
func CreatePort(r *client.ReleaseResource, containerName string, proto string, num int, pub bool, entry string, enum int) error {

	port := &common.Port{
		Protocol:         proto,
		Number:           num,
		Public:           pub,
		EntrypointDomain: &entry,
		ExternalNumber:   enum,
	}

	fault := false
	for _, container := range r.Containers {
		if container.Name == containerName {
			container.Ports = append(container.Ports, port)
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

// UpdatePort updates a port for a container resource.
func UpdatePort(r *client.ReleaseResource, containerName string, proto string, num int, pub bool, entry string, enum int) error {

	fault := false
	for ci, container := range r.Containers {
		if container.Name == containerName {
			if len(container.Ports) == 0 {
				return errors.New("This container has no ports.")
			}
			for pi, port := range container.Ports {
				if port.Number == num {
					if proto != "" {
						r.Containers[ci].Ports[pi].Protocol = proto
					}
					if num != 0 {
						r.Containers[ci].Ports[pi].Number = num
					}
					if pub {
						r.Containers[ci].Ports[pi].Public = pub
					}
					if entry != "" {
						r.Containers[ci].Ports[pi].EntrypointDomain = &entry
					}
					if enum != 0 {
						r.Containers[ci].Ports[pi].ExternalNumber = enum
					}
				} else {
					return errors.New("Port not found.")
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

// DeletePort deletes a port for a container resource.
func DeletePort(r *client.ReleaseResource, containerName string, num int) error {

	fault := false
	for ci, container := range r.Containers {
		if container.Name == containerName {
			if len(container.Ports) == 0 {
				return errors.New("This container has no ports.")
			}
			for pi, port := range container.Ports {
				if port.Number == num {
					r.Containers[ci].Ports = append(r.Containers[ci].Ports[:pi], r.Containers[ci].Ports[pi+1:]...)
				} else {
					return errors.New("Port not found.")
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
