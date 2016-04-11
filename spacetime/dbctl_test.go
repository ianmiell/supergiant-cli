package spacetime

import (
	"os"
	"os/user"
	"testing"
)

func TestGetConfigFileTest(t *testing.T) { //

	os.Setenv("SG_CLI_TEST_MODE", "true")

	sgdir := "/tmp/.supergiant"
	sgconfig := "" + sgdir + "/sgconfig.json" //

	fileString := getConfigFile() //

	if fileString != sgconfig {
		t.Error("ERROR getConfigFile: expected", sgconfig, "-- But got:", fileString)
	}
	os.Setenv("SG_CLI_TEST_MODE", "false")
}
func TestGetConfigFileProd(t *testing.T) {

	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	sgdir := "" + user.HomeDir + "/.supergiant"
	sgconfig := "" + sgdir + "/sgconfig.json"

	fileString := getConfigFile() //

	if fileString != sgconfig {
		t.Error("ERROR getConfigFile: expected", sgconfig, "-- But got:", fileString)
	}

}
