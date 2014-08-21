package main

import (
	"testing"
)

/**
 * Basically, verify current version follows SemVer guidelines
 */
func TestAPIVersion(t *testing.T) {
	t.Log("Testing api version: only v0 is supported")
	api := NewIntuitionAPI("http://localhost:5000")
	if api.Version != "v0" {
		t.Errorf("API Version not supported: %s", api.Version)
	}
}

func TestURLSetting(t *testing.T) {
	example_url := "http://localhost:5000"
	api := NewIntuitionAPI(example_url)
	base_url := api.Api.Api.BaseUrl.String()
	if base_url != example_url {
		t.Errorf("Unable to set api url: %s", base_url)
	}
}
