package shell

import (
	"bytes"
	"fmt"
	"github.com/apex/log"
	"github.com/davidalpert/go-git-mob/internal/diagnostics"
	"os/exec"
	"strings"
)

// SilentRun runs the given command in a shell.
func SilentRun(name string, arg ...string) (string, int, error) {
	c := exec.Command(name, arg...)

	lg := diagnostics.Log.WithFields(log.Fields{
		"method": "SilentRun",
		"cmd":    fmt.Sprintf("%s %s", name, strings.Join(arg, " ")),
	})

	var out bytes.Buffer
	c.Stdout = &out
	var stdErr bytes.Buffer
	c.Stderr = &stdErr

	if err := c.Run(); err != nil {
		lg.WithError(err).Error("command failed")
		if exitError, ok := err.(*exec.ExitError); ok {
			return "", 0, fmt.Errorf("nonzero exit code: %d: %s\nexitError.Stderr: %s\ncmd.Stderr: %s\ncmd.Stdout: %s", exitError.ExitCode(), exitError.Error(), string(exitError.Stderr), stdErr.String(), out.String())
		}
		return "", 0, fmt.Errorf("%s;%s", stdErr.String(), out.String())
	}

	lg.WithFields(log.Fields{
		"stdout": out.String(),
		"stderr": stdErr.String(),
	}).Debug("command succeeded")

	return strings.TrimSpace(out.String()), c.ProcessState.ExitCode(), nil
}
