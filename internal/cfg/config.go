package cfg

import (
	"fmt"
	"github.com/davidalpert/go-git-mob/internal/authors"
	"github.com/davidalpert/go-git-mob/internal/env"
	"github.com/go-git/go-git/v5/config"
	"os"
	"path"
)

// Get gets the (last) value for the given option key.
func Get(key string) (string, error) {
	return "", nil
}

// GetAll gets all values for a multi-valued option key.
func GetAll(key string) ([]string, error) {
	return nil, nil
}

// ResetMob clears out the co-authors from global git config
func ResetMob() error {
	c, err := config.LoadConfig(config.GlobalScope)
	if err != nil {
		return err
	}

	if c.Raw.HasSection("git-mob") {
		s := c.Raw.Section("git-mob")
		s.RemoveOption("co-author")
		return writeConfig(c)
	}

	return nil
}

// writeConfig saves the in-memory git config back to the global gitconfig file
func writeConfig(c *config.Config) error {
	b, err := c.Marshal()
	if err != nil {
		return err
	}

	return os.WriteFile(GlobalConfigFilePath, b, os.ModePerm)
}

func AddCoAuthors(aa ...authors.Author) error {
	c, err := config.LoadConfig(config.GlobalScope)
	if err != nil {
		return err
	}

	for _, a := range aa {
		c.Raw.AddOption("git-mob", "", "co-author", fmt.Sprintf("%s <%s>", a.Name, a.Email))
	}

	if len(aa) > 0 {
		return writeConfig(c)
	}

	return nil
}

// GetUser builds an authors.Author from the current configured user
func GetUser() (*authors.Author, error) {
	c, err := config.LoadConfig(config.GlobalScope)
	if err != nil {
		return nil, err
	}

	return &authors.Author{
		Name:  c.User.Name,
		Email: c.User.Email,
	}, nil
}

// GetCoAuthors gets the current list of co-authors from git config
func GetCoAuthors() ([]authors.Author, error) {
	//fmt.Printf("GetCoAuthors\n")
	c, err := config.LoadConfig(config.GlobalScope)
	if err != nil {
		return nil, err
	}

	if c.Raw.HasSection("git-mob") {
		oo := c.Raw.Section("git-mob").OptionAll("co-author")
		aa := make([]authors.Author, len(oo))
		for i, o := range oo {
			//fmt.Printf("found option: %s\n", o)
			if a, err := authors.ParseOne(o); err != nil {
				return nil, fmt.Errorf("failed to parse co-author from config option: '%s'", o)
			} else {
				aa[i] = a
			}
		}
		return aa, nil
	}

	return nil, nil
}

func SetCoAuthors() error {
	return nil
}

func ReadAllCoAuthorsFromFile() (map[string]authors.Author, error) {
	c, e := authors.ReadCoAuthorsContent()
	if e != nil {
		return nil, e
	}

	return c.CoAuthorsByInitial, nil
}

var (
	GlobalConfigFilePath string
)

const (
	EnvKeyCoauthorsPath = "GITMOB_COAUTHORS_PATH"
)

func init() {
	GlobalConfigFilePath = env.GetValueOrDefault(EnvKeyCoauthorsPath, path.Join(env.HomeDir, ".gitconfig"))
}
