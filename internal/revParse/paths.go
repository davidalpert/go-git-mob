package revParse

import (
	"github.com/go-git/go-git/v5"
	"path"
)

// InsideWorkTree checks if the current working directory is inside the working tree of a git repository.
// returns true if the cwd in a git repository.
func InsideWorkTree() bool {
	_, err := git.PlainOpenWithOptions(".", &git.PlainOpenOptions{
		DetectDotGit:          true,
		EnableDotGitCommonDir: false,
	})
	if err == git.ErrRepositoryNotExists {
		return false
	}
	return true
}

// TopLevelDirectory computes the path to the top-level directory of the git repository.
func TopLevelDirectory() string {
	r, err := git.PlainOpenWithOptions(".", &git.PlainOpenOptions{
		DetectDotGit:          true,
		EnableDotGitCommonDir: false,
	})
	if err != nil {
		panic(err)
	}

	w, err := r.Worktree()
	if err != nil {
		panic(err)
	}

	return w.Filesystem.Root()
}

// GitPath resolves the given path to the .git directory (GIT_DIR).
// from https://git-scm.com/book/en/v2/Git-Internals-Environment-Variables#_repository_locations
// GIT_DIR is the location of the .git folder. If this isn’t specified,
// Git walks up the directory tree until it gets to ~ or /, looking for a
// .git directory at every step.
func GitPath(rel ...string) string {
	return path.Join(append([]string{TopLevelDirectory(), ".git"}, rel...)...)
}
