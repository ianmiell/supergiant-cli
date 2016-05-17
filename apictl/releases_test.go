package apictl

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// Tests the GetReleaseExample function
func TestGetReleaseExample(t *testing.T) {
	Convey("When a user requests a example release output", t, func() {
		Convey("We would expect the function to produce yaml and json output.", func() {
			err := GetReleaseExample()
			So(err, ShouldBeNil)
		})
	})
}

// Tests the GetRelease function
func TestGetRelease(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"name\":\"test\",\"custom_deploy_script\":null,\"current_release_id\":\"20160509194833\",\"target_release_id\":null,\"created\":\"Mon, 09 May 2016 19:48:33 UTC\",\"updated\":\"Mon, 09 May 2016 19:48:48 UTC\",\"tags\":{},\"addresses\":{\"external\":[{\"port\":\"test\",\"address\":\"tcp://test:80\"}],\"internal\":null}}")
	}))
	testURL := strings.Replace(ts.URL, "https://", "", -1)
	Convey("When fetching a supergiant release.", t, func() {
		//setup
		sg, err := NewClient(testURL, "test", "test")
		So(err, ShouldBeNil)

		Convey("We would expect the fetch to be successful", func() {
			release, err := sg.GetRelease("test", "test")
			So(err, ShouldBeNil)
			So(release, ShouldNotBeNil)
		})
		Convey("We would expect the fetch to fail if the component does not exist.", func() {
			release, err := sg.GetRelease("cheese", "cheese")
			So(err, ShouldBeNil)
			So(release, ShouldNotBeNil)
		})
	})
}
