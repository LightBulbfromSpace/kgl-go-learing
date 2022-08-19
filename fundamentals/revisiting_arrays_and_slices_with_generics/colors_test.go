package revisiting_arrays_and_slices

import (
	"testing"
)

//все нуждается в доработке

func TestMixColor(t *testing.T) {
	black, err := NewColor("#000000", 0.5)
	if err != nil {
		t.Error(err)
	}
	white, err := NewColor("#FFFFFF", 0.5)
	if err != nil {
		t.Error(err)
	}
	gray, err := NewColor("#7F7F7F", 0.5)
	got, err := ColorMixer(black, white)
	if err != nil {
		t.Error(err)
	}
	if got != gray {
		t.Errorf("Expected to get %X%X%X, alpha: %f, but got %X%X%X, alpha: %f",
			gray.R, gray.G, gray.B, gray.A, got.R, got.G, got.B, got.A)
	}
}
