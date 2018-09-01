package status

import (
	"regexp"
	"testing"
)

func TestGetVersion(t *testing.T) {
	version := getVersion()
	match, err := regexp.MatchString(`^v[0-9]+\.[0-9]+\.[0-9]+$`, version)
	if !match || err != nil {
		t.Fail()
	}
}
