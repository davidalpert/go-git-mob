package shell

import (
	"bytes"
	"fmt"
	"github.com/davidalpert/go-git-mob/internal/env"
	"os/exec"
	"strings"
)

// SilentRun runs the given command in a shell.
func SilentRun(name string, arg ...string) (string, int, error) {
	c := exec.Command(name, arg...)
	//c.Stdin = strings.NewReader("and old falcon")

	if env.GetValueOrDefaultBool("GITMOB_DEBUG", false) {
		fmt.Printf("SilentRun: %s %s\n", name, strings.Join(arg, " "))
	}

	var out bytes.Buffer
	c.Stdout = &out
	var stdErr bytes.Buffer
	c.Stderr = &stdErr

	if err := c.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return "", 0, fmt.Errorf("nonzero exit code: %d: %s %s", exitError.ExitCode(), out.String(), stdErr.String())
		}
		return "", 0, err
	}

	if env.GetValueOrDefaultBool("GITMOB_DEBUG", false) {
		fmt.Println(out.String())
	}

	return strings.TrimSpace(out.String()), c.ProcessState.ExitCode(), nil
}
