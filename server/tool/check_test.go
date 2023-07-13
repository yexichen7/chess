package tool

import "testing"

func TestLengthCheck(t *testing.T) {
	type args struct {
		ss string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		// TODO: Add test cases.
		{
			"1",
			args{"8"},
			"8",
			true,
		}, {
			"1",
			args{"88"},
			"88",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := LengthCheck(tt.args.ss)
			if got != tt.want {
				t.Errorf("LengthCheck() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("LengthCheck() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPostLengthCheck(t *testing.T) {
	type args struct {
		ss string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		// TODO: Add test cases.
		{
			"1",
			args{"8"},
			"8",
			true,
		},
		{
			"1",
			args{"88"},
			"88",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := PostLengthCheck(tt.args.ss)
			if got != tt.want {
				t.Errorf("PostLengthCheck() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("PostLengthCheck() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestScoreCheck(t *testing.T) {
	type args struct {
		score float32
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		// TODO: Add test cases.
		{
			"1",
			args{8},
			"8",
			true,
		},
		{
			"1",
			args{88},
			"8",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ScoreCheck(tt.args.score)
			if got != tt.want {
				t.Errorf("ScoreCheck() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ScoreCheck() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
