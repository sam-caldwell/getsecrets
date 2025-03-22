package op

import (
	"os/exec"
)

func ListItems(vaultName string) (result string, err error) {
	var stdout []byte
	stdout, err = exec.Command(cmd1pwd, "item", "list", "--vault", vaultName, "--format", "json").Output()
	if err != nil {
		return "", err
	}
	return string(stdout), err
}
