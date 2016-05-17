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

// Tests the CreateVolume function
func TestCreateVolume(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"name\":\"test\",\"custom_deploy_script\":null,\"current_release_id\":\"20160509194833\",\"target_release_id\":null,\"created\":\"Mon, 09 May 2016 19:48:33 UTC\",\"updated\":\"Mon, 09 May 2016 19:48:48 UTC\",\"tags\":{},\"addresses\":{\"external\":[{\"port\":\"test\",\"address\":\"tcp://test:80\"}],\"internal\":null}}")
	}))
	testURL := strings.Replace(ts.URL, "https://", "", -1)
	Convey("When creating a new Supergiant Volume.", t, func() {
		//setup
		sg, err := NewClient(testURL, "test", "test")
		So(err, ShouldBeNil)
		release, err := sg.GetRelease("test", "test")
		So(err, ShouldBeNil)

		Convey("We would expect the create to be successful", func() {
			err := CreateVolume(release, "test", "test", 40)
			So(err, ShouldBeNil)
		})
		Convey("We would expect the create to fail if there is an api error.", func() {
			ts.Close()
			err := CreateVolume(release, "test", "test", 40)
			So(err.Error(), ShouldContainSubstring, "Put https")
		})
	})

}

// Tests the UpdateVolume function
func TestUpdateVolume(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"name\":\"test\",\"custom_deploy_script\":null,\"current_release_id\":\"20160509194833\",\"target_release_id\":null,\"created\":\"Mon, 09 May 2016 19:48:33 UTC\",\"updated\":\"Mon, 09 May 2016 19:48:48 UTC\",\"tags\":{},\"addresses\":{\"external\":[{\"port\":\"test\",\"address\":\"tcp://test:80\"}],\"internal\":null}}")
	}))
	testURL := strings.Replace(ts.URL, "https://", "", -1)
	Convey("When updating a Supergiant Volume.", t, func() {
		sg, err := NewClient(testURL, "test", "test")
		So(err, ShouldBeNil)
		release, err := sg.GetRelease("test", "test")
		So(err, ShouldBeNil)

		// TEST CASES
		Convey("We would expect update of a volume to error if the component has no volumes.", func() {
			err = UpdateVolume(release, "test", "test", 20)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "This Component has not volumes.")
		})
		// One Volume
		release.Volumes = append(release.Volumes, &common.VolumeBlueprint{
			Name: common.IDString("test"),
			Type: "test",
			Size: 30,
		})
		Convey("We would expect update of a volume to not error if the volume is found.", func() {
			err = UpdateVolume(release, "test", "test", 40)
			So(err, ShouldBeNil)
		})
		Convey("We would expect update of a volume to error if there are volumes but the volume is not found.", func() {
			err = UpdateVolume(release, "cheese", "test", 40)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "Volume not found.")
		})
		Convey("We would expect the volume update to error if there is an api error.", func() {
			ts.Close()
			err = UpdateVolume(release, "test", "test", 40)
			So(err.Error(), ShouldContainSubstring, "Put https")
		})
	})

}

// Tests the DeleteVolume function
func TestDeleteVolume(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"name\":\"test\",\"custom_deploy_script\":null,\"current_release_id\":\"20160509194833\",\"target_release_id\":null,\"created\":\"Mon, 09 May 2016 19:48:33 UTC\",\"updated\":\"Mon, 09 May 2016 19:48:48 UTC\",\"tags\":{},\"addresses\":{\"external\":[{\"port\":\"test\",\"address\":\"tcp://test:80\"}],\"internal\":null}}")
	}))
	testURL := strings.Replace(ts.URL, "https://", "", -1)
	Convey("When deleting a Supergiant Volume.", t, func() {
		//setup steps
		sg, err := NewClient(testURL, "test", "test")
		So(err, ShouldBeNil)
		release, err := sg.GetRelease("test", "test")
		So(err, ShouldBeNil)

		Convey("We would expect the volume delete to error if it dos not have any volumes.", func() {
			delerr := DeleteVolume(release, "cheese")
			So(delerr, ShouldNotBeNil)
			So(delerr.Error(), ShouldEqual, "This Component has not volumes.")
		})

		// One Volume
		release.Volumes = append(release.Volumes, &common.VolumeBlueprint{
			Name: common.IDString("test"),
			Type: "test",
			Size: 30,
		})

		Convey("We would expect the volume to delete without error.", func() {
			delerr := DeleteVolume(release, "test")
			So(delerr, ShouldBeNil)
		})
		Convey("We would expect the volume delete to error if it dos not exist.", func() {
			delerr := DeleteVolume(release, "cheese")
			So(delerr, ShouldNotBeNil)
			So(delerr.Error(), ShouldEqual, "Volume not found.")
		})
		Convey("We would expect the volume delete to error if there is an api error.", func() {
			ts.Close()
			delerr := DeleteVolume(release, "test")
			So(delerr.Error(), ShouldContainSubstring, "Put https")
		})
	})

}
