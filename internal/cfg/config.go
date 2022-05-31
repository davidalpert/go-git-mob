package cfg

import (
	"encoding/json"
	"fmt"
	"github.com/davidalpert/go-git-mob/internal/env"
	"github.com/go-git/go-git/v5/config"
	"io/ioutil"
	"os"
	"path"
	"regexp"
)

// Get gets the (last) value for the given option key.
func Get(key string) (string, error) {
	return "", nil
}

// GetAll gets all values for a multi-valued option key.
func GetAll(key string) ([]string, error) {
	return nil, nil
}

func GetCoAuthors() ([]Author, error) {
	//fmt.Printf("GetCoAuthors\n")
	c, err := config.LoadConfig(config.GlobalScope)
	if err != nil {
		return nil, nil
	}

	if c.Raw.HasSection("git-mob") {
		oo := c.Raw.Section("git-mob").OptionAll("co-author")
		aa := make([]Author, len(oo))
		for i, o := range oo {
			//fmt.Printf("found option: %s\n", o)
			res := reCoauthor.FindAllStringSubmatch(o, 1)
			if len(res) > 0 {
				aa[i] = Author{
					Name:  res[0][1],
					Email: res[0][2],
				}
			} else {
				return nil, fmt.Errorf("failed to parse co-author from config option: '%s'", o)
			}
		}
		return aa, nil
	}

	return nil, nil
}

func SetCoAuthors() error {
	return nil
}

func ReadAllCoAuthorsFromFile() (map[string]Author, error) {
	if err := EnsureCoauthorsFileExists(); err != nil {
		return nil, err
	}

	b, err := ioutil.ReadFile(CoAuthorsFilePath)
	if os.IsNotExist(err) {

	}

	var c CoAuthorsFileContent
	if err = json.Unmarshal(b, &c); err != nil {
		return nil, err
	}

	return c.CoAuthorsByInitial, nil
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
	reCoauthor        *regexp.Regexp
)

const (
	EnvKeyCoauthorsPath = "GITMOB_COAUTHORS_PATH"
)

func init() {
	CoAuthorsFilePath = env.GetValueOrDefault(EnvKeyCoauthorsPath, path.Join(env.HomeDir, ".git-coauthors"))
	reCoauthor = regexp.MustCompile(`(?P<Name>[^<]+)\s+\<(?P<Email>[^>]+)\>`)
}
