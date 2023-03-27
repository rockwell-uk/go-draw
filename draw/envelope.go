package draw

type Envelope struct {
	Min, Max []float64
}

func (e Envelope) Dx() float64 {
	return e.Max[0] - e.Min[0]
}

func (e Envelope) Dy() float64 {
	return e.Max[1] - e.Min[1]
}

func (e Envelope) Px(x float64) float64 {
	return (x - e.Min[0]) / e.Dx()
}

func (e Envelope) Py(y float64) float64 {
	return (y - e.Min[1]) / e.Dy()
}

func ToEnvelope(lineCoords [][]float64) (Envelope, error) {
	minX := lineCoords[0][0]
	maxX := lineCoords[2][0]
	minY := lineCoords[0][1]
	maxY := lineCoords[2][1]

	return Envelope{
		[]float64{
			minX,
			minY,
		},
		[]float64{
			maxX,
			maxY,
		},
	}, nil
}
