package go301

import "testing"

func TestHello(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestHello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Hello()
		})
	}
}
