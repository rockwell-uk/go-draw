package draw

import (
	"reflect"
	"testing"
)

func TestEnvelope(t *testing.T) {
	tests := map[string]struct {
		dim      float64
		expected Envelope
	}{
		"4000": {
			4000,
			Envelope{
				[]float64{
					0,
					0,
				},
				[]float64{
					4000,
					4000,
				},
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			tileWidth := tt.dim
			tileHeight := tt.dim

			bounds := [][]float64{
				{0, 0},
				{0, tileWidth},
				{tileWidth, tileHeight},
				{tileHeight, 0},
				{0, 0},
			}

			envelope, err := ToEnvelope(bounds)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(tt.expected, envelope) {
				t.Fatalf("expected %#v, actual %#v", tt.expected, envelope)
			}
		})
	}
}
