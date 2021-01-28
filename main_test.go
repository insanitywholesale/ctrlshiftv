package main

import (
	"testing"
	"ctrlshiftv/paste"
)

// Test is still incomplete
// Should be integration test maybe
func TestGet(t *testing.T) {
	// initialize mock repo
	repo := chooseRepo()
	if repo == nil {
		t.Error("repo oopsie")
	}
	// make redirect service
	service := paste.NewPasteService(repo)
	t.Log("service", service)
/*
	// create router based on the above service
	r := makeRouter(service)
	// create and start a test server
	testServer := httptest.NewServer(r)
	// do a simple Get request on preexisting redirect
	// said redirect can be found in the mock repo source
	res, err := http.Get(testServer.URL + "/1234")
	// be responsible and close the response body
	res.Body.Close()
	if err != nil {
		t.Error("tfw GET error:", err)
	}
	// close the test server
	testServer.Close()
*/
}
