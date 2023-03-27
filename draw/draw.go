package draw

import (
	"errors"
	"image/color"
	"math"

	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"golang.org/x/image/font"
)

func DrawCoordLine(gc draw2d.GraphicContext, lineCoords [][]float64, lineWidth float64, fillColor color.Color, strokeWidth float64, strokeColor color.Color, scale func(x, y float64) (float64, float64)) error {
	if lineWidth == 0.0 {
		return errors.New("line width cannot be zero")
	}

	// first line is for the stroke (beneath)
	gc.SetStrokeColor(strokeColor)
	gc.SetLineWidth(lineWidth + strokeWidth)

	err := lineCoordSeq(gc, &lineCoords, scale)
	if err != nil {
		return (err)
	}
	gc.Stroke()

	// the actual line
	gc.SetStrokeColor(fillColor)
	gc.SetLineWidth(lineWidth)

	err = lineCoordSeq(gc, &lineCoords, scale)
	if err != nil {
		return (err)
	}
	gc.Stroke()

	return nil
}

func DrawRune(gc *draw2dimg.GraphicContext, pos []float64, f font.Face, rotation float64, char rune) error {
	radians := rotation * (math.Pi / 180)
	rm := draw2d.NewRotationMatrix(radians)

	trlx, trly := rm.InverseTransformPoint(pos[0], pos[1])

	gc.Save()
	gc.Rotate(radians)
	gc.StrokeStringAt(string(char), trlx, trly)
	gc.FillStringAt(string(char), trlx, trly)
	gc.Restore()

	return nil
}

func lineCoordSeq(gc draw2d.GraphicContext, cs *[][]float64, scale func(x, y float64) (float64, float64)) error {
	if cs == nil {
		return errors.New("coord seq cannot be nil")
	}

	csd := *cs

	gc.MoveTo(scale(csd[0][0], csd[0][1]))

	for i := 1; i < len(csd); i++ {
		x, y := scale(csd[i][0], csd[i][1])
		gc.LineTo(x, y)
	}

	return nil
}
