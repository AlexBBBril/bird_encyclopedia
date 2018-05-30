package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
)

func TestRouter(t *testing.T) {
	r := newRouter()

	// Documentation : https://golang.org/pkg/net/http/httptest/#NewServer
	mockServer := httptest.NewServer(r)
	resp, err := http.Get(mockServer.URL + "/hello")

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be ok, got %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	respString := string(b)
	expected := "Hello World!"

	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}
}
