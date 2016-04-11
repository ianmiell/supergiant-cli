package spacetime

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
)

func testMode() bool {
	mode := os.Getenv("SG_CLI_TEST_MODE")
	if mode == "true" {
		return true
	}
	return false
}

// get the spacetime db file.
func getConfigFile() string {

	// get the current user.
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	// based on user data or test flag declare the config file.
	var sgdir string
	var sgconfig string
	if testMode() {
		sgdir = "/tmp/.supergiant"
		sgconfig = "" + sgdir + "/sgconfig.json"
	} else {
		sgdir = "" + user.HomeDir + "/.supergiant"
		sgconfig = "" + sgdir + "/sgconfig.json"
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
func (p *Spacetime) writeConfig() error {
	// Get file.
	config := getConfigFile()

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
func loadConfig() (*Spacetime, error) {
	// Get file.
	config := getConfigFile()

	// Read data from file.
	file, err := ioutil.ReadFile(config)
	if err != nil {
		panic(err)
	}

	// Load data into new spacetime object.
	pvdr := new(Spacetime)
	if err = json.Unmarshal(file, &pvdr); err != nil {
		pvdr = &Spacetime{Providers: map[string]Provider{}}
		return pvdr, nil
	}

	return pvdr, nil
}
