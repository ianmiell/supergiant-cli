package spacetime

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

// TODO: create details list function to look at a kube in detail.

//ListKubes dumps out data from the json db to the CLI
func ListKubes(kube string) { //list all kubes
	list, _ := loadConfig()

	if kube != "" {
		k, _ := GetKube(kube)
		k.listDetailed()
		return
	}

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, "Name\tProvider \tAZ\tIP\tStatus\tCore Status\tManaged\tSelected\t")

	for pn, provider := range list.Providers {
		for _, region := range provider.Regions {
			for _, kube := range region.Kubes {
				fmt.Fprintln(w, ""+strings.TrimSpace(kube.Name)+"\t"+strings.TrimSpace(pn)+" \t"+strings.TrimSpace(kube.AZ)+"\t"+strings.TrimSpace(kube.IP)+"\t"+strings.TrimSpace(kube.Status)+"\t"+strconv.FormatBool(kube.CoreInstalled)+"\t"+strconv.FormatBool(kube.Managed)+"\t"+strconv.FormatBool(kube.Selected)+"\t")
			}
		}
	}
	w.Flush()
}

func (k *Kube) listDetailed() {
	fmt.Println(`
Name: ` + k.Name + `
	User: ` + k.User + `
	Pass: ` + k.Pass + `
	IP: ` + k.IP + `
	Minion Size: ` + k.MinionSize + `
	Kube Version: ` + k.KubeVersion + `
	Kube Status: ` + k.Status + `
  Core Installed?: ` + strconv.FormatBool(k.CoreInstalled) + `
  Region: ` + k.Region + `
  AZ: ` + k.AZ + `
  Provider: ` + k.Provider + `
  Log: ` + k.Log + `
	Super Giant URL: ` + k.SgURL + `
		`)
}
