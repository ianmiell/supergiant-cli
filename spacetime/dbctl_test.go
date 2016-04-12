package spacetime

import (
	"os"
	"testing"

	"github.com/mitchellh/go-homedir"
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

	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}

	sgdir := "" + home + "/.supergiant"
	sgconfig := "" + sgdir + "/sgconfig.json"

	fileString := getConfigFile() //

	if fileString != sgconfig {
		t.Error("ERROR getConfigFile: expected", sgconfig, "-- But got:", fileString)
	}

}
