/*
Copyright 2022 The kubeall.com Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	TestHarborAddress  = ""
	TestHarborUsername = ""
	TestHarborPwd      = ""
)

func testInit() {
	if d, e := ioutil.ReadFile("test.json"); e == nil {
		var m map[string]string
		if json.Unmarshal(d, &m) == nil {

			if k, o := m["TestHarborAddress"]; o {
				TestHarborAddress = k
			}
			if k, o := m["TestHarborUsername"]; o {
				TestHarborUsername = k
			}
			if k, o := m["TestHarborPwd"]; o {
				TestHarborPwd = k
			}
		}
	}
}

// setup sets up a test HTTP server along with a harbor.Client that is
// configured to talk to that test server.  Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup(t *testing.T) (*http.ServeMux, *httptest.Server, *Client) {
	// mux is the HTTP request multiplexer used with the test server.
	mux := http.NewServeMux()

	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(mux)

	// client is the Gitlab client being tested.
	client, err := NewClient("", "", "")
	if err != nil {
		server.Close()
		t.Fatalf("Failed to create client: %v", err)
	}

	return mux, server, client
}

// teardown closes the test HTTP server.
func teardown(server *httptest.Server) {
	server.Close()
}
