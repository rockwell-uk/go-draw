package draw

import (
	"errors"
	"math"
)

func Circle(origin []float64, radius float64, numPoints int) ([][]float64, error) {
	d := float64(360)

	var ai float64
	points := [][]float64{}
	ai = d / float64(numPoints) * math.Pi / 180

	c := 0.0
	for i := 0; i <= numPoints; i++ {
		c += ai
		if c > d {
			break
		}
		x := radius*math.Cos(c) + origin[0]
		y := radius*math.Sin(c) + origin[1]

		points = append(points, []float64{x, y})
	}

	n := len(points)
	if n == 0 {
		return [][]float64{}, errors.New("no points in circle")
	}

	return points, nil
}
