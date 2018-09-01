package status

import (
	"os/exec"
	"strings"
)

func getVersion() string {
	out, err := exec.Command("git", "describe", "--tags").Output()
	if err != nil {
		return "v?.?.?"
	}
	version := string(out)
	return strings.TrimSpace(version)
}
