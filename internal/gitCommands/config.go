// Package gitCommands maps to git-mob/src/git-commands.js
package gitCommands

import (
	"fmt"
	"github.com/davidalpert/go-git-mob/internal/authors"
	"github.com/davidalpert/go-git-mob/internal/gitConfig"
	"github.com/davidalpert/go-git-mob/internal/shell"
	"strings"
)

/*
  silentRun --> refParse.SilentRun(..)
  get --> gitConfig.Get(..)
  getAll --> gitConfig.GetAll(..)
  set --> gitConfig.Set(..)
  add --> gitConfig.Add(..)
  has --> gitConfig.Has(..)
  gitPath --> revParse.GitPath(..)
  insideWorkTree --> revParse.InsideWorkTre()
  topLevelDirectory --> revParse.TopLevelDirectory()
*/

// ShortLogAuthorSummary returns a list of existing authors of the git repository.
func ShortLogAuthorSummary() (map[string]authors.Author, error) {
	o, _, err := shell.SilentRun("git", "shortlog", "--summary", "--email", "--number", "HEAD")
	if err != nil {
		return nil, fmt.Errorf("error reading git shortlog: %v", err)
	}

	result := make(map[string]authors.Author, 0)
	for _, line := range strings.Split(o, "\n") {
		parts := strings.Split(line, "\t")
		if len(parts) > 3 {
			return nil, fmt.Errorf("shortlog line contained %d parts; expected 3: %#v", len(parts), parts)
		}
		s := strings.Join(parts[1:], " ")
		if a, err := authors.ParseOne(s); err != nil {
			return nil, fmt.Errorf("parsing '%s' as author: %v", s, err)
		} else {
			result[a.InitialsFromName()] = a
		}
	}
	return result, nil
}

func GetTemplatePath() string {
	return gitConfig.Get("commit.template")
}

func SetTemplatePath(path string) error {
	return gitConfig.SetGlobal("commit.template", path)
}

func HasTemplatePath() bool {
	return gitConfig.Has("commit.template")
}

func UsingLocalTemplate() bool {
	return gitConfig.HasLocal("commit.template")
}

func UsingGlobalTemplate() bool {
	return gitConfig.HasGlobal("commit.template")
}

func GetGlobalTemplate() string {
	return gitConfig.GetGlobal("commit.template")
}
