package apictl

//import (
//	"encoding/json"
//	"io/ioutil"
//	"os"
//	"os/user"
//)//

//func testMode() bool {
//	mode := os.Getenv("SG_CLI_TEST_MODE")
//	if mode == "true" {
//		return true
//	}
//	return false
//}//

//// get the spacetime db file.
//func getConfigFile() string {//

//	// get the current user.
//	user, err := user.Current()
//	if err != nil {
//		panic(err)
//	}//

//	// based on user data or test flag declare the config file.
//	var sgdir string
//	var appcache string
//	if testMode() {
//		sgdir = "/tmp/.supergiant"
//		appcache = "" + sgdir + "/appcache.json"
//	} else {
//		sgdir = "" + user.HomeDir + "/.supergiant"
//		appcache = "" + sgdir + "/appcache.json"
//	}//

//	// Make sure the appcache directory exists.
//	_, err = os.Stat(sgdir)
//	// else create it.
//	if os.IsNotExist(err) {
//		err := os.Mkdir(sgdir, 0700)
//		if err != nil {
//			panic(err)
//		}
//	}//

//	// Make sure appcache file exists.
//	_, err = os.Stat(appcache)
//	// else create it.
//	if os.IsNotExist(err) {
//		_, err := os.Create(appcache)
//		if err != nil {
//			panic(err)
//		} else {
//			return appcache
//		}
//	}//

//	return appcache//

//}//

//// Wride a spacetime object to the config db.
//func (p *cache) writeConfig() error {
//	// Get file.
//	config := getConfigFile()//

//	// Convert object to json.
//	data, err := json.MarshalIndent(p, "", "  ")
//	if err != nil {
//		return err
//	}//

//	// Write json output to file.
//	err = ioutil.WriteFile(config, data, 0644)
//	if err != nil {
//		return err
//	}
//	return nil
//}//

//// Fetch spacetime object from the json db file.
//func loadConfig() (*cache, error) {
//	// Get file.
//	config := getConfigFile()//

//	// Read data from file.
//	file, err := ioutil.ReadFile(config)
//	if err != nil {
//		panic(err)
//	}//

//	// Load data into new spacetime object.
//	pvdr := new(cache)
//	if err = json.Unmarshal(file, &pvdr); err != nil {
//		pvdr = &cache{}
//		return pvdr, nil
//	}//

//	return pvdr, nil
//}
