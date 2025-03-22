package op

import (
	"os/exec"
)

func GetItem(vaultName, itemName string) (result string, err error) {
	var stdout []byte
	stdout, err = exec.Command(cmd1pwd, "item", "get", "--vault", vaultName, itemName, "--reveal", "--format", "json").Output()
	if err != nil {
		return "", err
	}
	return string(stdout), err
}
