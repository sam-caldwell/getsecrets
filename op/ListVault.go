package op

import "os/exec"

func ListVault() (string, error) {
	output, err := exec.Command(cmd1pwd, "vault", "list", "--format", "json").Output()
	return string(output), err
}
