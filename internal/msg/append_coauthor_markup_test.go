package msg

import (
	"github.com/davidalpert/go-git-mob/internal/authors"
	"reflect"
	"testing"
)

func TestAppendCoauthorMarkup(t *testing.T) {
	tests := []struct {
		name        string
		haveAuthors []authors.Author
		haveMsg     string
		wantMsg     string
		wantErr     bool
	}{
		{
			name:        "empty message empty coauthors",
			haveMsg:     "",
			haveAuthors: []authors.Author{},
			wantMsg:     "",
			wantErr:     false,
		},
		{
			name:    "empty message one coauthor",
			haveMsg: "",
			haveAuthors: []authors.Author{
				{
					Name:  "Hoban Washburne",
					Email: "wash@serenity.com",
				},
			},
			wantMsg: `

Co-Authored-By: Hoban Washburne <wash@serenity.com>

`,
			wantErr: false,
		},
		{
			name:    "empty message two coauthors",
			haveMsg: "",
			haveAuthors: []authors.Author{
				{
					Name:  "Hoban Washburne",
					Email: "wash@serenity.com",
				},
				{
					Name:  "Zoe Washburne",
					Email: "zoe@serenity.com",
				},
			},
			wantMsg: `

Co-Authored-By: Hoban Washburne <wash@serenity.com>
Co-Authored-By: Zoe Washburne <zoe@serenity.com>

`,
			wantErr: false,
		},
		{
			name: "message with comments, two coauthors",
			haveMsg: `add something awesome

# Please enter the commit message for your changes. Lines starting
# with '#' will be ignored, and an empty message aborts the commit.
#
# On branch 23-feat-append-to-commit-message
# Your branch is up to date with 'origin/23-feat-append-to-commit-message'.
#
`,
			haveAuthors: []authors.Author{
				{
					Name:  "Hoban Washburne",
					Email: "wash@serenity.com",
				},
				{
					Name:  "Zoe Washburne",
					Email: "zoe@serenity.com",
				},
			},
			wantMsg: `add something awesome

Co-Authored-By: Hoban Washburne <wash@serenity.com>
Co-Authored-By: Zoe Washburne <zoe@serenity.com>

# Please enter the commit message for your changes. Lines starting
# with '#' will be ignored, and an empty message aborts the commit.
#
# On branch 23-feat-append-to-commit-message
# Your branch is up to date with 'origin/23-feat-append-to-commit-message'.
#
`,
			wantErr: false,
		},
		{
			name:        "existing coauthors empty coauthors",
			haveAuthors: []authors.Author{},
			haveMsg: `

Co-Authored-By: Hoban Washburne <wash@serenity.com>
Co-Authored-By: Zoe Washburne <zoe@serenity.com>
`,
			wantMsg: `

Co-Authored-By: Hoban Washburne <wash@serenity.com>
Co-Authored-By: Zoe Washburne <zoe@serenity.com>

`,
			wantErr: false,
		},
		{
			name: "existing coauthors one coauthor; existing authors preserved while new authors are added",
			haveAuthors: []authors.Author{
				{
					Name:  "Zoe Washburne",
					Email: "zoe@serenity.com",
				},
			},
			haveMsg: `

Co-Authored-By: Hoban Washburne <wash@serenity.com>
`,
			wantMsg: `

Co-Authored-By: Hoban Washburne <wash@serenity.com>
Co-Authored-By: Zoe Washburne <zoe@serenity.com>

`,
			wantErr: false,
		},
		{
			name: "existing coauthors one duplicate coauthor; duplicates are removed",
			haveAuthors: []authors.Author{
				{
					Name:  "Zoe Washburne",
					Email: "zoe@serenity.com",
				},
			},
			haveMsg: `

Co-Authored-By: Hoban Washburne <wash@serenity.com>
Co-Authored-By: Zoe Washburne <zoe@serenity.com>
`,
			wantMsg: `

Co-Authored-By: Hoban Washburne <wash@serenity.com>
Co-Authored-By: Zoe Washburne <zoe@serenity.com>

`,

			wantErr: false,
		},
		{
			name: "existing coauthors one duplicate coauthor; duplicates are removed",
			haveAuthors: []authors.Author{
				{
					Name:  "Bob Doe",
					Email: "bob@findmypast.com",
				},
			},
			haveMsg: `empty mobbed commit

Co-Authored-By: Amy Doe <amy@findmypast.com>
`,
			wantMsg: `empty mobbed commit

Co-Authored-By: Amy Doe <amy@findmypast.com>
Co-Authored-By: Bob Doe <bob@findmypast.com>

`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AppendCoauthorMarkup(tt.haveAuthors, []byte(tt.haveMsg))
			if (err != nil) != tt.wantErr {
				t.Errorf("AppendCoauthorMarkup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotMsg := string(got)
			if !reflect.DeepEqual(gotMsg, tt.wantMsg) {
				t.Errorf("AppendCoauthorMarkup()\ngot = '%s'\nwant = '%s'", gotMsg, tt.wantMsg)
			}
		})
	}
}
