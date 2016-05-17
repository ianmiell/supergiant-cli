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

// Tests the CreatePort function
func TestCreatePort(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"name\":\"test\",\"custom_deploy_script\":null,\"current_release_id\":\"20160509194833\",\"target_release_id\":null,\"created\":\"Mon, 09 May 2016 19:48:33 UTC\",\"updated\":\"Mon, 09 May 2016 19:48:48 UTC\",\"tags\":{},\"addresses\":{\"external\":[{\"port\":\"test\",\"address\":\"tcp://test:80\"}],\"internal\":null}}")
	}))
	testURL := strings.Replace(ts.URL, "https://", "", -1)
	Convey("When creating a new Supergiant Port.", t, func() {
		//setup
		sg, err := NewClient(testURL, "test", "test")
		So(err, ShouldBeNil)
		release, err := sg.GetRelease("test", "test")
		So(err, ShouldBeNil)

		// one container
		release.Containers = append(release.Containers, &common.ContainerBlueprint{
			Name: "test",
		})

		Convey("We would expect the create to fail if the container does not exist.", func() {
			err := CreatePort(release, "cheese", "HTTP", 1234, true, "test", 1234)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "Container Does Not Exist...")
		})

		Convey("We would expect the create to be successful.", func() {
			err := CreatePort(release, "test", "HTTP", 1234, true, "test", 1234)
			So(err, ShouldBeNil)
		})
		Convey("We would expect the create to fail if there is an api error.", func() {
			ts.Close()
			err := CreatePort(release, "test", "HTTP", 1234, true, "test", 1234)
			So(err.Error(), ShouldContainSubstring, "Put https")
		})
	})
}

// Tests the UpdatePort function
func TestUpdatePort(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"name\":\"test\",\"custom_deploy_script\":null,\"current_release_id\":\"20160509194833\",\"target_release_id\":null,\"created\":\"Mon, 09 May 2016 19:48:33 UTC\",\"updated\":\"Mon, 09 May 2016 19:48:48 UTC\",\"tags\":{},\"addresses\":{\"external\":[{\"port\":\"test\",\"address\":\"tcp://test:80\"}],\"internal\":null}}")
	}))
	testURL := strings.Replace(ts.URL, "https://", "", -1)
	Convey("When updating a new Supergiant Port.", t, func() {
		//setup
		sg, err := NewClient(testURL, "test", "test")
		So(err, ShouldBeNil)
		release, err := sg.GetRelease("test", "test")
		So(err, ShouldBeNil)

		// one container
		release.Containers = append(release.Containers, &common.ContainerBlueprint{
			Name: "test",
		})

		Convey("We would expect the update to fail if the container does not have ports.", func() {
			err := UpdatePort(release, "test", "HTTP", 1234, true, "test", 1234)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "This container has no ports.")
		})

		// one container
		release.Containers = []*common.ContainerBlueprint{}
		release.Containers = append(release.Containers, &common.ContainerBlueprint{
			Name: "test",
			Ports: []*common.Port{
				&common.Port{
					Protocol:         "HTTP",
					Number:           1234,
					Public:           true,
					EntrypointDomain: common.IDString("test"),
					ExternalNumber:   1234,
				},
			},
		})

		Convey("We would expect the update to fail if the container does not exist.", func() {
			err := UpdatePort(release, "cheese", "HTTP", 1234, true, "test", 1234)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "Container Does Not Exist...")
		})

		Convey("We would expect the update to fail if the container has ports, but we didn't find the right one.", func() {
			err := UpdatePort(release, "test", "HTTP", 4567, true, "test", 1234)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "Port not found.")
		})

		Convey("We would expect the update to be successful.", func() {
			err := UpdatePort(release, "test", "HTTP", 1234, true, "test", 1234)
			So(err, ShouldBeNil)
		})
		Convey("We would expect the update to fail if there is an api error.", func() {
			ts.Close()
			err := UpdatePort(release, "test", "HTTP", 1234, true, "test", 1234)
			So(err.Error(), ShouldContainSubstring, "Put https")
		})
	})
}

// Tests the DeletePort function
func TestDeletePort(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"name\":\"test\",\"custom_deploy_script\":null,\"current_release_id\":\"20160509194833\",\"target_release_id\":null,\"created\":\"Mon, 09 May 2016 19:48:33 UTC\",\"updated\":\"Mon, 09 May 2016 19:48:48 UTC\",\"tags\":{},\"addresses\":{\"external\":[{\"port\":\"test\",\"address\":\"tcp://test:80\"}],\"internal\":null}}")
	}))
	testURL := strings.Replace(ts.URL, "https://", "", -1)
	Convey("When deleting a Supergiant Port.", t, func() {
		//setup
		sg, err := NewClient(testURL, "test", "test")
		So(err, ShouldBeNil)
		release, err := sg.GetRelease("test", "test")
		So(err, ShouldBeNil)

		// one container
		release.Containers = append(release.Containers, &common.ContainerBlueprint{
			Name: "test",
		})

		Convey("We would expect the delete to fail if the container does not have ports.", func() {
			err := DeletePort(release, "test", 1234)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "This container has no ports.")
		})

		// one container
		release.Containers = []*common.ContainerBlueprint{}
		release.Containers = append(release.Containers, &common.ContainerBlueprint{
			Name: "test",
			Ports: []*common.Port{
				&common.Port{
					Protocol:         "HTTP",
					Number:           1234,
					Public:           true,
					EntrypointDomain: common.IDString("test"),
					ExternalNumber:   1234,
				},
			},
		})

		Convey("We would expect the delete to fail if the container does not exist.", func() {
			err := DeletePort(release, "cheese", 1234)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "Container Does Not Exist...")
		})

		Convey("We would expect the delete to be successful.", func() {
			err := DeletePort(release, "test", 1234)
			So(err, ShouldBeNil)
		})

		Convey("We would expect the delete to fail if the container has ports, but we didn't find the right one.", func() {
			err := DeletePort(release, "test", 4567)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "Port not found.")
		})

		Convey("We would expect the delete to fail if there is an api error.", func() {
			ts.Close()
			err := DeletePort(release, "test", 1234)
			So(err.Error(), ShouldContainSubstring, "Put https")
		})
	})
}
