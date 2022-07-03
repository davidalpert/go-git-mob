package revParse

import (
	"fmt"
	"github.com/davidalpert/go-git-mob/internal/shell"
	"path"
	"strings"
)

// InsideWorkTree checks if the current working directory is inside the working tree of a git repository.
// returns true if the cwd in a git repository.
func InsideWorkTree() bool {
	_, exitCode, err := shell.SilentRun("git", "rev-parse", "--is-inside-work-tree")

	return err == nil && exitCode == 0
}

// TopLevelDirectory computes the path to the top-level directory of the git repository.
func TopLevelDirectory() (string, error) {
	output, statusCode, err := shell.SilentRun("git", "rev-parse", "--show-toplevel")
	if err != nil {
		return "", err
	}
	if statusCode != 0 {
		return "", fmt.Errorf("TopLevelDirectory: expected 0 but got %d", statusCode)
	}
	return strings.TrimSpace(output), nil
}

// GitPath resolves the given path to the .git directory (GIT_DIR).
// from https://git-scm.com/book/en/v2/Git-Internals-Environment-Variables#_repository_locations
// GIT_DIR is the location of the .git folder. If this isn’t specified,
// Git walks up the directory tree until it gets to ~ or /, looking for a
// .git directory at every step.
func GitPath(rel ...string) (string, error) {
	tld, err := TopLevelDirectory()
	if err != nil {
		return "", err
	}
	return path.Join(append([]string{tld, ".git"}, rel...)...), nil
}

func GitPathRelativeToTopLevelDirectory(rel ...string) string {
	return path.Join(append([]string{".git"}, rel...)...)
}
