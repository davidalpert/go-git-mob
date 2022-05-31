package msg

import (
	"fmt"
	"github.com/davidalpert/go-git-mob/internal/authors"
	"github.com/davidalpert/go-git-mob/internal/cfg"
	"github.com/davidalpert/go-git-mob/internal/env"
	"io/ioutil"
	"os"
	"strings"
)

// FormatCoAuthorList takes a list of authors in the form "name <email>" and
// formats them as a newline-separated list of Co-authored-by tags
func FormatCoAuthorList(coAuthorList []authors.Author) string {
	tags := make([]string, len(coAuthorList))
	for i, a := range coAuthorList {
		tags[i] = fmt.Sprintf("Co-authored-by: %s <%s>", a.Name, a.Email)
	}
	return strings.Join(tags, "\n")
}

const (
	EnvKeyGitMessagePath = "GITMOB_MESSAGE_PATH"
)

func GitMessagePath() string {
	return env.GetValueOrDefault(EnvKeyGitMessagePath, cfg.GitPath(".gitmessage"))
}

func CommitTemplatePath() string {
	s := env.GetValueOrDefaultString(EnvKeyGitMessagePath, cfg.Get("commit.template"))
	if s == "" {
		s = GitMessagePath()
	}
	return s
}

func WriteGitMessage(coAuthorList ...authors.Author) error {
	p := GitMessagePath()

	content := "\n" + "\n" + FormatCoAuthorList(coAuthorList)

	return ioutil.WriteFile(p, []byte(content), os.ModePerm)
}
