package bash

import (
	"bytes"
	"os/exec"
)

//Run run bash command
func Run(s string) (bytes.Buffer, error) {
	cmd := exec.Command("/bin/bash", "-c", s)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out, err
}
