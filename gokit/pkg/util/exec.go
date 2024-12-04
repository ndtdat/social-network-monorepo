package util

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func Exec(commands []string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", strings.Join(commands, ";")) //nolint:gosec
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf(err.Error() + ": " + stderr.String())
	}

	return out.String(), nil
}
