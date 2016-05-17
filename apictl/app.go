package apictl

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/supergiant/supergiant/client"
)

// GetApp gets a app object in context.
func (sg *SGClient) GetApp(appName string) (*client.AppResource, error) {

	app, err := sg.Apps().Get(&appName)
	if err != nil {
		return nil, err
	}
	return app, nil
}

// CreateApp makes a new supergiant app.
func (sg *SGClient) CreateApp(name string) error {
	app := &client.App{
		Name: &name,
	}

	_, err := sg.Apps().Create(app)
	if err != nil {
		return err
	}
	return nil
}

//DestroyApp deletes a app
func (sg *SGClient) DestroyApp(name string) error {

	app, err := sg.Apps().Get(&name)
	if err != nil {
		return err
	}

	err = app.Delete()
	if err != nil {
		return err
	}

	return nil
}

// ListApps sends a list of apps to stdout
func (sg *SGClient) ListApps(name string) error {

	list, err := sg.Apps().List()
	if err != nil {
		return err
	}

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, "Name\tComponents\tCluster\tCreate Time\t")

	for _, app := range list.Items {
		appName := *app.Name
		createTime := app.Created.String()
		comps, err := sg.getAppComps(appName)
		compCount := strconv.Itoa(len(comps))
		if err != nil {
			compCount = "Error"
		}

		fmt.Fprintln(w, ""+appName+"\t"+compCount+"\tUnknown\t"+createTime+"\t")
	}
	w.Flush()
	return nil
}

func (sg *SGClient) getAppComps(appName string) ([]*client.Component, error) {
	var rcomps []*client.Component

	app, err := sg.Apps().Get(&appName)
	if err != nil {
		return rcomps, err
	}

	comps, err := app.Components().List()
	if err != nil {
		return rcomps, err
	}

	for _, comp := range comps.Items {
		if *app.Name == appName {
			rcomps = append(rcomps, comp.Component)
		}
	}
	return rcomps, nil
}
