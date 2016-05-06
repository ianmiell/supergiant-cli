package sgcore

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/briandowns/spinner"
	"github.com/supergiant/guber"
	"github.com/supergiant/supergiant-cli/apictl"
	"github.com/supergiant/supergiant-cli/spacetime"
)

func initNamespace(c guber.Client) error {
	namespace := &guber.Namespace{
		Metadata: &guber.Metadata{
			Name: "supergiant",
		},
	}

	out, err := c.Namespaces().Create(namespace)
	if err != nil {
		return err
	}
	fmt.Println(out)
	return nil
}

// InstallSGCore installs a spergiant instance to a running kubernetes cluster.
func InstallSGCore(name string) error {

	// Fetch the referenced kube.
	kube, err := spacetime.GetKube(name)
	if err != nil {
		return err
	}

	// Create a new kubernetes client.
	client := guber.NewClient(kube.IP, kube.User, kube.Pass, true)

	//Create the "supergiant" namespace.
	err = initNamespace(client)
	if err != nil {
		fmt.Println("WARN: Namespace exists.")
	}
	//Start ETCD in the namspace.
	provider, err := kube.GetProvider()
	if err != nil {
		return err
	}
	err = initETCD(client, provider, kube)
	if err != nil {
		fmt.Println("WARN: ETCD exists.")
	} //
	err = initETCDBrowser(client)
	if err != nil {
		fmt.Println("WARN: ETCD BROWSER exists.")
	} //
	//Start the supergiant api.
	err = initSGAPI(client, kube, "")
	if err != nil {
		fmt.Println("WARN: ETCD BROWSER exists.", err)
	} //
	fmt.Println("Waiting for core to settle... This can take a few minutes.")
	s := spinner.New(spinner.CharSets[1], 100*time.Millisecond)
	s.Start()
	for i := 0; i < 200; i++ {

		_, err = apictl.GetApp("")
		if err != nil {
			time.Sleep(2 * time.Second)
		} else {
			s.Stop()
			fmt.Println("Supergiant API verified...")
			break
		}
	}
	if err != nil {
		return errors.New("Supergiant API Install Failed.")
	}

	//Start the dashboard.
	dash, err := initDash(client, "")
	if err != nil {
		return err
	}

	fmt.Println("Waiting port detect...")
	s = spinner.New(spinner.CharSets[1], 100*time.Millisecond)
	s.Start()
	var service *guber.Service
	for i := 0; i < 200; i++ {

		service, err = client.Services("supergiant").Get("sg-ui-public")
		if err != nil {
			time.Sleep(2 * time.Second)
		} else {
			s.Stop()
			fmt.Println("Port detected...")
			break
		}
	}
	if err != nil {
		return errors.New("Supergiant UI Port detection failed.")
	}

	var uiPort int
	for _, port := range service.Spec.Ports {
		if port.Name == "9001" {
			uiPort = port.NodePort
		}
	}

	// Success
	kube.CoreInstalled = true
	kube.SgURL = "https://" + kube.User + ":" + kube.Pass + "@" + kube.IP + "/api/v1/proxy/namespaces/supergiant/services/supergiant-api:frontend/"
	kube.DashURL = "http://" + dash + ":" + strconv.Itoa(uiPort) + ""
	kube.Update()

	fmt.Println("Waiting Dashboard to be active...")
	s = spinner.New(spinner.CharSets[1], 100*time.Millisecond)
	s.Start()
	for i := 0; i < 200; i++ {
		err := err
		resp, err := http.Get(kube.DashURL)
		if err != nil {
			time.Sleep(2 * time.Second)
		} else {
			if resp.StatusCode == 200 {
				s.Stop()
				fmt.Println("Dashboard Detected...")
				fmt.Println("Dashboard URL:", kube.DashURL)
				break
			}
		}
	}
	if err != nil {
		return errors.New("Dashboard detection failed.")
	}

	return nil
}

// DestroySGCore deletes a supegiant instance from a kubrnetes cluster.
func DestroySGCore(name string) error {

	// Fetch the referenced kube.
	kube, err := spacetime.GetKube(name)
	if err != nil {
		return err
	}

	// Create a client connection to kubernetes.
	client := guber.NewClient(kube.IP, kube.User, kube.Pass, true)

	// Get a list of namespaces.
	namespaces := client.Namespaces()
	err = destroyDash()
	if err != nil {
		return err
	}
	// Delete the supergiant namespace (and all resources).
	namespaces.Delete("supergiant")

	// Success.
	kube.CoreInstalled = false
	kube.SgURL = ""
	kube.Update()
	return nil
}

// UpdateSGCore updates the installation of supergiant to the default
// or the specified version of supergiant-api
func UpdateSGCore(kubeName string, version string) error {
	// Fetch the referenced kube.
	kube, err := spacetime.GetKube(kubeName)
	if err != nil {
		return err
	}
	// Create a new kubernetes client.
	client := guber.NewClient(kube.IP, kube.User, kube.Pass, true)

	_, err = client.Namespaces().Get("supergiant")
	if err != nil {
		return errors.New("Supergiant does not appear to be installed on this cluster. You must install Supergiant with \"create core\" before you can upgrade.")
	}

	// Delete the RC and wait.
	var svCount int
	for {
		err := client.Services("supergiant").Delete("supergiant-api")
		if err != nil {
			break
		}

		fmt.Println("Waiting for service to delete...")
		time.Sleep(5 * time.Second)
		svCount++

		if svCount > 25 {
			return errors.New("Upgrade failed while removing supergiant service.")
		}
	}

	// Delete the RC and wait.
	var rcCount int
	for {
		err := client.ReplicationControllers("supergiant").Delete("supergiant-api")
		if err != nil {
			break
		}

		fmt.Println("Waiting for rc to delete...")
		time.Sleep(5 * time.Second)
		rcCount++

		if rcCount > 25 {
			return errors.New("Upgrade failed while removing supergiant replication controller.")
		}
	}

	// Delete the pods and wait.
	q := &guber.QueryParams{
		LabelSelector: "instance=supergiant-api",
	}

	pods, err := client.Pods("supergiant").Query(q)
	if err == nil {
		// if err is nil, meaning pods do exist after RC delete.
		for _, pod := range pods.Items {
			// Delete each pod and wait.
			var podCounter int
			for {
				err := client.Pods("supergiant").Delete(pod.Metadata.Name)
				if err != nil {
					break
				}
				fmt.Println("Waiting for pods to delete...")
				time.Sleep(5 * time.Second)
				podCounter++
				if podCounter > 25 {
					return errors.New("Upgrade failed while removing supergiant pods.")
				}
			}
		}
	}

	fmt.Println("Installing Supergiant...")
	//Start the supergiant api.
	err = initSGAPI(client, kube, version)
	if err != nil {
		return err
	}

	return nil
}

func UpdateDash(name string, version string) error {

	// Fetch the referenced kube.
	kube, err := spacetime.GetKube(name)
	if err != nil {
		return err
	}

	// Create a new kubernetes client.
	client := guber.NewClient(kube.IP, kube.User, kube.Pass, true)
	// destroy old dash
	fmt.Println("Removing old Dashboard...")
	err = destroyDash()
	if err != nil {
		fmt.Println("WARN: Dash does not appear to exist. trying to continue.")
	}
	fmt.Println("Waiting for Dashboard to delete...")
	time.Sleep(5 * time.Second)

	fmt.Println("Updating dashboard to version: " + version + "...")
	//Start the dashboard.
	dash, err := initDash(client, version)
	if err != nil {
		return err
	}

	fmt.Println("Waiting port detect...")
	s := spinner.New(spinner.CharSets[1], 100*time.Millisecond)
	s.Start()
	var service *guber.Service
	for i := 0; i < 200; i++ {

		service, err = client.Services("supergiant").Get("sg-ui-public")
		if err != nil {
			time.Sleep(2 * time.Second)
		} else {
			s.Stop()
			fmt.Println("Port detected...")
			break
		}
	}
	if err != nil {
		return errors.New("Supergiant UI Port detection failed.")
	}

	var uiPort int
	for _, port := range service.Spec.Ports {
		if port.Name == "9001" {
			uiPort = port.NodePort
		}
	}

	// Success
	kube.CoreInstalled = true
	kube.DashURL = "http://" + dash + ":" + strconv.Itoa(uiPort) + ""
	kube.Update()

	fmt.Println("Waiting Dashboard to be active...")
	s = spinner.New(spinner.CharSets[1], 100*time.Millisecond)
	s.Start()
	for i := 0; i < 200; i++ {
		err := err
		resp, err := http.Get(kube.DashURL)
		if err != nil {
			time.Sleep(2 * time.Second)
		} else {
			if resp.StatusCode == 200 {
				s.Stop()
				fmt.Println("Dashboard Detected...")
				fmt.Println("Dashboard URL:", kube.DashURL)
				break
			}
		}
	}
	if err != nil {
		return errors.New("Dashboard detection failed.")
	}
	return nil
}
