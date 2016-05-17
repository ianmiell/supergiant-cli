package apictl

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// Tests the CreateApp function
func TestCreateApp(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"name\":\"test\",\"custom_deploy_script\":null,\"current_release_id\":\"20160509194833\",\"target_release_id\":null,\"created\":\"Mon, 09 May 2016 19:48:33 UTC\",\"updated\":\"Mon, 09 May 2016 19:48:48 UTC\",\"tags\":{},\"addresses\":{\"external\":[{\"port\":\"test\",\"address\":\"tcp://test:80\"}],\"internal\":null}}")
	}))
	testURL := strings.Replace(ts.URL, "https://", "", -1)
	Convey("When creating a supergiant app.", t, func() {
		//setup
		sg, err := NewClient(testURL, "test", "test")
		So(err, ShouldBeNil)

		Convey("We would expect the create to be successful", func() {
			err := sg.CreateApp("test")
			So(err, ShouldBeNil)
		})
	})
}

// Tests the DestroyApp function
func TestDestroyApp(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"name\":\"test\",\"custom_deploy_script\":null,\"current_release_id\":\"20160509194833\",\"target_release_id\":null,\"created\":\"Mon, 09 May 2016 19:48:33 UTC\",\"updated\":\"Mon, 09 May 2016 19:48:48 UTC\",\"tags\":{},\"addresses\":{\"external\":[{\"port\":\"test\",\"address\":\"tcp://test:80\"}],\"internal\":null}}")
	}))
	testURL := strings.Replace(ts.URL, "https://", "", -1)
	Convey("When destroying a supergiant app.", t, func() {
		//setup
		sg, err := NewClient(testURL, "test", "test")
		So(err, ShouldBeNil)

		Convey("We would expect the destroy to be successful", func() {
			err := sg.DestroyApp("test")
			So(err, ShouldBeNil)
		})
	})
}

// Tests the ListApps function
func TestListApps(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"name\":\"test\",\"custom_deploy_script\":null,\"current_release_id\":\"20160509194833\",\"target_release_id\":null,\"created\":\"Mon, 09 May 2016 19:48:33 UTC\",\"updated\":\"Mon, 09 May 2016 19:48:48 UTC\",\"tags\":{},\"addresses\":{\"external\":[{\"port\":\"test\",\"address\":\"tcp://test:80\"}],\"internal\":null}}")
	}))
	testURL := strings.Replace(ts.URL, "https://", "", -1)
	Convey("When listing supergiant apps.", t, func() {
		//setup
		sg, err := NewClient(testURL, "test", "test")
		So(err, ShouldBeNil)

		Convey("We would expect the list to be successful", func() {
			err := sg.ListApps("test")
			So(err, ShouldBeNil)
		})
	})
}

// Tests the getAppComps function
func TestGetAppComps(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"name\":\"test\",\"custom_deploy_script\":null,\"current_release_id\":\"20160509194833\",\"target_release_id\":null,\"created\":\"Mon, 09 May 2016 19:48:33 UTC\",\"updated\":\"Mon, 09 May 2016 19:48:48 UTC\",\"tags\":{},\"addresses\":{\"external\":[{\"port\":\"test\",\"address\":\"tcp://test:80\"}],\"internal\":null}}")
	}))
	testURL := strings.Replace(ts.URL, "https://", "", -1)
	Convey("When listing components for a supergiant app.", t, func() {
		//setup
		sg, err := NewClient(testURL, "test", "test")
		So(err, ShouldBeNil)

		Convey("We would expect the get to be successful", func() {
			_, err := sg.getAppComps("test")
			So(err, ShouldBeNil)
		})
	})
}
