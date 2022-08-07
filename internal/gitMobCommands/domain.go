// Package gitMobCommands contains methods ported from git-mob/src/git-mob-commands.js
package gitMobCommands

import (
	"fmt"
	"github.com/davidalpert/go-git-mob/internal/authors"
	"github.com/davidalpert/go-git-mob/internal/checkAuthor"
	"github.com/davidalpert/go-git-mob/internal/gitConfig"
	"strings"
)

// GetCoAuthors gets the current list of co-authors from git config
func GetCoAuthors() ([]authors.Author, error) {
	configValues, err := gitConfig.GetAllGlobal("git-mob.co-author")
	if err != nil {
		return nil, err
	}

	aa := make([]authors.Author, len(configValues))
	for i, o := range configValues {
		//fmt.Printf("found option: %s\n", o)
		if a, err := authors.ParseOne(o); err != nil {
			return nil, fmt.Errorf("failed to parse co-author from config option: '%s'", o)
		} else {
			aa[i] = a
		}
	}
	return aa, nil
}

func IsCoAuthorsSet() bool {
	return gitConfig.HasGlobal("git-mob.co-author")
}

func AddCoAuthor(a authors.Author) error {
	return gitConfig.AddGlobal("git-mob.co-author", a.String())
}

func AddCoAuthors(aa ...authors.Author) error {
	for _, a := range aa {
		if err := AddCoAuthor(a); err != nil {
			return err
		}
	}

	return nil
}

// ResetMob clears out the co-authors from global git config
func ResetMob() error {
	return gitConfig.RemoveAllGlobal("git-mob.co-author")
}

func UseLocalTemplate() bool {
	v := gitConfig.GetLocal("git-mob.use-local-template")
	return strings.EqualFold(v, "true")
}

// GetGitAuthor builds an authors.Author from the current configured
// user; returns an error if git config is missing user.name or user.email
func GetGitAuthor() (*authors.Author, error) {
	name := gitConfig.Get("user.name")
	email := gitConfig.Get("user.email")

	err := checkAuthor.ConfigWarning(name, email)
	return &authors.Author{
		Name:  name,
		Email: email,
	}, err
}

// SetGitAuthor sets the primary git author from the given authors.Author
func SetGitAuthor(a *authors.Author) error {
	if err := gitConfig.Set("user.name", a.Name); err != nil {
		return err
	}
	if err := gitConfig.Set("user.email", a.Email); err != nil {
		return err
	}
	return nil
}
