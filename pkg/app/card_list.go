package main

import (
	"image"

	"github.com/Valentin-Foucher/rss-feed-aggregator/pkg/rss"

	"gioui.org/io/key"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/unit"
)

var scrollOffset unit.Dp = 0

func CardListLayout(gtx layout.Context, items []rss.IItem) layout.Dimensions {
	for _, e := range gtx.Events(0) {
		if e, ok := e.(key.Event); ok && e.State == key.Press {
			if e.Name == "U" {
				scrollOffset = scrollOffset - 100
				if scrollOffset < 0 {
					scrollOffset = 0
				}
			} else if e.Name == "D" {
				scrollOffset = scrollOffset + 100
				if scrollOffset < 0 {
					scrollOffset = 0
				}
			}
		}
	}

	eventArea := clip.Rect(
		image.Rectangle{
			Min: image.Point{0, 0},
			Max: image.Point{gtx.Constraints.Max.X, gtx.Constraints.Max.Y},
		},
	).Push(gtx.Ops)

	key.InputOp{
		Keys: key.Set("(Shift)-U|(Shift)-D"),
		Tag:  0,
	}.Add(gtx.Ops)

	list := layout.List{Axis: layout.Vertical, Alignment: layout.Middle, Position: layout.Position{Offset: int(scrollOffset)}}

	articles := layout.Inset{Bottom: unit.Dp(40)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return list.Layout(gtx, len(items), func(gtx layout.Context, i int) layout.Dimensions {
			item := items[i]
			return layout.Inset{Top: unit.Dp(15), Right: unit.Dp(15), Bottom: unit.Dp(35), Left: unit.Dp(15)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return Card(gtx, item)
			})
		})
	})

	eventArea.Pop()

	return articles
}
