package spacetime

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestListKubes(t *testing.T) {
	os.Setenv("SG_CLI_TEST_MODE", "true")
	// ensure test db is destroyed
	sgdir := "/tmp/.supergiant"
	sgconfig := "" + sgdir + "/sgconfig.json"
	os.Remove(sgconfig)
	pvdrconfig := "" + sgdir + "/provider.json"
	os.Remove(pvdrconfig)

	NewProvider("test", "test", "test", "test", false)
	NewKube("test", "test", "test", "test", "test", "test", "test", false)

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	outC := make(chan string)

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	ListKubes("")

	// back to normal state
	w.Close()
	os.Stdout = old // restoring the real stdout
	out := <-outC

	expected := `Name	Provider 	AZ	IP	Status	Core Status	Managed	Selected
test	test 		test		buildingfalse		true	true`

	result := strings.TrimSpace(out)
	if result != expected {
		t.Error("ERROR ListKubes: expected,\n", expected, "\n -- But got,\n", out)
	}
}
