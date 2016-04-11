package spacetime

import "errors"

// Check if a kube exists at all...
func (s *Spacetime) doesKubeExist(name string) bool {
	for _, provider := range s.Providers {
		for _, region := range provider.Regions {
			for id := range region.Kubes {
				if id == name {
					return true
				}
			}
		}
	}
	return false
}

// Add a new provider to spacetime
func (s *Spacetime) addNewProvider(k *Kube) error {
	// if provider does not exist in spacetime add it
	if _, ok := s.Providers[k.Provider]; !ok {
		s.Providers[k.Provider] = Provider{map[string]Region{
			k.Region: Region{map[string]Kube{
				k.Name: *k,
			},
			},
		},
		}
		return nil
	}
	//else err
	return errors.New("Provider already exists..")
}

// Add a new region to a provider
func (s *Spacetime) addNewRegion(k *Kube) error {
	// if region does not exist in provider add it.
	if _, ok := s.Providers[k.Provider].Regions[k.Region]; !ok {
		s.Providers[k.Provider].Regions[k.Region] = Region{map[string]Kube{
			k.Name: *k,
		},
		}
		return nil
	}
	//else err
	return errors.New("Region already exists..")
}

// Add a new kube to a region
func (s *Spacetime) addNewKube(k *Kube) error {
	//if kube does not exist in region add it.
	if _, ok := s.Providers[k.Provider].Regions[k.Region].Kubes[k.Name]; !ok {
		s.Providers[k.Provider].Regions[k.Region].Kubes[k.Name] = *k
		return nil
	}
	//else error
	return errors.New("Kube already exists..")
}

// AddKubeRecord adds a new kubernetes cluster record to a spacetime object
func (s *Spacetime) AddKubeRecord(k *Kube) error { //

	//check if new kube exists anywhere...
	result := s.doesKubeExist(k.Name)
	if result {
		return errors.New("Kubernetes cluster: " + k.Name + " already exists.")
	}
	// try to add a new provider
	err := s.addNewProvider(k)
	if err != nil {
		//if provider already exists try to add a new region
		err := s.addNewRegion(k)
		if err != nil {
			// if region already exists try to add a new kube
			s.addNewKube(k)
		}
	}
	return nil

}
