package opencv

import "testing"

func TestOpenCV002(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestOpenCV002"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			OpenCV002()
		})
	}
}
