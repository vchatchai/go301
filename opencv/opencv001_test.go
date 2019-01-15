package opencv

import "testing"

func TestOpenCV001(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestOpenCV001"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			OpenCV001()
		})
	}
}
