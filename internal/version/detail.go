// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at 2022-08-06 10:26:33.611772 -0500 CDT m=+0.053536751
package version

import (
	"fmt"
	"os/user"
	"runtime"
	"sort"
	"strings"
)

// Detail provides an easy global way to
var Detail = NewVersionDetail()

// NewVersionDetail builds a new version DetailStruct
func NewVersionDetail() DetailStruct {
	s := DetailStruct{
		AppName:              "git-mob",
		BuildDate:            "2022-08-06 10:26:33.611772 -0500 CDT m=+0.053536751",
        CoreVersion:          "0.5.0",
		GitBranch:            "main",
		GitCommit:            "9cfd2d8",
		GitCommitSummary:     "Merge pull request #48 from davidalpert/GH-43-fix-packaged-version",
		GitDirty:             true,
		GitDirtyHasModified:  true,
		GitDirtyHasStaged:    false,
		GitDirtyHasUntracked: false,
		Version:              "0.5.0+9cfd2d8.dirty",
	}
	s.UserAgentString = s.ToUserAgentString()
	if s.GitDirty {
		s.GitWorkingState = "dirty"
	}
	return s
}

// DetailStruct provides an easy way to grab all the govvv version details together
type DetailStruct struct {
	AppName              string `json:"app_name"`
	BuildDate            string `json:"build_date"`
	CoreVersion          string `json:"core_version"`
	GitBranch            string `json:"branch"`
	GitCommit            string `json:"commit"`
	GitCommitSummary     string `json:"commit_summary"`
	GitDirty             bool `json:"dirty"`
	GitDirtyHasModified  bool `json:"dirty_modified"`
	GitDirtyHasStaged    bool `json:"dirty_staged"`
	GitDirtyHasUntracked bool `json:"dirty_untracked"`
	GitWorkingState      string `json:"working_state"`
	GitSummary           string `json:"summary"`
	UserAgentString      string `json:"user_agent"`
	Version              string `json:"version"`
}

// String implements Stringer
func (d *DetailStruct) String() string {
	if d == nil {
		return "n/a"
	}
	return fmt.Sprintf("%s %s", d.AppName, d.Version)
}

// ToUserAgentString formats a DetailStruct as a User-Agent string
func (s DetailStruct) ToUserAgentString() string {
	productName := s.AppName
	productVersion := s.Version

	productDetails := map[string]string{ }

	user, err := user.Current()
	if err == nil {
		username := user.Username
		if username == "" {
			username = "unknown"
		}

		productDetails["user"] = username // strings.Replace(user.Username, "a-", 1) // this is a northfield convention
	}

	detailParts := []string{}
	for k, v := range productDetails {
		detailParts = append(detailParts, fmt.Sprintf("%s: %s", k, v))
	}
	sort.Slice(detailParts, func(i, j int) bool {
		return detailParts[i] < detailParts[j]
	})
	productDetail := strings.Join(detailParts, ", ")

	return fmt.Sprintf("%s/%s (%s) %s (%s)", productName, productVersion, productDetail, runtime.GOOS, runtime.GOARCH)
}
