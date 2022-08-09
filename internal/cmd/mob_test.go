package cmd

import (
	"github.com/davidalpert/go-git-mob/internal/authors"
	"github.com/davidalpert/go-printers/v1"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestMobOptions_Run(t *testing.T) {
	type fields struct {
	}
	tests := []struct {
		name                   string
		Initials               []string
		ListOnly               bool
		PrintMob               bool
		PrintVersion           bool
		CurrentGitUser         authors.Author
		AllCoAuthorsByInitials map[string]authors.Author
		wantStdout             string
		wantStderr             string
		wantErr                bool
	}{
		{
			name:           "Scenario 01 - list",
			Initials:       []string{"ad", "bd"},
			ListOnly:       true,
			PrintMob:       false,
			PrintVersion:   false,
			CurrentGitUser: authors.Author{Name: "Jane Doe", Email: "jane@example.com"},
			AllCoAuthorsByInitials: map[string]authors.Author{
				"bd": authors.Author{Name: "Bob Doe", Email: "bob@example.com"},
				"ad": authors.Author{Name: "Amy Doe", Email: "amy@example.com"},
			},
			wantStdout: strings.TrimLeft(`
ad Amy Doe amy@example.com
bd Bob Doe bob@example.com
`, "\r\n"),
			wantErr: false,
		},
		// TODO: this test doesn't work because the newer implementation of cmd/mob.go couples Run to the external git configuration
		//		{
		//			name:           "Scenario 01 - list",
		//			Initials:       []string{"ad", "bd"},
		//			ListOnly:       false,
		//			PrintMob:       true,
		//			PrintVersion:   false,
		//			CurrentGitUser: authors.Author{Name: "Jane Doe", Email: "jane@example.com"},
		//			AllCoAuthorsByInitials: map[string]authors.Author{
		//				"bd": authors.Author{Name: "Bob Doe", Email: "bob@example.com"},
		//				"ad": authors.Author{Name: "Amy Doe", Email: "amy@example.com"},
		//			},
		//			wantStdout: strings.TrimLeft(`
		//Warning: git-mob uses global git config.
		//Using local commit.template could mean your template does not have selected co-authors appended after switching projects
		//Jane Doe <jane@example.com>
		//Amy Doe <amy@example.com>
		//Bob Doe <bob@example.com>
		//`, "\r\n"),
		//			wantErr: false,
		//		},
	}
	for _, tt := range tests {
		testStreams, _, stdOut, _ := printers.NewTestIOStreams()
		t.Run(tt.name, func(t *testing.T) {
			o := &MobOptions{
				PrinterOptions:         printers.NewPrinterOptions().WithStreams(testStreams),
				Initials:               tt.Initials,
				ListOnly:               tt.ListOnly,
				PrintMob:               tt.PrintMob,
				PrintVersion:           tt.PrintVersion,
				CurrentGitUser:         &tt.CurrentGitUser,
				AllCoAuthorsByInitials: tt.AllCoAuthorsByInitials,
			}
			if err := o.Run(); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}

			assert.Equal(t, tt.wantStdout, stdOut.String())
		})
	}
}
