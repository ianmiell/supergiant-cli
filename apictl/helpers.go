package apictl

import (
	"errors"

	"github.com/supergiant/supergiant/common"
)

var (
	noData = "No Data"
)

func commonErrorParse(err error, objectName string) error {
	switch err.Error() {
	case "Request failed with status 404 Not Found":
		return errors.New("" + objectName + ": Was not found.")

	}
	return err
}

func checkForNilString(s *string) string {
	if s == nil {
		return noData
	}
	return *s
}

func checkforNILTime(t *common.Timestamp) string {
	if t == nil {
		return noData
	}
	return t.String()
}

func clientConnectionError(e error) error {
	return errors.New("Client connection error to supergiant core: " + e.Error() + "")
}

func appGetError(a string, e error) error {
	return errors.New("Failed to fetch app, " + a + ": " + e.Error() + "")
}

func compGetError(c string, e error) error {
	return errors.New("Failed to fetch Component, " + c + ": " + e.Error() + "")
}

func releaseGetError(c string, e error) error {
	return errors.New("Failed to fetch release info for component, " + c + ": " + e.Error() + "")
}
