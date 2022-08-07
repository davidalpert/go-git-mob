package checkAuthor

import "fmt"

// ConfigWarning validates that we did not receive an error loading primary git details;
// ported from git-mob/src/check-author.js
func ConfigWarning(name string, email string) error {
	errMsg := "warning: Missing information for the primary author. Set with:\n"
	var missingConfig = false
	if name == "" {
		errMsg = errMsg + "\n$ git config --global user.name \"Jane Doe\""
		missingConfig = true
	}

	if email == "" {
		errMsg = errMsg + "\n$ git config --global user.email \"jane@example.com\""
		missingConfig = true
	}

	if missingConfig {
		return fmt.Errorf(errMsg)
	}
	return nil
}
