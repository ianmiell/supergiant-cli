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

// Tests the CreateCreateMount function
func TestCreateMount(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"name\":\"test\",\"custom_deploy_script\":null,\"current_release_id\":\"20160509194833\",\"target_release_id\":null,\"created\":\"Mon, 09 May 2016 19:48:33 UTC\",\"updated\":\"Mon, 09 May 2016 19:48:48 UTC\",\"tags\":{},\"addresses\":{\"external\":[{\"port\":\"test\",\"address\":\"tcp://test:80\"}],\"internal\":null}}")
	}))
	testURL := strings.Replace(ts.URL, "https://", "", -1)
	Convey("When creating a new Supergiant Mount.", t, func() {
		//setup
		sg, err := NewClient(testURL, "test", "test")
		So(err, ShouldBeNil)
		release, err := sg.GetRelease("test", "test")
		So(err, ShouldBeNil)

		Convey("We would expect the create to be fail if the volume does not exist.", func() {
			err := CreateMount(release, "cheese", "test", "test")
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "Volume Does Not Exist...")
		})

		// One Volume
		release.Volumes = append(release.Volumes, &common.VolumeBlueprint{
			Name: common.IDString("test"),
			Type: "test",
			Size: 30,
		})

		// one container
		release.Containers = append(release.Containers, &common.ContainerBlueprint{
			Name: "test",
		})

		Convey("We would expect the create to fail if the container does not exist.", func() {
			err := CreateMount(release, "cheese", "test", "test")
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "Container Does Not Exist...")
		})

		Convey("We would expect the create to be successful if the volume exists.", func() {
			err := CreateMount(release, "test", "test", "test")
			So(err, ShouldBeNil)
		})
		Convey("We would expect the create to fail if there is an api error.", func() {
			ts.Close()
			err := CreateMount(release, "test", "test", "test")
			So(err.Error(), ShouldContainSubstring, "Put https")
		})
	})

}

// Tests the CreateCreateMount function
func TestUpdateMount(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"name\":\"test\",\"custom_deploy_script\":null,\"current_release_id\":\"20160509194833\",\"target_release_id\":null,\"created\":\"Mon, 09 May 2016 19:48:33 UTC\",\"updated\":\"Mon, 09 May 2016 19:48:48 UTC\",\"tags\":{},\"addresses\":{\"external\":[{\"port\":\"test\",\"address\":\"tcp://test:80\"}],\"internal\":null}}")
	}))
	testURL := strings.Replace(ts.URL, "https://", "", -1)
	Convey("When updating a Supergiant Mount.", t, func() {
		//setup
		sg, err := NewClient(testURL, "test", "test")
		So(err, ShouldBeNil)
		release, err := sg.GetRelease("test", "test")
		So(err, ShouldBeNil)

		Convey("We would expect the update to be fail if the volume does not exist.", func() {
			err := UpdateMount(release, "cheese", "test", "test")
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "Volume Does Not Exist...")
		})

		// One Volume
		release.Volumes = append(release.Volumes, &common.VolumeBlueprint{
			Name: common.IDString("test"),
			Type: "test",
			Size: 30,
		})

		// one container
		release.Containers = append(release.Containers, &common.ContainerBlueprint{
			Name: "test",
		})

		Convey("We would expect the update to fail if the container does not exist.", func() {
			err := UpdateMount(release, "cheese", "test", "test")
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "Container Does Not Exist...")
		})

		Convey("We would expect the update to be fail if there are no mounts.", func() {
			err := UpdateMount(release, "test", "test", "test")
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "This container has no mounts.")
		})

		release.Containers = []*common.ContainerBlueprint{}
		release.Containers = append(release.Containers, &common.ContainerBlueprint{
			Name: "test",
			Mounts: []*common.Mount{
				&common.Mount{
					Path:   "/test",
					Volume: common.IDString("test"),
				},
			},
		})

		Convey("We would expect the update to be successful if the volume exists.", func() {
			err := UpdateMount(release, "test", "/test", "test")
			So(err, ShouldBeNil)
		})
		Convey("We would expect the update to be fail if there are mounts, but we can't find the mount specified.", func() {
			err := UpdateMount(release, "test", "cheese", "test")
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "Mount not found.")
		})
		Convey("We would expect the update to fail if there is an api error.", func() {
			ts.Close()
			err := UpdateMount(release, "test", "/test", "test")
			So(err.Error(), ShouldContainSubstring, "Put https")
		})
	})

}

// Tests the CreateDeleteMount function
func TestDeleteMount(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"name\":\"test\",\"custom_deploy_script\":null,\"current_release_id\":\"20160509194833\",\"target_release_id\":null,\"created\":\"Mon, 09 May 2016 19:48:33 UTC\",\"updated\":\"Mon, 09 May 2016 19:48:48 UTC\",\"tags\":{},\"addresses\":{\"external\":[{\"port\":\"test\",\"address\":\"tcp://test:80\"}],\"internal\":null}}")
	}))
	testURL := strings.Replace(ts.URL, "https://", "", -1)
	Convey("When deleting a Supergiant Mount.", t, func() {
		//setup
		sg, err := NewClient(testURL, "test", "test")
		So(err, ShouldBeNil)
		release, err := sg.GetRelease("test", "test")
		So(err, ShouldBeNil)

		// One Volume
		release.Volumes = append(release.Volumes, &common.VolumeBlueprint{
			Name: common.IDString("test"),
			Type: "test",
			Size: 30,
		})
		// one container
		release.Containers = append(release.Containers, &common.ContainerBlueprint{
			Name: "test",
		})

		Convey("We would expect the delete to fail if there are no mounts.", func() {
			err := DeleteMount(release, "test", "test")
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "This container has no mounts.")
		})

		// one container
		release.Containers = []*common.ContainerBlueprint{}
		release.Containers = append(release.Containers, &common.ContainerBlueprint{
			Name: "test",
			Mounts: []*common.Mount{
				&common.Mount{
					Path:   "/test",
					Volume: common.IDString("test"),
				},
			},
		})

		Convey("We would expect the delete to fail if the container does not exist.", func() {
			err := DeleteMount(release, "cheese", "test")
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "Container Does Not Exist...")
		})

		Convey("We would expect the delete to be successful if the path exists.", func() {
			err := DeleteMount(release, "test", "/test")
			So(err, ShouldBeNil)
		})
		Convey("We would expect the delete to fail if there are mounts, but we can't find the mount specified.", func() {
			err := DeleteMount(release, "test", "cheese")
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "Mount not found.")
		})
		Convey("We would expect the delete to fail if there is an api error.", func() {
			ts.Close()
			err := DeleteMount(release, "test", "/test")
			So(err.Error(), ShouldContainSubstring, "Put https")
		})
	})

}
