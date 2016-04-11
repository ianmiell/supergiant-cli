package apictl

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/supergiant/supergiant/client"
)

var errorString = "Error"

// DeployComponent Deploys a component live.
func DeployComponent(compName string, appName string) error {
	app, err := GetApp(appName)
	if err != nil {
		return err
	}

	comp, err := app.Components().Get(&compName)
	if err != nil {
		return err
	}

	err = comp.Deploy()
	if err != nil {
		return err
	}

	return nil
}

// GetComponent gets a component object in context.
func GetComponent(compName string, appName string) (*client.ComponentResource, error) {
	app, err := GetApp(appName)
	if err != nil {
		return nil, err
	}

	comp, err := app.Components().Get(&compName)
	if err != nil {
		return nil, err
	}
	return comp, nil
}

// CreateComponent creates a new component
func CreateComponent(compName string, appName string) error {
	// Get app
	app, err := GetApp(appName)
	if err != nil {
		return commonErrorParse(err, "Application Get, "+appName+"")
	}
	// mock comp
	comp := &client.Component{
		Name: &compName,
	}
	//create the comp
	newComp, err := app.Components().Create(comp)
	if err != nil {
		return commonErrorParse(err, "Component Create, "+*comp.Name+"")
	}
	// mock release
	release := releaseBoiler
	// Edit logic here.

	// Create Release.
	_, err = newComp.Releases().Create(release)
	if err != nil {
		return err
	}
	return nil
}

// DestroyComponent destroys a component
func DestroyComponent(compName string, appName string) error {
	app, err := GetApp(appName)
	if err != nil {
		return commonErrorParse(err, "Application Get, "+appName+"")
	}

	comp, err := app.Components().Get(&compName)
	if err != nil {
		return err
	}

	err = comp.Delete()
	if err != nil {
		return err
	}

	return nil
}

// ListComponents sends a list of componants to stdout
func ListComponents(appName string) error {
	app, err := GetApp(appName)
	if err != nil {
		return commonErrorParse(err, "Application Get, "+appName+"")
	}

	list, err := app.Components().List()
	if err != nil {
		return err
	}

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, "Name\tContainers\tCluster\tCreate Time\t")

	for _, comp := range list.Items {

		release, err := comp.TargetRelease()
		if err != nil {
			return err
		}
		// Get Name, create time
		compName := *comp.Name
		createTime := comp.Created.String()
		if err != nil {
			return err
		}
		containersCount := strconv.Itoa(len(release.Containers))

		fmt.Fprintln(w, ""+compName+"\t"+containersCount+"\tUnknown\t"+createTime+"\t")
	}
	w.Flush()
	return nil
}

// ListAllComponents sends a list of componants to stdout
func ListAllComponents() error {
	sg, err := getClient()
	if err != nil {
		return err
	}

	apps, err := sg.Apps().List()
	if err != nil {
		return err
	}

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, "Name\tApp\tContainers\tCluster\tCreate Time\tUpdate Time\t")

	for _, app := range apps.Items {

		list, err := app.Components().List()
		if err != nil {
			return err
		}

		for _, comp := range list.Items {
			release, err := comp.TargetRelease()
			if err != nil {
				return err
			}
			appName := checkForNilString(app.Name)
			compName := checkForNilString(comp.Name)
			createTime := comp.Created.String()
			updateTime := ""
			if comp.Updated == nil {
				updateTime = noData
			} else {
				updateTime = comp.Updated.String()
			}
			containersCount := strconv.Itoa(len(release.Containers))

			fmt.Fprintln(w, ""+compName+"\t"+appName+"\t"+containersCount+"\tUnknown\t"+createTime+"\t"+updateTime+"\t")
		}
	}
	w.Flush()
	return nil
}

// ComponentDetails returns a detailed description of a component to stdout.
func ComponentDetails(appName string, compName string) error {

	release, err := GetRelease(appName, compName)
	if err != nil {
		return err
	}

	fmt.Println(`
---- Release Information ----
Application: ` + appName + `
Component: ` + compName + `
Commited: ` + strconv.FormatBool(release.Committed) + `
Retired: ` + strconv.FormatBool(release.Retired) + `
Instances: ` + strconv.Itoa(release.InstanceCount) + `
Termination Grace Period: ` + strconv.Itoa(release.TerminationGracePeriod) + `
  Tags:
 --------------`)

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, "  Key\tValue\t")

	for Key, value := range release.Tags {
		fmt.Fprintln(w, "  "+Key+"\t"+value+"\t")
	}
	w.Flush()
	fmt.Println(`
  Volumes:
 --------------`)
	if len(release.Volumes) == 0 {
		fmt.Println(`      ` + noData + ``)
	}

	for _, volume := range release.Volumes {
		name := checkForNilString(volume.Name)
		size := strconv.Itoa(volume.Size)
		vtype := volume.Type
		fmt.Println(`  Name: ` + name + ` Size: ` + size + ` Type: ` + vtype + ``)
	}

	// Containers
	fmt.Println(`
  Containers:`)
	if len(release.Containers) == 0 {
		fmt.Println("  " + noData + "")
	}
	for _, container := range release.Containers {
		fmt.Println("  **************")
		// Build vars
		cpuMax := strconv.Itoa(int(container.CPU.Max))
		cpuMin := strconv.Itoa(int(container.CPU.Min))
		ramMax := strconv.Itoa(int(container.RAM.Max))
		ramMin := strconv.Itoa(int(container.RAM.Min))

		fmt.Println(`  Container: ` + container.Name + `
    Image: ` + container.Image + `
    Resources: CPU: max ` + cpuMax + ` / min ` + cpuMin + ` RAM: max ` + ramMax + ` / min ` + ramMin + ``)

		//Mounts

		fmt.Println(`    Mounts:
   --------------`)
		if len(container.Mounts) == 0 {
			fmt.Println(`      ` + noData + ``)
		}

		for _, mount := range container.Mounts {
			path := mount.Path
			volume := checkForNilString(mount.Volume)
			fmt.Println(`      Path: ` + path + ` Volume: ` + volume + ``)
		}

		// Ports

		fmt.Println(`
    Ports:
   --------------`)
		if len(container.Ports) == 0 {
			fmt.Println(`      ` + noData + ``)
		}

		for _, port := range container.Ports {
			protocal := port.Protocol
			pnumber := strconv.Itoa(port.Number)
			epnumber := strconv.Itoa(port.ExternalNumber)
			public := strconv.FormatBool(port.Public)
			entrypoint := checkForNilString(port.EntrypointDomain)
			fmt.Println(`    --Protocol: ` + protocal + ` Public: ` + public + ` Entrypoint: ` + entrypoint + ``)
			fmt.Println(`      Port Number: ` + pnumber + ` Ext Port: ` + epnumber + ``)
		}

		//ENV

		fmt.Println(`
    ENV Vars:
   --------------`)
		if len(container.Env) == 0 {
			fmt.Println(`      ` + noData + ``)
		}

		for _, env := range container.Env {
			name := env.Name
			value := env.Value
			fmt.Println(`      Name: ` + name + ` Value: ` + value + ``)
		}

		//Mounts

		fmt.Println(`
    Commands:
   --------------`)
		if len(container.Command) == 0 {
			fmt.Println(`      ` + noData + ``)
		}

		for _, command := range container.Command {
			fmt.Println(`      --- ` + command + ``)
		}
		fmt.Println("  **************")
	}
	return nil
}
