package apictl

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/supergiant/supergiant/client"
)

// CreateEntryPoint makes a new supergiant endpoint.
func (sg *SGClient) CreateEntryPoint(name string) error {
	entrypoint := &client.Entrypoint{
		Domain: &name,
	}

	_, err := sg.Entrypoints().Create(entrypoint)
	if err != nil {
		return err
	}
	return nil
}

// DestroyEntryPoint deletes an entrypoint
func (sg *SGClient) DestroyEntryPoint(name string) error {
	entrypoint, err := sg.Entrypoints().Get(&name)
	if err != nil {
		return err
	}

	err = entrypoint.Delete()
	if err != nil {
		return err
	}

	return nil
}

// GetEntryURL returns the URL of an entrypoint
func (sg *SGClient) GetEntryURL(s string) (string, error) {
	list, err := sg.Entrypoints().List()
	if err != nil {
		return "", err
	}
	for _, entry := range list.Items {
		if *entry.Domain == s {
			return entry.Address, nil
		}
	}
	return "", nil
}

// ListEntryPoints sends a list of entrypoints to stdout.
func (sg *SGClient) ListEntryPoints(name string) error {
	list, err := sg.Entrypoints().List()
	if err != nil {
		return err
	}

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, "ID\tAddress\tComponents\tCluster\tCreate Time\tUpdate Time\t")

	for _, entrypoint := range list.Items {
		// build values.

		entrypointCreated := entrypoint.Created.String()
		if entrypointCreated == "" {
			entrypointCreated = noData
		}
		entrypointAddress := entrypoint.Address
		if entrypointAddress == "" {
			entrypointAddress = noData
		}
		entrypointDomain := *entrypoint.Domain
		if entrypoint.Domain == nil {
			entrypointDomain = noData
		}
		//entrypointUpdated := entrypoint.Updated.String()
		//if entrypointUpdated == "" {
		//	entrypointUpdated = noData
		//}
		//entrypointComps, err := getEntrypointComps(entrypointDomain)
		//if err != nil {
		//	return err
		//}
		//entrypointCompCount := strconv.Itoa(len(entrypointComps))

		fmt.Fprintln(w, ""+entrypointDomain+"\t"+entrypointAddress+"\t0\tUnknown\t"+entrypointCreated+"\tUnknown\t")
	}
	w.Flush()
	return nil
}

// slice of comps assigned to a entrypoint
func (sg *SGClient) getEntrypointComps(entrypointID string) ([]string, error) {
	var ecomps []string

	apps, err := sg.Apps().List()
	if err != nil {
		return ecomps, err
	}

	// nast biz here...
	for _, app := range apps.Items {
		comps, err := app.Components().List()
		if err != nil {
			return ecomps, err
		}
		for _, comp := range comps.Items {
			release, err := comp.CurrentRelease()
			if err != nil {
				return ecomps, err
			}
			for _, container := range release.Containers {
				for _, port := range container.Ports {
					entry := *port.EntrypointDomain
					if port.EntrypointDomain == nil {
						entry = "NULL"
					}
					if entry == entrypointID {
						ecomps = append(ecomps, entrypointID)
					}
				}
			}
		}

	}
	return ecomps, nil
}
