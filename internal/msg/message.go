package msg

import (
	"fmt"
	"github.com/davidalpert/go-git-mob/internal/authors"
	"github.com/davidalpert/go-git-mob/internal/cfg"
	"github.com/davidalpert/go-git-mob/internal/env"
	"github.com/davidalpert/go-git-mob/internal/revParse"
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

func GitMessagePath() (string, error) {
	p, err := revParse.GitPath(".gitmessage")
	if err != nil {
		return "", err
	}

	return env.GetValueOrDefault(EnvKeyGitMessagePath, p), nil
}

func CommitTemplatePath() (string, error) {
	s := env.GetValueOrDefaultString(EnvKeyGitMessagePath, cfg.Get("commit.template"))
	if s == "" {
		ss, err := GitMessagePath()
		if err != nil {
			return "", err
		}
		return ss, nil
	}
	return s, nil
}

func WriteGitMessage(coAuthorList ...authors.Author) error {
	p, err := GitMessagePath()
	if err != nil {
		return err
	}

	content := "\n" + "\n" + FormatCoAuthorList(coAuthorList)

	if _, err := os.Stat(p); err == nil {
		b, err := os.ReadFile(p)
		if err != nil {
			return fmt.Errorf("reading git message file: %v", err)
		}

		i := strings.Index(string(b), "\n\nCo-authored-by:")
		if i > 0 {
			content = string(b[:i]) + content
		} else {
			content = string(b) + content
		}
	}

	return ioutil.WriteFile(p, []byte(content), os.ModePerm)
}
