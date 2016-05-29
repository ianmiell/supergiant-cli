// Package spacetime controls kubernetes cluster creation, destruction,
//and stores the data in a flat file json database in the users home folder.
package spacetime

import (
	"fmt"
	"os"
)

// Region contains a map of kubernetes clusters.
type Region struct {
	Kubes map[string]Kube `json:"kubes"`
}

// Provider contains a map of regions.
type Provider struct {
	Regions map[string]Region `json:"regions"`
}

//Spacetime contains a map of providers.
type Spacetime struct {
	Providers map[string]Provider `json:"providers"`
}

// CommitAllKubes will mass install all kubes currently marked as staged.
//func CommitAllKubes(verbose bool) {
//	spacetime, err := loadConfig()
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(6)
//	}
//	// TODO: fork this call with go, notify user that things are happening and how to see it.
//	if verbose {
//		spacetime.commit(verbose)
//	} else {
//		go spacetime.commit(verbose)
//		fmt.Println("Commit action submitted...")
//		fmt.Println("Run \"sg spacetime list\" to track status.")
//		fmt.Println("You can also follow the commit log \"tail -f ~/.supergiant/<Cluster Name>/supergiant.log\"")
//	}
//}

// TODO: change sgconfig to spactime config

// RemoveKube removes a Kube object form the spacetime db.
func RemoveKube(kubeName string, force bool) {
	spacetime, err := loadConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(6)
	}

	spacetime.deleteKube(kubeName, force)
}
