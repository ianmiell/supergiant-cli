package apictl

import (
	"errors"

	"github.com/supergiant/supergiant/client"
	"github.com/supergiant/supergiant/common"
)

// CreateVolume creates a new volume for a resource.
func CreateVolume(r *client.ReleaseResource, name string, vtype string, size int) error {

	volume := &common.VolumeBlueprint{
		Name: &name,
		Type: vtype,
		Size: size,
	}

	r.Volumes = append(r.Volumes, volume)

	_, err := r.Save()
	if err != nil {
		return err
	}

	return nil
}

// UpdateVolume updates a volume for a resource.
func UpdateVolume(r *client.ReleaseResource, name string, vtype string, size int) error {
	/// Find our volume.
	if len(r.Volumes) == 0 {
		return errors.New("This Component has not volumes.")
	}
	for i, volume := range r.Volumes {
		if *volume.Name == name {
			// Edit the volume.
			if vtype != "gp2" {
				r.Volumes[i].Type = vtype
			}
			if size != 20 {
				r.Volumes[i].Size = size
			}
		} else {
			return errors.New("Volume not found.")
		}
	}

	_, err := r.Save()
	if err != nil {
		return err
	}

	return nil
}

// DeleteVolume creates a new volume for a resource.
func DeleteVolume(r *client.ReleaseResource, name string) error {

	if len(r.Volumes) == 0 {
		return errors.New("This Component has not volumes.")
	}
	for i, volume := range r.Volumes {
		if *volume.Name == name {
			r.Volumes = append(r.Volumes[:i], r.Volumes[i+1:]...)
		} else {
			return errors.New("Volume not found.")
		}
	}

	_, err := r.Save()
	if err != nil {
		return err
	}

	return nil
}

func volumeExist(r *client.ReleaseResource, volname string) bool {
	for _, volume := range r.Volumes {
		if *volume.Name == volname {
			return true
		}
	}
	return false
}
