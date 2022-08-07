package authors

import (
	"github.com/davidalpert/go-git-mob/internal/env"
	"path"
	"regexp"
)

var (
	CoAuthorsFilePath string
	reAuthorString    *regexp.Regexp
)

const (
	EnvKeyCoauthorsPath = "GITMOB_COAUTHORS_PATH"
)

func init() {
	CoAuthorsFilePath = env.GetValueOrDefault(EnvKeyCoauthorsPath, path.Join(env.HomeDir, ".git-coauthors"))
	reAuthorString = regexp.MustCompile(`(?P<Name>[^<]+)\s+\<(?P<Email>[^>]+)\>`)
}
