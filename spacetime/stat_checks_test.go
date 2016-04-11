package spacetime

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	kubeResp = `{
	  "kind": "NamespaceList",
	  "apiVersion": "v1",
	  "metadata": {
	    "selfLink": "/api/v1/namespaces",
	    "resourceVersion": "test"
	  },
	  "items": [
	    {
	      "metadata": {
	        "name": "default",
	        "selfLink": "/api/v1/namespaces/default",
	        "uid": "test",
	        "resourceVersion": "test",
	        "creationTimestamp": "test"
	      },
	      "spec": {
	        "finalizers": [
	          "kubernetes"
	        ]
	      },
	      "status": {
	        "phase": "Active"
	      }
	    },
	    {
	      "metadata": {
	        "name": "kube-system",
	        "selfLink": "/api/v1/namespaces/kube-system",
	        "uid": "test",
	        "resourceVersion": "test",
	        "creationTimestamp": "test"
	      },
	      "spec": {
	        "finalizers": [
	          "kubernetes"
	        ]
	      },
	      "status": {
	        "phase": "Active"
	      }
	    }
	  ]
	}`
)

// Check that it works
func TestCheckForLife(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, kubeResp)
	}))
	defer ts.Close()

	url := strings.Replace(ts.URL, "https://", "", -1)
	err := checkForLife(url, "", "")
	if err != nil {
		t.Error("ERROR checkforlife: Expected to get a successful check. But check failed.")
	}
}

// Check what happens if the request fails.
func TestCheckForLifeClientRequestFail(t *testing.T) {
	err := checkForLife("test", "", "")
	if err == nil {
		t.Error("ERROR checkforlife: Expected to get a failed check. But check passed.")
	}
}

// Check that it works, but does not give us the expected response
func TestCheckForLifeUnexpectedResponse(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Donkey")
	}))
	defer ts.Close()

	url := strings.Replace(ts.URL, "https://", "", -1)
	err := checkForLife(url, "", "")
	if err.Error() != "Kubernetes appears to be up, but the kube-system namespace is not there." {
		t.Error("ERROR checkforlife: Expected error msg, \"Kubernetes appears to be up, but the kube-system namespace is not there.\". But got,\"", err.Error(), "\"")
	}
}
