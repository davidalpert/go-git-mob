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

	var out bytes.Buffer
	c.Stdout = &out
	var stdErr bytes.Buffer
	c.Stderr = &stdErr

	lg := diagnostics.Log.WithFields(log.Fields{
		"method": "SilentRun",
		"cmd":    fmt.Sprintf("%s %s", name, strings.Join(arg, " ")),
	})

	if err := c.Run(); err != nil {
		lg = lg.WithFields(log.Fields{
			"stdout": out.String(),
			"stderr": stdErr.String(),
		})
		if exitError, ok := err.(*exec.ExitError); ok {
			lg.WithFields(log.Fields{
				"exit.code":   exitError.ExitCode(),
				"exit.error":  exitError.Error(),
				"exit.stderr": string(exitError.Stderr),
			}).WithError(err).Error("command failed")
			return "", exitError.ExitCode(), fmt.Errorf("nonzero exit code: %d: %s", exitError.ExitCode(), exitError.Error())
		}
		lg.WithError(err).Error("command failed without exitError")
		return "", -1, fmt.Errorf("%s;%s", stdErr.String(), out.String())
	}

	lg.WithFields(log.Fields{
		"stdout": out.String(),
		"stderr": stdErr.String(),
	}).Debug("command succeeded")

	return strings.TrimSpace(out.String()), c.ProcessState.ExitCode(), nil
}
