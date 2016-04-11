package spacetime

import (
	"errors"

	"github.com/supergiant/guber"
)

func checkForLife(ip string, user string, pass string) error {
	client := guber.NewClient(ip, user, pass, true)

	namespaces, err := client.Namespaces().List()
	if err != nil {
		return err
	}
	for _, items := range namespaces.Items {
		if items.Metadata.Name == "kube-system" {
			return nil
		}
	}
	return errors.New("Kubernetes appears to be up, but the kube-system namespace is not there.")
}
