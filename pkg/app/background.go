package main

import (
	"image/color"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

func Background(gtx layout.Context, color color.NRGBA) layout.Dimensions {
	size := gtx.Constraints.Max

	defer clip.Rect{Max: size}.Push(gtx.Ops).Pop()
	paint.LinearGradientOp{Stop1: f32.Point{X: 0, Y: 0}, Color1: red, Stop2: f32.Point{X: 1000, Y: 1000}, Color2: blue}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	return layout.Dimensions{Size: size}
}
