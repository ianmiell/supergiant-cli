package spacetime

import (
	"errors"
	"os"
)

//Kube describes a kubernetes cluster.
type Kube struct {
	Name          string `json:"name"`
	User          string `json:"user"`
	Pass          string `json:"pass"`
	IP            string `json:"ip"`
	SGID          string `json:"sg_id"`
	MinionSize    string `json:"minionSize"`
	KubeVersion   string `json:"kubeVersion"`
	Status        string `json:"status"`
	CoreInstalled bool   `json:"coreInstalled"`
	SgURL         string `json:"sgurl"`
	DashURL       string `json:"dashurl"`
	Region        string `json:"region"`
	AZ            string `json:"az"`
	Provider      string `json:"provider"`
	Log           string `json:"log"`
	Managed       bool   `json:"managed"`
	Selected      bool   `json:"selected"`
}

// NewKube creates a new Kube object and adds it to the current Spacetime database.
func NewKube(providerName string, regionName string, kubeName string, user string, pass string, az string, version string, verbose bool) error {

	// Fetch the current config
	spacetime, err := loadConfig()
	if err != nil {
		return err
	}

	// Build the new Kube object.
	kube := &Kube{
		Name:        kubeName,
		User:        user,
		Pass:        pass,
		Status:      "building",
		Region:      regionName,
		AZ:          az,
		Provider:    providerName,
		KubeVersion: version,
		Managed:     true,
		Selected:    true,
	}
	// add the new record to the Spacetime database.
	err = spacetime.AddKubeRecord(kube)
	if err != nil {
		return err
	}
	// Write the updated database to the database file.
	err = spacetime.writeConfig()
	if err != nil {
		return err
	}

	mode := os.Getenv("SG_CLI_TEST_MODE")
	if mode == "" {
		if verbose {
			err = kube.init(verbose)
		} else {
			go kube.init(verbose)
		}
		if err != nil {
			return err
		}
	}

	return nil

}

// DestroyKube will destroy a kube on it's provider and remove
// the kube from the spacetime database.
func DestroyKube(name string, force bool) error {

	kube, err := GetKube(name)
	if err != nil {
		return err
	}

	// Fetch the current config
	spacetime, err := loadConfig()
	if err != nil {
		return err
	}

	mode := os.Getenv("SG_CLI_TEST_MODE")
	if mode == "" && kube.Managed {
		err = kube.destroy(true)
		if err != nil {
			return err
		}
	}

	// add the new record to the Spacetime database.
	err = spacetime.deleteKube(name, true)
	if err != nil {
		return err
	}

	return nil

}

// RetryKube will re-run the terraform build for a given kube.
// This is useful if the build fails for some reason.
func RetryKube(name string, verbose bool) error {
	kube, err := GetKube(name)
	if err != nil {
		return err
	}

	mode := os.Getenv("SG_CLI_TEST_MODE")
	if mode == "" {
		err = kube.init(verbose)
		if err != nil {
			return err
		}
	}
	return nil
}

//GetKube Retruns a Kube object from the spacetime database.
func GetKube(name string) (*Kube, error) {
	list, _ := loadConfig()
	for _, provider := range list.Providers {
		for _, region := range provider.Regions {
			for _, kube := range region.Kubes {
				if kube.Name == name {
					return &kube, nil
				}
			}
		}
	}
	return nil, errors.New("Kube not found.")

}

// SetDefaultKube sets the default kube for supergiant
// context.
func SetDefaultKube(name string) error {
	// Get the old selected
	oldSelected, err := GetDefaultKube()
	if err != nil {
		// if no old selected.
		kube, gerr := GetKube(name)
		if gerr != nil {
			return gerr
		}
		kube.Selected = true
		err = kube.Update()
		if err != nil {
			return err
		}
		return nil
	}

	// if oldSelected, set old to false
	oldSelected.Selected = false
	err = oldSelected.Update()
	if err != nil {
		return err
	}

	// set requested to true.
	kube, err := GetKube(name)
	if err != nil {
		return err
	}
	kube.Selected = true
	err = kube.Update()
	if err != nil {
		return err
	}
	return nil
}

// GetDefaultKube gets the current selected default kube
func GetDefaultKube() (*Kube, error) {
	var kubeName Kube
	var faultCheck int
	// Fetch the current config
	spacetime, err := loadConfig()
	if err != nil {
		return &kubeName, err
	}

	for _, provider := range spacetime.Providers {
		for _, region := range provider.Regions {
			for _, kube := range region.Kubes {
				if kube.Selected {
					kubeName = kube
					faultCheck++
				}
			}
		}
	}

	if faultCheck > 1 {
		return &kubeName, errors.New("More then one kube is selected.")
	} else if faultCheck == 0 {
		return &kubeName, errors.New("No default kube is selected. Please use the select action or --kube flag to set context.")
	}

	return &kubeName, nil
}

// ImportKube imports a non-managed kube into the local spacetime database.
func ImportKube(name string, IP string, user string, pass string, region string, az string, provider string) error {
	// Fetch the current config
	spacetime, err := loadConfig()
	if err != nil {
		return err
	}

	// Build the new Kube object.
	kube := &Kube{
		Name:     name,
		User:     user,
		Pass:     pass,
		IP:       IP,
		Status:   "imported",
		Managed:  false,
		Selected: false,
		Region:   region,
		AZ:       az,
		Provider: provider,
	}

	err = spacetime.AddKubeRecord(kube)
	if err != nil {
		return err
	}

	// Write the updated database to the database file.
	err = spacetime.writeConfig()
	if err != nil {
		return err
	}

	err = SetDefaultKube(name)
	if err != nil {
		return err
	}

	return nil
}

// GetProvider returns the provider assigned to a kube resource.
func (k *Kube) GetProvider() (*ProviderConfig, error) {
	// Fetch the current config
	providerConfigs, err := loadProviderConfig()
	if err != nil {
		return nil, err
	}
	provider, err := providerConfigs.get(k.Provider)
	if err != nil {
		return nil, err
	}
	return provider, nil
}

// Update will update a kube's data in the Spacetime
//database. This is used to relay info to the database such as
// status, IP info, URLs etc.
func (k *Kube) Update() error {
	// Fetch the current config
	spacetime, err := loadConfig()
	if err != nil {
		return err
	}

	spacetime.Providers[k.Provider].Regions[k.Region].Kubes[k.Name] = *k

	// Write the updated database to the database file.
	err = spacetime.writeConfig()
	if err != nil {
		return err
	}
	return nil
}

func (k *Kube) fail() error {
	k.Status = "Failed"
	k.Update()
	return nil
}

func (s *Spacetime) deleteKube(kubeName string, force bool) error {
	for pn, provider := range s.Providers {
		for rn, region := range provider.Regions {
			for _, kube := range region.Kubes {
				if kube.Name == kubeName {
					delete(s.Providers[pn].Regions[rn].Kubes, kube.Name)
				}
			}
			if len(region.Kubes) == 0 {
				delete(s.Providers[pn].Regions, rn)
			}
		}
		if len(provider.Regions) == 0 {
			delete(s.Providers, pn)
		}
	}

	err := s.writeConfig()
	if err != nil {
		return err
	}
	return nil
}
