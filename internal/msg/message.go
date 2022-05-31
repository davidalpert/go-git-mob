package msg

import (
	"fmt"
	"github.com/davidalpert/go-git-mob/internal/cfg"
	"strings"
)

// FormatCoAuthorList takes a list of authors in the form "name <email>" and
// formats them as a newline-separated list of Co-authored-by tags
func FormatCoAuthorList(coAuthorList []cfg.Author) string {
	tags := make([]string, len(coAuthorList))
	for i, a := range coAuthorList {
		tags[i] = fmt.Sprintf("Co-authored-by: %s <%s>", a.Name, a.Email)
	}
	return strings.Join(tags, "\n")
}
