package sc_executor

import (
	"os/exec"
)

func scriptExec(scriptContent string) (string, error) {
	cmd := exec.Command("bash", "-c", scriptContent)

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	outputStr := string(output)

	return outputStr, nil
}
