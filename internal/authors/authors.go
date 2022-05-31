package authors

import (
	"encoding/json"
	"fmt"
	"github.com/davidalpert/go-git-mob/internal/env"
	"os"
	"path"
	"regexp"
	"strings"
)

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (a Author) String() string {
	return fmt.Sprintf("%s <%s>", a.Name, a.Email)
}

func (a Author) CoauthorTag() string {
	return fmt.Sprintf("Co-Authored-By: %s <%s>", a.Name, a.Email)
}

// MustParseOne parses an author string into an Author and panics if parsing fails
func MustParseOne(s string) Author {
	if a, err := ParseOne(s); err != nil {
		panic(err)
	} else {
		return a
	}
}

// ParseOne parses an author string into an Author
func ParseOne(s string) (Author, error) {
	res := reAuthorString.FindAllStringSubmatch(s, 1)

	if len(res) > 0 {
		return Author{
			Name:  res[0][1],
			Email: res[0][2],
		}, nil
	}

	return Author{}, fmt.Errorf("failed to parse co-author from config option: '%s'", s)
}

type CoAuthorsFileContent struct {
	CoAuthorsByInitial map[string]Author `json:"coauthors"`
}

func ReadCoAuthorsContent() (CoAuthorsFileContent, error) {
	return ReadCoAuthorsContentFromFile(CoAuthorsFilePath)
}

func ReadCoAuthorsContentFromFile(filePath string) (CoAuthorsFileContent, error) {
	b, err := os.ReadFile(filePath)
	if err != nil {
		return CoAuthorsFileContent{}, err
	}
	return ReadCoAuthorsContentFromBytes(b)
}

func ReadCoAuthorsContentFromBytes(b []byte) (CoAuthorsFileContent, error) {
	var c CoAuthorsFileContent
	err := json.Unmarshal(b, &c)
	return c, err
}

func (f CoAuthorsFileContent) LookupByInitials(initials ...string) []Author {
	parts := make([]Author, 0)
	for _, i := range initials {
		for initial, author := range f.CoAuthorsByInitial {
			if strings.EqualFold(i, initial) {
				parts = append(parts, author)
				continue
			}
		}
	}

	return parts
}

func (f CoAuthorsFileContent) FindAndFormatAsList(initials ...string) []string {
	aa := f.LookupByInitials(initials...)
	result := make([]string, len(aa))

	for i, a := range aa {
		result[i] = a.String()
	}

	return result
}

func (f CoAuthorsFileContent) FindInitialsFromCoAuthorStrings(ss ...string) []string {
	result := make([]string, 0)

	for _, s := range ss {
		a := MustParseOne(s)
		for k, v := range f.CoAuthorsByInitial {
			if strings.EqualFold(a.Email, v.Email) {
				result = append(result, k)
			}
		}
	}

	return result
}

func EnsureCoauthorsFileExists() error {
	if _, err := os.Stat(CoAuthorsFilePath); os.IsNotExist(err) {
		cc := CoAuthorsFileContent{
			CoAuthorsByInitial: make(map[string]Author, 0),
		}
		b, err := json.Marshal(cc)
		if err != nil {
			return err
		}

		return os.WriteFile(CoAuthorsFilePath, b, os.ModePerm)
	}
	return nil
}

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
