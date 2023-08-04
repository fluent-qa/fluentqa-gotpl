package utils

import "testing"

func TestGenUUID(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{name: "Get UUID Successfully", want: 36},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenUUID(); len(got) != tt.want {
				t.Errorf("GenUUID() = %v, want %v", len(got), tt.want)
			}
		})
	}
}
