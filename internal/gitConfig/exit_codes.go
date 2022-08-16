package gitConfig

import "fmt"

// re: https://git-scm.com/docs/git-config#_description
// [the `git config`] command will fail with non-zero status upon error. Some exit codes are...

// ExitCode of the `git config` command.
type ExitCode int

// ExitCodes
const (
	UnknownExitCode ExitCode = iota - 1
	CommandSuccess           // 0
	SectionOrKeyIsInvalid
	SectionOrNameNotProvided
	FileIsInvalid
	FileCannotBeWritten
	CannotUnsetOptionWhichDoesNotExist
	CannotUnsetSetMultipleLinesMatched
	InvalidRegexp
)

var exitCodeNames = [...]string{
	CommandSuccess:                     "On success, the command returns the exit code 0.",
	SectionOrKeyIsInvalid:              "The section or key is invalid (ret = 1)",
	SectionOrNameNotProvided:           "no section or name was provided (ret = 2)",
	FileIsInvalid:                      "the config file is invalid (ret = 3)",
	FileCannotBeWritten:                "the config file cannot be written (ret = 4)",
	CannotUnsetOptionWhichDoesNotExist: "you try to unset an option which does not exist (ret = 5)",
	CannotUnsetSetMultipleLinesMatched: "you try to unset/set an option for which multiple lines match (ret = 5)",
	InvalidRegexp:                      "you try to use an invalid regexp (ret = 6)",
}

func (c ExitCode) String() string {
	if c.IsKnown() {
		return exitCodeNames[c]
	}

	if c == UnknownExitCode {
		return "unknown"
	}

	return "undocumented"
}

func (c ExitCode) IsKnown() bool {
	return UnknownExitCode < c && c <= InvalidRegexp
}

// Errorf wraps an error with the known interpretation from git
func (c ExitCode) Errorf(err error) error {
	if err == nil {
		return nil
	}

	if c.IsKnown() {
		return fmt.Errorf("%s: %v", c, err)
	}
	return fmt.Errorf("%s: %v", UnknownExitCode, err)
}
