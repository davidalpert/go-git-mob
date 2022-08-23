// Package gitMessage
package gitMessage

import (
	"fmt"
	"github.com/davidalpert/go-git-mob/internal/authors"
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

func WriteGitMessage(coAuthorList ...authors.Author) error {
	return Write(Path(), coAuthorList...)
}

func Write(p string, coAuthorList ...authors.Author) error {
	content := "\n" + "\n" + FormatCoAuthorList(coAuthorList)

	if _, err := os.Stat(p); err == nil {
		b, err := os.ReadFile(p)
		if err != nil {
			return fmt.Errorf("reading git message file: %v", err)
		}

		content = replaceCoauthors(b, content)
	}

	return ioutil.WriteFile(p, []byte(content), os.ModePerm)
}

func replaceCoauthors(b []byte, content string) string {
	i := strings.Index(string(b), "\n\nCo-authored-by:")
	if i > 0 {
		content = string(b[:i]) + content
	} else {
		content = string(b) + content
	}
	return content
}
