// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at 2023-03-10 07:51:03.650135 -0600 CST m=+0.036639917
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
		BuildDate:            "2023-03-10 07:51:03.650135 -0600 CST m=+0.036639917",
		CoreVersion:          "0.10.0",
		GitBranch:            "main",
		GitCommit:            "851165e",
		GitCommitSummary:     "fix: build the correct cmd/git-mob package on release",
		GitDirty:             false,
		GitDirtyHasModified:  false,
		GitDirtyHasStaged:    false,
		GitDirtyHasUntracked: false,
		Version:              "0.10.0+851165e",
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
