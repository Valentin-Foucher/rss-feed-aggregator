package main

import (
	"fmt"
	"image"
	"image/color"
	"time"

	"gioui.org/font"
	"gioui.org/font/gofont"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/Valentin-Foucher/rss-feed-aggregator/pkg/rss"
	"github.com/fstanis/screenresolution"
	"github.com/pkg/browser"
)

var (
	th               = material.NewTheme(gofont.Collection())
	screenResolution = screenresolution.GetPrimary()
	grey             = color.NRGBA{R: 0xDE, G: 0xDE, B: 0xDE, A: 0xFF}
	isHovering       = false
)

func cardBackground(gtx layout.Context) layout.Dimensions {
	size := gtx.Constraints.Min
	roundness := 15

	defer clip.RRect{Rect: image.Rect(0, 0, size.X, size.Y+20), SE: roundness, SW: roundness, NW: roundness, NE: roundness}.Push(gtx.Ops).Pop()
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

func cardDetails(gtx layout.Context, source string, date *time.Time) layout.Dimensions {
	cardDetails := material.Body1(th, fmt.Sprintf("%s - %s", source, date.Format("2006-01-02")))
	cardDetails.Alignment = text.Start
	cardDetails.Font.Style = font.Style(font.Italic)
	cardDetails.Layout(gtx)

	return layout.Dimensions{Size: image.Pt(screenResolution.Width, screenResolution.Height/100)}
}

func cardDescription(gtx layout.Context, description string) layout.Dimensions {
	cardContent := material.Body1(th, description)
	cardContent.Alignment = text.Start
	cardContent.Layout(gtx)

	return layout.Dimensions{Size: image.Pt(screenResolution.Width, screenResolution.Height/13)}
}

func cardLink(gtx layout.Context, link string) layout.Dimensions {
	// widget content
	cardContent := material.Body1(th, link)
	cardContent.Alignment = text.Start
	cardContent.Font.Style = font.Italic
	cardContent.Color = blue
	cardContent.Layout(gtx)

	for _, e := range gtx.Events(cardContent) {
		if e, ok := e.(pointer.Event); ok {
			switch e.Type {
			case pointer.Press:
				browser.OpenURL(link)
			case pointer.Enter:
				isHovering = true
			case pointer.Leave:
				isHovering = false
				pointer.Cursor.Add(pointer.CursorDefault, gtx.Ops)
			}
		}
	}

	if isHovering {
		pointer.Cursor.Add(pointer.CursorPointer, gtx.Ops)
	} else {
		pointer.Cursor.Add(pointer.CursorDefault, gtx.Ops)
	}

	// event capture area
	area := clip.Rect(image.Rect(0, 0, screenResolution.Width, screenResolution.Height/18)).Push(gtx.Ops)
	pointer.InputOp{
		Types: pointer.Press | pointer.Release | pointer.Leave | pointer.Enter,
		Tag:   cardContent,
		Grab:  false,
	}.Add(gtx.Ops)
	area.Pop()

	return layout.Dimensions{Size: image.Pt(screenResolution.Width, screenResolution.Height/18)}
}

func cardContent(gtx layout.Context, item rss.IItem) layout.Dimensions {
	return layout.Inset{Top: unit.Dp(5), Bottom: unit.Dp(5), Left: unit.Dp(25), Right: unit.Dp(25)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Flexed(0.1, func(gtx layout.Context) layout.Dimensions {
				return cardTitle(gtx, item.Title())
			}),
			layout.Rigid(layout.Spacer{Width: unit.Dp(gtx.Constraints.Max.X), Height: unit.Dp(20)}.Layout),
			layout.Flexed(0.1, func(gtx layout.Context) layout.Dimensions {
				return cardDetails(gtx, item.Source(), item.ParsedPublishedDate())
			}),
			layout.Rigid(layout.Spacer{Width: unit.Dp(gtx.Constraints.Max.X), Height: unit.Dp(30)}.Layout),
			layout.Flexed(0.7, func(gtx layout.Context) layout.Dimensions {
				return cardDescription(gtx, item.Description())
			}),
			layout.Rigid(layout.Spacer{Width: unit.Dp(gtx.Constraints.Max.X), Height: unit.Dp(70)}.Layout),
			layout.Flexed(0.1, func(gtx layout.Context) layout.Dimensions {
				return cardLink(gtx, item.Link())
			}))
	})
}

func Card(gtx layout.Context, item rss.IItem) layout.Dimensions {
	return layout.Stack{Alignment: layout.S}.Layout(gtx,
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			return cardBackground(gtx)
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return cardContent(gtx, item)
		}),
	)
}
