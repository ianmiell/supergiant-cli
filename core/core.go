package sgcore

import (
	"errors"
	"fmt"
	"time"

	"github.com/supergiant/guber"
	"github.com/supergiant/supergiant-cli/spacetime"
)

func initNamespace(c *guber.Client) error {
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
		return err
	}

	// Start ETCD in the namspace.
	provider, err := kube.GetProvider()
	if err != nil {
		return err
	}
	err = initETCD(client, provider, kube)
	if err != nil {
		return err
	}

	err = initETCDBrowser(client)
	if err != nil {
		return err
	}

	//Start the supergiant api.
	err = initSGAPI(client, kube, "")
	if err != nil {
		return err
	}

	// Success
	kube.CoreInstalled = true
	kube.SgURL = "https://" + kube.User + ":" + kube.Pass + "@" + kube.IP + "/api/v1/proxy/namespaces/supergiant/services/supergiant-api:frontend/"
	kube.Update()

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
		found, _ := client.Services("supergiant").Delete("supergiant-api")
		if !found {
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
		found, _ := client.ReplicationControllers("supergiant").Delete("supergiant-api")
		if !found {
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
				found, _ := client.Pods("supergiant").Delete(pod.Metadata.Name)
				if !found {
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
