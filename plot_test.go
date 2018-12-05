package go301

import "testing"

func TestPlot(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestPlot"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Plot()
		})
	}
}
