package dao

import "testing"

func TestRUNDB(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"1"},
		{"as"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RUNDB()
		})
	}
}
