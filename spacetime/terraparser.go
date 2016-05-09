package spacetime

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type state struct {
	Modules []struct {
		Resources struct {
			AwsInstanceKubeMaster struct {
				Primary struct {
					Attributes struct {
						PublicIP string `json:"public_ip"`
					} `json:"attributes"`
				} `json:"primary"`
			} `json:"aws_instance.kube_master"`
			AwsSecurityGroupELBSg struct {
				Primary struct {
					ID string `json:"id"`
				} `json:"primary"`
			} `json:"aws_security_group.elb_sg"`
		} `json:"resources"`
	} `json:"modules"`
}

func terraStateParser(stateFile string) (string, string, error) {
	_, err := os.Stat(stateFile)
	if err != nil {
		return "", "", err
	}
	// Read data from file.
	file, err := ioutil.ReadFile(stateFile)
	if err != nil {
		return "", "", err
	}

	// Load data into new state object.
	pvdr := new(state)
	if err = json.Unmarshal(file, &pvdr); err != nil {
		return "", "", err
	}

	return pvdr.Modules[0].Resources.AwsInstanceKubeMaster.Primary.Attributes.PublicIP,
		pvdr.Modules[0].Resources.AwsSecurityGroupELBSg.Primary.ID,
		nil

	//fmt.Println(pvdr)
}
