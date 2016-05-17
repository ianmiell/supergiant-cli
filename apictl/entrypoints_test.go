package apictl

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// Tests the CreateEntryPoint function
func TestCreateEntryPoint(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"name\":\"test\",\"custom_deploy_script\":null,\"current_release_id\":\"20160509194833\",\"target_release_id\":null,\"created\":\"Mon, 09 May 2016 19:48:33 UTC\",\"updated\":\"Mon, 09 May 2016 19:48:48 UTC\",\"tags\":{},\"addresses\":{\"external\":[{\"port\":\"test\",\"address\":\"tcp://test:80\"}],\"internal\":null}}")
	}))
	testURL := strings.Replace(ts.URL, "https://", "", -1)
	Convey("When creating a new Supergiant Entrypoint.", t, func() {
		//setup
		sg, err := NewClient(testURL, "test", "test")
		So(err, ShouldBeNil)

		Convey("We would expect the create to be successful", func() {
			err := sg.CreateEntryPoint("test")
			So(err, ShouldBeNil)
		})
		Convey("We would expect the create to fail if there is an api error.", func() {
			ts.Close()
			err := sg.CreateEntryPoint("test")
			So(err.Error(), ShouldContainSubstring, "Post https")
		})
	})
}

// Tests the DestroyEntryPoint function
func TestDestroyEntryPoint(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"name\":\"test\",\"custom_deploy_script\":null,\"current_release_id\":\"20160509194833\",\"target_release_id\":null,\"created\":\"Mon, 09 May 2016 19:48:33 UTC\",\"updated\":\"Mon, 09 May 2016 19:48:48 UTC\",\"tags\":{},\"addresses\":{\"external\":[{\"port\":\"test\",\"address\":\"tcp://test:80\"}],\"internal\":null}}")
	}))
	testURL := strings.Replace(ts.URL, "https://", "", -1)
	Convey("When destroying a Supergiant Entrypoint.", t, func() {
		//setup
		sg, err := NewClient(testURL, "test", "test")
		So(err, ShouldBeNil)

		Convey("We would expect the destroy to be successful", func() {
			err := sg.DestroyEntryPoint("test")
			So(err, ShouldBeNil)
		})
		Convey("We would expect the destroy to fail if there is an api error.", func() {
			ts.Close()
			err := sg.DestroyEntryPoint("test")
			So(err.Error(), ShouldContainSubstring, "Get https")
		})
	})
}

// Tests the GetEntryURL function
func TestGetEntryURL(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"name\":\"test\",\"custom_deploy_script\":null,\"current_release_id\":\"20160509194833\",\"target_release_id\":null,\"created\":\"Mon, 09 May 2016 19:48:33 UTC\",\"updated\":\"Mon, 09 May 2016 19:48:48 UTC\",\"tags\":{},\"addresses\":{\"external\":[{\"port\":\"test\",\"address\":\"tcp://test:80\"}],\"internal\":null}}")
	}))
	testURL := strings.Replace(ts.URL, "https://", "", -1)
	Convey("When getting a Supergiant Entrypoint URL.", t, func() {
		//setup
		sg, err := NewClient(testURL, "test", "test")
		So(err, ShouldBeNil)

		Convey("We would expect the get to be successful", func() {
			_, err := sg.GetEntryURL("test")
			So(err, ShouldBeNil)
		})
		Convey("We would expect the get to fail if there is an api error.", func() {
			ts.Close()
			_, err := sg.GetEntryURL("test")
			So(err.Error(), ShouldContainSubstring, "Get https")
		})
	})
}

// Tests the ListEntryPoints function
func TestListEntryPoints(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"name\":\"test\",\"custom_deploy_script\":null,\"current_release_id\":\"20160509194833\",\"target_release_id\":null,\"created\":\"Mon, 09 May 2016 19:48:33 UTC\",\"updated\":\"Mon, 09 May 2016 19:48:48 UTC\",\"tags\":{},\"addresses\":{\"external\":[{\"port\":\"test\",\"address\":\"tcp://test:80\"}],\"internal\":null}}")
	}))
	testURL := strings.Replace(ts.URL, "https://", "", -1)
	Convey("When listing a Supergiant Entrypoints.", t, func() {
		//setup
		sg, err := NewClient(testURL, "test", "test")
		So(err, ShouldBeNil)

		Convey("We would expect the list to be successful", func() {
			err := sg.ListEntryPoints("test")
			So(err, ShouldBeNil)
		})
		Convey("We would expect the list to fail if there is an api error.", func() {
			ts.Close()
			err := sg.ListEntryPoints("test")
			So(err.Error(), ShouldContainSubstring, "Get https")
		})
	})
}

// Tests the getEntrypointComps function
func TestGetEntrypointComps(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"name\":\"test\",\"custom_deploy_script\":null,\"current_release_id\":\"20160509194833\",\"target_release_id\":null,\"created\":\"Mon, 09 May 2016 19:48:33 UTC\",\"updated\":\"Mon, 09 May 2016 19:48:48 UTC\",\"tags\":{},\"addresses\":{\"external\":[{\"port\":\"test\",\"address\":\"tcp://test:80\"}],\"internal\":null}}")
	}))
	testURL := strings.Replace(ts.URL, "https://", "", -1)
	Convey("When listing a Supergiant Entrypoints components.", t, func() {
		//setup
		sg, err := NewClient(testURL, "test", "test")
		So(err, ShouldBeNil)

		Convey("We would expect the list to be successful", func() {
			_, err := sg.getEntrypointComps("test")
			So(err, ShouldBeNil)
		})
		Convey("We would expect the list to fail if there is an api error.", func() {
			ts.Close()
			_, err := sg.getEntrypointComps("test")
			So(err.Error(), ShouldContainSubstring, "Get https")
		})
	})
}
