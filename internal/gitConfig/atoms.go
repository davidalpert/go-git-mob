// Package gitConfig provides wrappers around common git configuration commands
// as refactored from git-mob/src/git-commands.js
package gitConfig

import (
	"fmt"
	"github.com/davidalpert/go-git-mob/internal/shell"
	"strings"
)

// Get gets the (last) value for the given option key.
func Get(key string) string {
	o, _, err := shell.SilentRun("git", "config", "--get", key)
	if err != nil {
		return ""
	}
	return o
}

// GetLocal gets the (last) local value for the given option key.
func GetLocal(key string) string {
	o, _, err := shell.SilentRun("git", "config", "--local", "--get", key)
	if err != nil {
		return ""
	}
	return o
}

// GetGlobal gets the (last) global value for the given option key.
func GetGlobal(key string) string {
	o, _, err := shell.SilentRun("git", "config", "--global", "--get", key)
	if err != nil {
		return ""
	}
	return o
}

// GetAll gets all values for a multi-valued option key.
func GetAll(key string) ([]string, error) {
	o, exitCode, err := shell.SilentRun("git", "config", "--get-all", key)
	if err != nil {
		return make([]string, 0), ExitCode(exitCode).Errorf(err)
	}
	return strings.Split(o, "\n"), nil
}

// GetAllGlobal gets all values for a multi-valued option key.
func GetAllGlobal(key string) ([]string, error) {
	o, exitCode, err := shell.SilentRun("git", "config", "--global", "--get-all", key)
	if err != nil {
		return make([]string, 0), ExitCode(exitCode).Errorf(err)
	}
	return strings.Split(o, "\n"), nil
}

// Set sets the option, overwriting the existing value if one exists.
func Set(key string, value string) error {
	//const { status } = SilentRun(`git config ${key} "${value}"`);
	_, exitCode, err := shell.SilentRun("git", "config", key, value)
	if err != nil {
		return ExitCode(exitCode).Errorf(fmt.Errorf("set(%#v, %#v): %v", key, value, err))
	}
	return nil
}

// SetGlobal sets the global option, overwriting the existing value if one exists.
func SetGlobal(key string, value string) error {
	//const { status } = SilentRun(`git config ${key} "${value}"`);
	_, exitCode, err := shell.SilentRun("git", "config", "--global", key, value)
	if err != nil {
		return ExitCode(exitCode).Errorf(fmt.Errorf("option '%s' has multiple values. Cannot overwrite multiple values for option '%s' with a single value", key, key))
	}
	return nil
}

// Add adds a new line to the option without altering any existing values.
func Add(key string, value string) error {
	_, exitCode, err := shell.SilentRun("git", "config", "--add", key, value)
	return ExitCode(exitCode).Errorf(err)
}

// AddGlobal adds a new line to the global option without altering any existing values.
func AddGlobal(key string, value string) error {
	_, exitCode, err := shell.SilentRun("git", "config", "--global", "--add", key, value)
	return ExitCode(exitCode).Errorf(err)
}

// Has checks if the given option exists in the merged configuration.
func Has(key string) bool {
	_, _, err := shell.SilentRun("git", "config", key)
	return err == nil
}

// HasLocal checks if the given option exists in the local configuration.
func HasLocal(key string) bool {
	_, _, err := shell.SilentRun("git", "config", "--local", key)
	return err == nil
}

// HasGlobal checks if the given option exists in the global configuration.
func HasGlobal(key string) bool {
	_, _, err := shell.SilentRun("git", "config", "--global", key)
	return err == nil
}

// RemoveSection removes the given section from the configuration.
func RemoveSection(key string) error {
	if Has(key) {
		_, exitCode, err := shell.SilentRun("git", "config", "--remove-section", key)
		return ExitCode(exitCode).Errorf(err)
	}
	return nil
}

// RemoveSectionGlobal removes the given section from the global configuration.
func RemoveSectionGlobal(key string) error {
	if HasGlobal(key) {
		_, exitCode, err := shell.SilentRun("git", "config", "--global", "--remove-section", key)
		return ExitCode(exitCode).Errorf(err)
	}
	return nil
}

// Remove removes the given key from the configuration.
func Remove(key string) error {
	if Has(key) {
		_, exitCode, err := shell.SilentRun("git", "config", "--unset", key)
		return ExitCode(exitCode).Errorf(err)
	}
	return nil
}

// RemoveGlobal removes the given key from the configuration.
func RemoveGlobal(key string) error {
	if HasGlobal(key) {
		_, exitCode, err := shell.SilentRun("git", "config", "--global", "--unset", key)
		return ExitCode(exitCode).Errorf(err)
	}
	return nil
}

// RemoveAll removes all the given keys from the configuration.
func RemoveAll(key string) error {
	if Has(key) {
		_, exitCode, err := shell.SilentRun("git", "config", "--unset-all", key)
		return ExitCode(exitCode).Errorf(err)
	}
	return nil
}

// RemoveAllGlobal removes all the given keys from the configuration.
func RemoveAllGlobal(key string) error {
	if HasGlobal(key) {
		_, exitCode, err := shell.SilentRun("git", "config", "--global", "--unset-all", key)
		return ExitCode(exitCode).Errorf(err)
	}
	return nil
}
