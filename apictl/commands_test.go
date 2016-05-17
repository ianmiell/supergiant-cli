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

// Tests the CreateCommand function
func TestCreateCommand(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"name\":\"test\",\"custom_deploy_script\":null,\"current_release_id\":\"20160509194833\",\"target_release_id\":null,\"created\":\"Mon, 09 May 2016 19:48:33 UTC\",\"updated\":\"Mon, 09 May 2016 19:48:48 UTC\",\"tags\":{},\"addresses\":{\"external\":[{\"port\":\"test\",\"address\":\"tcp://test:80\"}],\"internal\":null}}")
	}))
	testURL := strings.Replace(ts.URL, "https://", "", -1)
	Convey("When creating a new Supergiant command.", t, func() {
		//setup
		sg, err := NewClient(testURL, "test", "test")
		So(err, ShouldBeNil)
		release, err := sg.GetRelease("test", "test")
		So(err, ShouldBeNil)
		// one container
		release.Containers = append(release.Containers, &common.ContainerBlueprint{
			Name: "test",
		})
		Convey("We would expect the create to pass.", func() {
			err := CreateCommand(release, "test", []string{"test", "test"})
			So(err, ShouldBeNil)
		})

		Convey("We would expect the create to fail if there is an api error.", func() {
			ts.Close()
			err := CreateCommand(release, "test", []string{"test", "test"})
			So(err.Error(), ShouldContainSubstring, "Put https")
		})
	})
}

// Tests the DeleteCommand function
func TestDeleteCommand(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"name\":\"test\",\"custom_deploy_script\":null,\"current_release_id\":\"20160509194833\",\"target_release_id\":null,\"created\":\"Mon, 09 May 2016 19:48:33 UTC\",\"updated\":\"Mon, 09 May 2016 19:48:48 UTC\",\"tags\":{},\"addresses\":{\"external\":[{\"port\":\"test\",\"address\":\"tcp://test:80\"}],\"internal\":null}}")
	}))
	testURL := strings.Replace(ts.URL, "https://", "", -1)
	Convey("When deleting a Supergiant command.", t, func() {
		//setup
		sg, err := NewClient(testURL, "test", "test")
		So(err, ShouldBeNil)
		release, err := sg.GetRelease("test", "test")
		So(err, ShouldBeNil)
		// one container
		release.Containers = append(release.Containers, &common.ContainerBlueprint{
			Name: "test",
		})
		Convey("We would expect the delete to fail if the container has no commands.", func() {
			err := DeleteCommand(release, "test", []string{"test"})
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "This container has no Commands.")
		})

		release.Containers = []*common.ContainerBlueprint{}
		release.Containers = append(release.Containers, &common.ContainerBlueprint{
			Name:    "test",
			Command: []string{"test"},
		})

		Convey("We would expect the delete to pass.", func() {
			err := DeleteCommand(release, "test", []string{"test"})
			So(err, ShouldBeNil)
		})

		Convey("We would expect the delete to fail if there is an api error.", func() {
			ts.Close()
			err := DeleteCommand(release, "test", []string{"test"})
			So(err.Error(), ShouldContainSubstring, "Put https")
		})
	})
}
