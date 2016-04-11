package spacetime

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestProvider(t *testing.T) {

	// Test init
	os.Setenv("SG_CLI_TEST_MODE", "true")
	sgdir := "/tmp/.supergiant"
	pvdrconfig := "" + sgdir + "/provider.json"
	os.Remove(pvdrconfig)

	Convey("When creating a new provider.", t, func() {
		NewProvider("test", "test", "test", "test", false)
		Convey("The contents of the providers db should match the expect content.", func() {
			buf, err := ioutil.ReadFile(pvdrconfig)
			if err != nil {
				t.Error("Provider db read failed.")
			}
			actual := string(buf)

			expected := `{
  "providerConfigs": {
    "test": {
      "name": "test",
      "type": "test",
      "accessKey": "test",
      "secretKey": "test",
      "sshKey": "",
      "staged": "staged",
      "log": ""
    }
  }
}`
			So(actual, ShouldEqual, expected)
		})
	})

	Convey("When a user lists all providers with the CLI.", t, func() {
		// Capture STDOUT
		old := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w
		outC := make(chan string)

		go func() {
			var buf bytes.Buffer
			io.Copy(&buf, r)
			outC <- buf.String()
		}()

		ListProvider("")
		// back to normal state
		w.Close()
		os.Stdout = old // restoring the real stdout
		out := <-outC

		Convey("The user should get the expected content.", func() {
			expected := `ProviderType	Status 	Access Key	Secret Key	
test	test	staged 	test		test`

			So(strings.TrimSpace(out), ShouldEqual, expected)
		})
	})

	Convey("When a user lists a specific provider with the CLI.", t, func() {
		// Capture STDOUT
		old := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w
		outC := make(chan string)

		go func() {
			var buf bytes.Buffer
			io.Copy(&buf, r)
			outC <- buf.String()
		}()

		ListProvider("test")
		// back to normal state
		w.Close()
		os.Stdout = old // restoring the real stdout
		out := <-outC

		Convey("The user should get the expected content.", func() {
			expected := `Provider Details:
	Name: test
	  Type: test
		Status: staged
		AWS Access Key: test
		AWS Secret Key: test
		Log File:`

			So(strings.TrimSpace(out), ShouldEqual, expected)
		})
	})

	os.Setenv("SG_CLI_TEST_MODE", "false")
}
