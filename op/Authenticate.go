package op

import (
	"fmt"
	"os/exec"
)

func Authenticate() error {
	cmd := exec.Command(cmd1pwd, "signin")
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("error signing into 1Password CLI: %v\nstdout:%v", err, string(output))
	}
	return err
}
