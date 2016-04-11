package spacetime

import (
	"errors"
	"fmt"
	"os"
	"text/tabwriter"
)

// ProviderConfig holds details about a ProviderConfig
type ProviderConfig struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
	SSHKey    string `json:"sshKey"`
	Status    string `json:"staged"`
	Log       string `json:"log"`
}

// ProviderConfigs holds a list of ProviderConfig objects.
type ProviderConfigs struct {
	ProviderConfigs map[string]ProviderConfig `json:"providerConfigs"`
}

// NewProvider adds a new provider configurations to the provider db.
func NewProvider(name string, accessKey string, secretKey string, hstPvdr string, verbose bool) error {

	// Fetch the current config
	providerConfigs, err := loadProviderConfig()
	if err != nil {
		return err
	}

	newProvider := &ProviderConfig{
		Name:      name,
		Type:      hstPvdr,
		AccessKey: accessKey,
		SecretKey: secretKey,
		Status:    "staged",
	}
	// TODO: Add function to create and attach a ssh key for the provider.
	err = providerConfigs.addProviderRecord(newProvider)
	if err != nil {
		return err
	}
	err = providerConfigs.writeProviderConfig()
	if err != nil {
		return err
	}
	err = newProvider.init(verbose)
	if err != nil {
		return err
	}
	return nil
}

// RebuildProvider rebuilds an existing provider.
func RebuildProvider(name string, verbose bool) error {
	// Fetch the current config
	providerConfigs, err := loadProviderConfig()
	if err != nil {
		return err
	}

	provider, _ := providerConfigs.get(name)

	if verbose {
		err = provider.init(verbose)
		if err != nil {
			return err
		}
	} else {
		go provider.init(verbose)
	}
	return nil
}

// ListProvider shows a list of provider configurations from the provider json db.
// If non empty string is passed, it will try to list a detailed view of the requested
// provider config.
func ListProvider(provider string) error {
	// get providers
	providers, err := loadProviderConfig()
	if err != nil {
		return err
	}

	// if user is looking for details
	if provider != "" {
		//search for the provider
		pvdr, _ := providers.get(provider)
		if pvdr == nil {
			return errors.New("Provider Not Found.")
		}
		pvdr.listProviderDetails()
		return nil
	}

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, "Provider\tType\tStatus \tAccess Key\tSecret Key\t")

	for _, providerConfig := range providers.ProviderConfigs {

		fmt.Fprintln(w, ""+providerConfig.Name+"\t"+providerConfig.Type+"\t"+providerConfig.Status+" \t"+providerConfig.AccessKey+"\t"+providerConfig.SecretKey+"\t")
	}
	w.Flush()
	return nil
}

// DeleteProvider removes a ProviderConfig from the provider db.
func DeleteProvider(name string) error {
	// Fetch the current config
	providerConfigs, err := loadProviderConfig()
	if err != nil {
		return err
	}

	err = providerConfigs.deleteProviderRecord(name)
	if err != nil {
		return err
	}
	err = providerConfigs.writeProviderConfig()
	if err != nil {
		return err
	}
	return nil
}

func (p *ProviderConfig) listProviderDetails() {
	fmt.Println(`
Provider Details:
	Name: ` + p.Name + `
	  Type: ` + p.Type + `
		Status: ` + p.Status + `
		AWS Access Key: ` + p.AccessKey + `
		AWS Secret Key: ` + p.SecretKey + `
		Log File: ` + p.Log + `
		`)

}

func (p *ProviderConfigs) get(name string) (*ProviderConfig, error) {
	for _, provider := range p.ProviderConfigs {
		if provider.Name == name {
			return &provider, nil
		}
	}
	return nil, errors.New("Provider not found.")
}

func (p *ProviderConfigs) providerExist(name string) bool {
	for key := range p.ProviderConfigs {
		if key == name {
			return true
		}
	}
	return false
}

func (p *ProviderConfigs) addProviderRecord(c *ProviderConfig) error {
	result := p.providerExist(c.Name)
	if result {
		return errors.New("Provider Already Exists.")
	}

	p.ProviderConfigs[c.Name] = *c

	return nil
}
func (p *ProviderConfigs) deleteProviderRecord(name string) error {
	delete(p.ProviderConfigs, name)
	return nil
}

func (p *ProviderConfig) update() error {
	// Fetch the current config
	// get providers
	providers, err := loadProviderConfig()
	if err != nil {
		return err
	}

	providers.ProviderConfigs[p.Name] = *p

	// Write the updated database to the database file.
	err = providers.writeProviderConfig()
	if err != nil {
		return err
	}
	return nil
}
