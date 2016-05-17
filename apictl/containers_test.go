package apictl

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/supergiant/supergiant/common"
)

// Tests the CreateContainer function
func TestCreateContainer(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"name\":\"test\",\"custom_deploy_script\":null,\"current_release_id\":\"20160509194833\",\"target_release_id\":null,\"created\":\"Mon, 09 May 2016 19:48:33 UTC\",\"updated\":\"Mon, 09 May 2016 19:48:48 UTC\",\"tags\":{},\"addresses\":{\"external\":[{\"port\":\"test\",\"address\":\"tcp://test:80\"}],\"internal\":null}}")
	}))
	testURL := strings.Replace(ts.URL, "https://", "", -1)
	Convey("When creating a new Supergiant container.", t, func() {
		//setup
		sg, err := NewClient(testURL, "test", "test")
		So(err, ShouldBeNil)
		release, err := sg.GetRelease("test", "test")
		So(err, ShouldBeNil)
		// one container
		release.Containers = append(release.Containers, &common.ContainerBlueprint{
			Name: "test",
		})

		release.Containers = append(release.Containers, &common.ContainerBlueprint{
			Name: "Hello World",
		})

		Convey("We would expect the release boiler to get removed.", func() {
			err := CreateContainer(release, "test", "test", "test", "test", "test", "test")
			So(err, ShouldBeNil)
		})

		Convey("We would expect the create to pass.", func() {
			err := CreateContainer(release, "test", "test", "test", "test", "test", "test")
			So(err, ShouldBeNil)
		})

		Convey("We would expect the create to fail if there is an api error.", func() {
			ts.Close()
			err := CreateContainer(release, "test", "test", "test", "test", "test", "test")
			So(err.Error(), ShouldContainSubstring, "Put https")
		})
	})
}

// Tests the UpdateContainer function
func TestUpdateContainer(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"name\":\"test\",\"custom_deploy_script\":null,\"current_release_id\":\"20160509194833\",\"target_release_id\":null,\"created\":\"Mon, 09 May 2016 19:48:33 UTC\",\"updated\":\"Mon, 09 May 2016 19:48:48 UTC\",\"tags\":{},\"addresses\":{\"external\":[{\"port\":\"test\",\"address\":\"tcp://test:80\"}],\"internal\":null}}")
	}))
	testURL := strings.Replace(ts.URL, "https://", "", -1)
	Convey("When updating a new Supergiant container.", t, func() {
		//setup
		sg, err := NewClient(testURL, "test", "test")
		So(err, ShouldBeNil)
		release, err := sg.GetRelease("test", "test")
		So(err, ShouldBeNil)

		Convey("We would expect the update to fail if the release has no containers.", func() {
			err := UpdateContainer(release, "test", "cheese", "", "", "", "")
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "This Component has not Containers.")
		})

		// one container
		release.Containers = append(release.Containers, &common.ContainerBlueprint{
			Name: "test",
		})

		Convey("We would expect the update to pass when updating an image.", func() {
			err := UpdateContainer(release, "test", "cheese", "", "", "", "")
			So(err, ShouldBeNil)
		})

		Convey("We would expect the update to fail when updating CPU and not providing both min and max value.", func() {
			err := UpdateContainer(release, "test", "cheese", "1", "", "", "")
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "Both CPU MAX, and MIN must be defined.")
		})

		Convey("We would expect the update to pass when CPU Max and MIN values.", func() {
			err := UpdateContainer(release, "test", "cheese", "1", "1", "", "")
			So(err, ShouldBeNil)
		})

		Convey("We would expect the update to fail when updating RAM and not providing both min and max value.", func() {
			err := UpdateContainer(release, "test", "cheese", "", "", "1", "")
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "Both RAM MAX, and MIN must be defined.")
		})

		Convey("We would expect the update to pass when RAM Max and MIN values.", func() {
			err := UpdateContainer(release, "test", "cheese", "", "", "1", "1")
			So(err, ShouldBeNil)
		})

		Convey("We would expect the update to fail if the release has containers, but the one specified is not found.", func() {
			err := UpdateContainer(release, "cheese", "cheese", "", "", "", "")
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "Container not found.")
		})

		Convey("We would expect the update to fail if there is an api error.", func() {
			ts.Close()
			err := UpdateContainer(release, "test", "", "", "", "", "")
			So(err.Error(), ShouldContainSubstring, "Put https")
		})
	})
}

// Tests the DeleteContainer function
func TestDeleteContainer(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"name\":\"test\",\"custom_deploy_script\":null,\"current_release_id\":\"20160509194833\",\"target_release_id\":null,\"created\":\"Mon, 09 May 2016 19:48:33 UTC\",\"updated\":\"Mon, 09 May 2016 19:48:48 UTC\",\"tags\":{},\"addresses\":{\"external\":[{\"port\":\"test\",\"address\":\"tcp://test:80\"}],\"internal\":null}}")
	}))
	testURL := strings.Replace(ts.URL, "https://", "", -1)
	Convey("When deleting a new Supergiant container.", t, func() {
		//setup
		sg, err := NewClient(testURL, "test", "test")
		So(err, ShouldBeNil)
		release, err := sg.GetRelease("test", "test")
		So(err, ShouldBeNil)
		// one container

		Convey("We would expect the delete to fail if the release has no containers.", func() {
			err := DeleteContainer(release, "test")
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "This Component has not Containers.")
		})

		release.Containers = append(release.Containers, &common.ContainerBlueprint{
			Name: "test",
		})

		Convey("We would expect the delete to pass.", func() {
			err := DeleteContainer(release, "test")
			So(err, ShouldBeNil)
		})

		Convey("We would expect the update to fail if the release has containers, but the one specified is not found.", func() {
			err := DeleteContainer(release, "cheese")
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "Container not found.")
		})

		Convey("We would expect the create to fail if there is an api error.", func() {
			ts.Close()
			err := DeleteContainer(release, "test")
			So(err.Error(), ShouldContainSubstring, "Put https")
		})
	})
}

// Tests the helloWorldExist function
func TestHelloWorldExist(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"name\":\"test\",\"custom_deploy_script\":null,\"current_release_id\":\"20160509194833\",\"target_release_id\":null,\"created\":\"Mon, 09 May 2016 19:48:33 UTC\",\"updated\":\"Mon, 09 May 2016 19:48:48 UTC\",\"tags\":{},\"addresses\":{\"external\":[{\"port\":\"test\",\"address\":\"tcp://test:80\"}],\"internal\":null}}")
	}))
	testURL := strings.Replace(ts.URL, "https://", "", -1)
	Convey("When checking if the boiler plate exists", t, func() {
		//setup
		sg, err := NewClient(testURL, "test", "test")
		So(err, ShouldBeNil)
		release, err := sg.GetRelease("test", "test")
		So(err, ShouldBeNil)

		Convey("We would expect the response to be false if it does not exist.", func() {
			So(helloWorldExist(release), ShouldEqual, false)
		})
		release.Containers = append(release.Containers, &common.ContainerBlueprint{
			Name: "test",
		})

		Convey("We would expect the response to be false if it does not exist, but the release has containers.", func() {
			So(helloWorldExist(release), ShouldEqual, false)
		})

		release.Containers = append(release.Containers, &common.ContainerBlueprint{
			Name: "Hello World",
		})

		Convey("We would expect the response to be true if it does exist.", func() {
			So(helloWorldExist(release), ShouldEqual, true)
		})
	})
}
