package gitConfig

import "testing"

func TestExitCode_String(t *testing.T) {
	tests := []struct {
		have ExitCode
		want int
	}{
		{
			have: CommandSuccess,
			want: 0,
		},
		{
			have: SectionOrKeyIsInvalid,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.have.String(), func(t *testing.T) {
			if got := int(tt.have); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}

			if got := ExitCode(tt.want); got != tt.have {
				t.Errorf("String() = %v, want %v", got, tt.have)
			}
		})
	}
}
