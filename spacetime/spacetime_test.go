package spacetime

import (
	"os"
	"reflect"
	"testing"
)

func TestNewKube(t *testing.T) {
	os.Setenv("SG_CLI_TEST_MODE", "true")
	// ensure test db is destroyed
	sgdir := "/tmp/.supergiant"
	sgconfig := "" + sgdir + "/sgconfig.json"
	os.Remove(sgconfig)
	pvdrconfig := "" + sgdir + "/provider.json"
	os.Remove(pvdrconfig)

	NewProvider("test", "test", "test", "test", false)
	NewKube("test", "test", "test", "test", "test", "test", "test", false)

	expected := &Spacetime{
		Providers: map[string]Provider{
			"test": Provider{map[string]Region{
				"test": Region{map[string]Kube{
					"test": Kube{
						Name:        "test",
						User:        "test",
						Pass:        "test",
						Status:      "building",
						Region:      "test",
						AZ:          "test",
						Provider:    "test",
						KubeVersion: "test",
						Managed:     true,
						Selected:    true,
					},
				},
				},
			},
			},
		},
	}

	spacetime, _ := loadConfig()
	if !reflect.DeepEqual(expected, spacetime) {
		t.Error("Error: NewKube: expected,", expected, "-- But got:", spacetime)
	}
	os.Setenv("SG_CLI_TEST_MODE", "false")
}
