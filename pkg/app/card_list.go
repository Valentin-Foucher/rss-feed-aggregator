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

func CardListLayout(gtx layout.Context, iterator rss.IItemsIterator) layout.Dimensions {
	items := iterator.Next()

	for _, e := range gtx.Events(0) {
		if e, ok := e.(key.Event); ok && e.State == key.Press {
			if e.Name == "U" {
				scrollOffset = scrollOffset - 50
				if scrollOffset < 0 {
					scrollOffset = 0
				}
			} else if e.Name == "D" {
				scrollOffset = scrollOffset + 50
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

	articles := list.Layout(gtx, len(items), func(gtx layout.Context, i int) layout.Dimensions {
		item := items[i]
		return layout.UniformInset(unit.Dp(15)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return Card(gtx, item.Title(), item.Description(), item.Link())
		})
	})

	eventArea.Pop()

	return articles
}
