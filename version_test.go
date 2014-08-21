package main

import (
	"strconv"
	"strings"
	"testing"
)

/**
 * Basically, verify current version follows SemVer guidelines
 */
func TestVersionSemVer(t *testing.T) {
	t.Logf("Testing version semantic (%s)", Version)
	details := strings.Split(Version, ".")
	if len(details) != 3 {
		t.Errorf("Version should provide major, minor and path informations: %s", Version)
	}
	if _, err := strconv.ParseInt(details[0], 2, 0); err != nil {
		t.Errorf(err.Error())
	}
	if _, err := strconv.ParseInt(details[1], 2, 0); err != nil {
		t.Errorf(err.Error())
	}

	patch := strings.Split(details[2], "-")
	if _, err := strconv.ParseInt(patch[0], 2, 0); err != nil {
		t.Errorf(err.Error())
	}
	if len(patch) > 2 {
		t.Error("last version part only provides patch number and pre-release info")

	}
}
