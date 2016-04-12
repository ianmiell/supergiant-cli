package spacetime

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/mitchellh/go-homedir"
)

// get the spacetime db file.
func getProviderConfigFile() string {

	// get the current user.
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}

	// based on user data or test flag declare the config file.
	var sgdir string
	var sgconfig string
	if testMode() {
		sgdir = "/tmp/.supergiant"
		sgconfig = "" + sgdir + "/provider.json"
	} else {
		sgdir = "" + home + "/.supergiant"
		sgconfig = "" + sgdir + "/provider.json"
	}

	// Make sure the sgconfig directory exists.
	_, err = os.Stat(sgdir)
	// else create it.
	if os.IsNotExist(err) {
		err := os.Mkdir(sgdir, 0700)
		if err != nil {
			panic(err)
		}
	}

	// Make sure sgconfig file exists.
	_, err = os.Stat(sgconfig)
	// else create it.
	if os.IsNotExist(err) {
		_, err := os.Create(sgconfig)
		if err != nil {
			panic(err)
		} else {
			return sgconfig
		}
	}

	return sgconfig

}

// Wride a spacetime object to the config db.
func (p *ProviderConfigs) writeProviderConfig() error {
	// Get file.
	config := getProviderConfigFile()

	// Convert object to json.
	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		panic(err)
	}

	// Write json output to file.
	err = ioutil.WriteFile(config, data, 0644)
	if err != nil {
		panic(err)
	}
	return nil
}

// Fetch spacetime object from the json db file.
func loadProviderConfig() (*ProviderConfigs, error) {
	// Get file.
	config := getProviderConfigFile()

	// Read data from file.
	file, err := ioutil.ReadFile(config)
	if err != nil {
		panic(err)
	}

	// Load data into new spacetime object.
	pvdr := new(ProviderConfigs)
	if err = json.Unmarshal(file, &pvdr); err != nil {
		pvdr = &ProviderConfigs{ProviderConfigs: map[string]ProviderConfig{}}
		return pvdr, nil
	}

	return pvdr, nil
}
