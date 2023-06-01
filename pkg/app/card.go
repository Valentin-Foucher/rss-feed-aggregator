package main

import (
	"image"

	"gioui.org/font"
	"gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/fstanis/screenresolution"
)

var (
	th               = material.NewTheme(gofont.Collection())
	screenResolution = screenresolution.GetPrimary()
)

func cardBackground(gtx layout.Context) layout.Dimensions {
	size := gtx.Constraints.Min
	roundness := 15

	defer clip.RRect{Rect: image.Rect(0, 0, size.X, size.Y), SE: roundness, SW: roundness, NW: roundness, NE: roundness}.Push(gtx.Ops).Pop()
	paint.ColorOp{Color: grey}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	return layout.Dimensions{Size: size}
}

func cardTitle(gtx layout.Context, title string) layout.Dimensions {
	cardTitle := material.H6(th, title)
	cardTitle.Alignment = text.Middle
	cardTitle.Font.Style = font.Style(font.Bold)
	cardTitle.Layout(gtx)

	return layout.Dimensions{Size: image.Pt(screenResolution.Width, screenResolution.Height/30)}
}

func cardDescription(gtx layout.Context, description string) layout.Dimensions {
	cardContent := material.Body1(th, description)
	cardContent.Alignment = text.Middle
	cardContent.Layout(gtx)

	return layout.Dimensions{Size: image.Pt(screenResolution.Width, screenResolution.Height/10)}
}

func cardContent(gtx layout.Context, title, description string) layout.Dimensions {
	return layout.Inset{Top: unit.Dp(5), Bottom: unit.Dp(5), Left: unit.Dp(25), Right: unit.Dp(25)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Flexed(0.3, func(gtx layout.Context) layout.Dimensions {
				return cardTitle(gtx, description)
			}),

			layout.Rigid(layout.Spacer{Width: unit.Dp(gtx.Constraints.Max.X), Height: unit.Dp(20)}.Layout),

			layout.Flexed(0.7, func(gtx layout.Context) layout.Dimensions {
				return cardDescription(gtx, description)
			}))
	})
}

func Card(gtx layout.Context, description, title string) layout.Dimensions {
	return layout.Stack{Alignment: layout.S}.Layout(gtx,
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			return cardBackground(gtx)
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return cardContent(gtx, title, description)
		}),
	)
}
