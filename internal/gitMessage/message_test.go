package gitMessage

import (
	"fmt"
	"testing"
)

func Test_replaceCoauthors(t *testing.T) {
	tests := []struct {
		name        string
		haveOld     string
		haveContent string
		want        string
	}{
		{
			haveOld:     "",
			haveContent: "\n\nCo-authored-by: Bob Doe <bob@findmypast.com>",
			want:        "\n\nCo-authored-by: Bob Doe <bob@findmypast.com>",
		},
		{
			haveOld: `some common message

Co-authored-by: Amy Doe <amy@findmypast.com>
`,
			haveContent: "\n\nCo-authored-by: Bob Doe <bob@findmypast.com>",
			want:        "some common message\n\nCo-authored-by: Bob Doe <bob@findmypast.com>",
		},
		{
			haveOld: `

Co-authored-by: Amy Doe <amy@findmypast.com>
`,
			haveContent: "\n\nCo-authored-by: Bob Doe <bob@findmypast.com>",
			want:        "\n\nCo-authored-by: Bob Doe <bob@findmypast.com>",
		},
		{
			haveOld: `

Co-authored-by: Amy Doe <amy@findmypast.com>
Co-authored-by: Bob Doe <amy@findmypast.com>
`,
			haveContent: "\n\nCo-authored-by: Bob Doe <bob@findmypast.com>",
			want:        "\n\nCo-authored-by: Bob Doe <bob@findmypast.com>",
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("scenario #%d", i+1), func(t *testing.T) {
			if got := replaceCoauthors([]byte(tt.haveOld), tt.haveContent); got != tt.want {
				t.Errorf("replaceCoauthors() = %v, want %v", got, tt.want)
			}
		})
	}
}
