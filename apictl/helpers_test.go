package apictl

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/supergiant/supergiant/common"
)

// Tests the commonErrorParse function
func TestCommonErrorParse(t *testing.T) {
	Convey("When parsing common errors.", t, func() {
		Convey("We expect expected errors to be custom parsed.", func() {
			err := commonErrorParse(errors.New("Request failed with status 404 Not Found"), "test")
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "test: Was not found.")
		})
		Convey("We expect non-expected errors to not be parsed.", func() {
			err := commonErrorParse(errors.New("test"), "test")
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "test")
		})
	})
}

// Tests the checkForNilString function
func TestCheckForNilString(t *testing.T) {
	Convey("When checking for a nil string.", t, func() {
		var pointer *string
		Convey("We expect nil strings to return with 'noData'", func() {
			string := checkForNilString(pointer)
			So(string, ShouldEqual, "No Data")
		})
		Convey("We expect non-nil sring pointers to be passed through as string.", func() {
			s := "test"
			string := checkForNilString(&s)
			So(string, ShouldEqual, "test")
		})
	})
}

// Tests the checkforNILTime function
func TestCheckforNILTime(t *testing.T) {
	Convey("When checking for a nil time.", t, func() {
		var pointer *common.Timestamp
		Convey("We expect nil time strings to return with 'noData'", func() {
			string := checkforNILTime(pointer)
			So(string, ShouldEqual, "No Data")
		})
		Convey("We expect non-nil sring pointers to be passed through as string.", func() {
			s := &common.Timestamp{}
			string := checkforNILTime(s)
			So(string, ShouldEqual, "0001-01-01 00:00:00 +0000 UTC")
		})
	})
}

// Tests the checkCommonErrors function
func TestCheckCommonErrors(t *testing.T) {
	Convey("When checking misc common errors.", t, func() {
		Convey("We all common errors to not be nil.", func() {
			e := errors.New("test")
			err := clientConnectionError(e)
			So(err, ShouldNotBeNil)

			err = appGetError("test", e)
			So(err, ShouldNotBeNil)

			err = compGetError("test", e)
			So(err, ShouldNotBeNil)

			err = releaseGetError("test", e)
			So(err, ShouldNotBeNil)
		})

	})
}
