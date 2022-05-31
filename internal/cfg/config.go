package cfg

import (
	"fmt"
	"github.com/davidalpert/go-git-mob/internal/env"
	"github.com/go-git/go-git/v5/config"
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

var (
	reCoauthor *regexp.Regexp
)

func init() {
	reCoauthor = regexp.MustCompile(`(?P<Name>[^<]+)\s+\<(?P<Email>[^>]+)\>`)
}
