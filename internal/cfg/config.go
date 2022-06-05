package cfg

import (
	"bytes"
	"fmt"
	"github.com/davidalpert/go-git-mob/internal/authors"
	"github.com/davidalpert/go-git-mob/internal/env"
	"github.com/go-git/go-git/v5/config"
	"os/exec"
	"path"
	"strings"
)

// Get gets the (last) value for the given option key.
func Get(key string) string {
	o, err := silentRun("git", "config", "--get", key)
	if err != nil {
		return ""
	}
	return o
}

// GetAll gets all values for a multi-valued option key.
func GetAll(key string) ([]string, error) {
	o, err := silentRun("git", "config", "--all", key)
	if err != nil {
		return make([]string, 0), err
	}
	return strings.Split(o, "\n"), nil
}

// Set sets the option, overwriting the existing value if one exists.
func Set(key string, value string) error {
	//const { status } = silentRun(`git config ${key} "${value}"`);
	_, err := silentRun("git", "config", key, value)
	if err != nil {
		return fmt.Errorf("option '%s' has multiple values. Cannot overwrite multiple values for option '%s' with a single value", key, key)
	}
	return nil
}

// SetGlobal sets the global option, overwriting the existing value if one exists.
func SetGlobal(key string, value string) error {
	//const { status } = silentRun(`git config ${key} "${value}"`);
	_, err := silentRun("git", "config", "--global", key, value)
	if err != nil {
		return fmt.Errorf("option '%s' has multiple values. Cannot overwrite multiple values for option '%s' with a single value", key, key)
	}
	return nil
}

// Add adds a new line to the option without altering any existing values.
func Add(key string, value string) error {
	_, err := silentRun("git", "config", "--add", key, value)
	return err
}

// AddGlobal adds a new line to the global option without altering any existing values.
func AddGlobal(key string, value string) error {
	_, err := silentRun("git", "config", "--global", "--add", key, value)
	return err
}

// silentRun runs the given command in a shell.
func silentRun(name string, arg ...string) (string, error) {
	c := exec.Command(name, arg...)
	//c.Stdin = strings.NewReader("and old falcon")

	if env.GetValueOrDefaultBool("GITMOB_DEBUG", false) {
		fmt.Printf("silentRun: %s %s\n", name, strings.Join(arg, " "))
	}

	var out bytes.Buffer
	c.Stdout = &out
	var stdErr bytes.Buffer
	c.Stderr = &stdErr

	if err := c.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return "", fmt.Errorf("nonzero exit code: %d: %s %s", exitError.ExitCode(), out.String(), stdErr.String())
		}
		return "", err
	}

	if env.GetValueOrDefaultBool("GITMOB_DEBUG", false) {
		fmt.Println(out.String())
	}

	return strings.TrimSpace(out.String()), nil
}

// Has checks if the given option exists in the configuration.
func Has(key string) bool {
	_, err := silentRun("git", "config", key)
	return err == nil
}

// HasGlobal checks if the given option exists in the global configuration.
func HasGlobal(key string) bool {
	_, err := silentRun("git", "config", "--global", key)
	return err == nil
}

// RemoveSection removes the given section from the configuration.
func RemoveSection(key string) error {
	if Has(key) {
		_, err := silentRun("git", "config", "--remove-section", key)
		return err
	}
	return nil
}

// RemoveSectionGlobal removes the given section from the global configuration.
func RemoveSectionGlobal(key string) error {
	if HasGlobal(key) {
		_, err := silentRun("git", "config", "--global", "--remove-section", key)
		return err
	}
	return nil
}

// Remove removes the given key from the configuration.
func Remove(key string) error {
	if Has(key) {
		_, err := silentRun("git", "config", "--unset", key)
		return err
	}
	return nil
}

// RemoveAll removes all the given keys from the configuration.
func RemoveAll(key string) error {
	if Has(key) {
		_, err := silentRun("git", "config", "--unset-all", key)
		return err
	}
	return nil
}

// RemoveGlobal removes the given key from the configuration.
func RemoveGlobal(key string) error {
	if HasGlobal(key) {
		_, err := silentRun("git", "config", "--global", "--unset", key)
		return err
	}
	return nil
}

// RemoveAllGlobal removes all the given keys from the configuration.
func RemoveAllGlobal(key string) error {
	if HasGlobal(key) {
		_, err := silentRun("git", "config", "--global", "--unset-all", key)
		return err
	}
	return nil
}

func AddCoAuthors(aa ...authors.Author) error {
	for _, a := range aa {
		if err := AddGlobal("git-mob.co-author", a.String()); err != nil {
			return err
		}
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

func ShortLogAuthorSummary() (map[string]authors.Author, error) {
	// git shortlog --summary --email --number HEAD'
	o, err := silentRun("git", "shortlog", "--summary", "--email", "--number", "HEAD")
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

var (
	GlobalConfigFilePath string
)

const (
	EnvKeyCoauthorsPath = "GITMOB_COAUTHORS_PATH"
)

func init() {
	GlobalConfigFilePath = env.GetValueOrDefault(EnvKeyCoauthorsPath, path.Join(env.HomeDir, ".gitconfig"))
}
