package main

import (
	"image"

	"gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/widget/material"
	"github.com/fstanis/screenresolution"
)

var (
	th               = material.NewTheme(gofont.Collection())
	screenResolution = screenresolution.GetPrimary()
)

func Card(gtx layout.Context, content, title string) layout.Dimensions {

	return layout.Stack{Alignment: layout.S}.Layout(gtx,
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			size := image.Pt(screenResolution.Width, screenResolution.Height/10+screenResolution.Height/30)
			roundness := 15

			defer clip.RRect{Rect: image.Rect(0, 0, size.X, size.Y), SE: roundness, SW: roundness, NW: roundness, NE: roundness}.Push(gtx.Ops).Pop()
			paint.ColorOp{Color: white}.Add(gtx.Ops)
			paint.PaintOp{}.Add(gtx.Ops)
			return layout.Dimensions{Size: size}
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Flexed(0.2, func(gtx layout.Context) layout.Dimensions {
					cardTitle := material.Body1(th, title)
					cardTitle.Alignment = text.Middle
					cardTitle.Layout(gtx)

					return layout.Dimensions{Size: image.Pt(screenResolution.Width, screenResolution.Height/30)}
				}),
				layout.Flexed(0.8, func(gtx layout.Context) layout.Dimensions {
					cardContent := material.Body2(th, content)
					cardContent.Alignment = text.Middle
					cardContent.Layout(gtx)

					return layout.Dimensions{Size: image.Pt(screenResolution.Width, screenResolution.Height/10)}
				}))
		}),
	)
}
