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

// Tests the CreateEnv function
func TestCreateEnv(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"name\":\"test\",\"custom_deploy_script\":null,\"current_release_id\":\"20160509194833\",\"target_release_id\":null,\"created\":\"Mon, 09 May 2016 19:48:33 UTC\",\"updated\":\"Mon, 09 May 2016 19:48:48 UTC\",\"tags\":{},\"addresses\":{\"external\":[{\"port\":\"test\",\"address\":\"tcp://test:80\"}],\"internal\":null}}")
	}))
	testURL := strings.Replace(ts.URL, "https://", "", -1)
	Convey("When creating a new Supergiant env.", t, func() {
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
			err := CreateEnv(release, "test", "test", "test")
			So(err, ShouldBeNil)
		})

		Convey("We would expect the create to fail if there is an api error.", func() {
			ts.Close()
			err := CreateEnv(release, "test", "test", "test")
			So(err.Error(), ShouldContainSubstring, "Put https")
		})
	})
}

// Tests the UpdateEnv function
func TestUpdateEnv(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"name\":\"test\",\"custom_deploy_script\":null,\"current_release_id\":\"20160509194833\",\"target_release_id\":null,\"created\":\"Mon, 09 May 2016 19:48:33 UTC\",\"updated\":\"Mon, 09 May 2016 19:48:48 UTC\",\"tags\":{},\"addresses\":{\"external\":[{\"port\":\"test\",\"address\":\"tcp://test:80\"}],\"internal\":null}}")
	}))
	testURL := strings.Replace(ts.URL, "https://", "", -1)
	Convey("When updating a new Supergiant env.", t, func() {
		//setup
		sg, err := NewClient(testURL, "test", "test")
		So(err, ShouldBeNil)
		release, err := sg.GetRelease("test", "test")
		So(err, ShouldBeNil)

		// one container
		release.Containers = append(release.Containers, &common.ContainerBlueprint{
			Name: "test",
		})

		Convey("We would expect the update to fail if there are no envs.", func() {
			err := UpdateEnv(release, "test", "test", "test2")
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "This container has no Envs.")
		})

		Convey("We would expect the update to fail if the container is not found.", func() {
			err := UpdateEnv(release, "cheese", "test", "test2")
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "Container Does Not Exist...")
		})

		// one container
		release.Containers = []*common.ContainerBlueprint{}
		release.Containers = append(release.Containers, &common.ContainerBlueprint{
			Name: "test",
			Env: []*common.EnvVar{
				&common.EnvVar{
					Name:  "test",
					Value: "test",
				},
			},
		})
		Convey("We would expect the update to pass.", func() {
			err := UpdateEnv(release, "test", "test", "test2")
			So(err, ShouldBeNil)
		})

		Convey("We would expect the update to fail if there are envs, but the requested one is not found.", func() {
			err := UpdateEnv(release, "test", "cheese", "test2")
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "Env not found.")
		})

		Convey("We would expect the update to fail if there is an api error.", func() {
			ts.Close()
			err := UpdateEnv(release, "test", "test", "test2")
			So(err.Error(), ShouldContainSubstring, "Put https")
		})
	})
}

// Tests the DeleteEnv function
func TestDeleteEnv(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"name\":\"test\",\"custom_deploy_script\":null,\"current_release_id\":\"20160509194833\",\"target_release_id\":null,\"created\":\"Mon, 09 May 2016 19:48:33 UTC\",\"updated\":\"Mon, 09 May 2016 19:48:48 UTC\",\"tags\":{},\"addresses\":{\"external\":[{\"port\":\"test\",\"address\":\"tcp://test:80\"}],\"internal\":null}}")
	}))
	testURL := strings.Replace(ts.URL, "https://", "", -1)
	Convey("When deleting a new Supergiant env.", t, func() {
		//setup
		sg, err := NewClient(testURL, "test", "test")
		So(err, ShouldBeNil)
		release, err := sg.GetRelease("test", "test")
		So(err, ShouldBeNil)
		// one container
		release.Containers = append(release.Containers, &common.ContainerBlueprint{
			Name: "test",
		})

		Convey("We would expect the update to fail if there are no envs.", func() {
			err := DeleteEnv(release, "test", "test")
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "This container has no Envs.")
		})

		Convey("We would expect the update to fail if the container is not found.", func() {
			err := DeleteEnv(release, "cheese", "test")
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "Container Does Not Exist...")
		})

		// one container
		release.Containers = []*common.ContainerBlueprint{}
		release.Containers = append(release.Containers, &common.ContainerBlueprint{
			Name: "test",
			Env: []*common.EnvVar{
				&common.EnvVar{
					Name:  "test",
					Value: "test",
				},
			},
		})

		Convey("We would expect the delete to pass.", func() {
			err := DeleteEnv(release, "test", "test")
			So(err, ShouldBeNil)
		})

		Convey("We would expect the update to fail if there are envs, but the requested one is not found.", func() {
			err := DeleteEnv(release, "test", "cheese")
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "Env not found.")
		})

		Convey("We would expect the create to fail if there is an api error.", func() {
			ts.Close()
			err := DeleteEnv(release, "test", "test")
			So(err.Error(), ShouldContainSubstring, "Put https")
		})
	})
}
