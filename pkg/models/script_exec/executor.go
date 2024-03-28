package sc_executor

import (
	"os/exec"

	"github.com/test/pkg/models"
)

func ScriptExec(s *models.Command) error {
	cmd := exec.Command("bash", "-c", s.Content)

	output, err := cmd.Output()
	if err != nil {
		return err
	}

	s.Exec_res = string(output)

	return nil
}
