package sc_executor

import (
	"os"
	"os/exec"
)

func scriptExec(filePath string) (string, error) {
	scriptContent, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	cmd := exec.Command("bash", "-c", string(scriptContent))

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	outputStr := string(output)

	return outputStr, nil
}
