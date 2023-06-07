package main

import (
	"image"
	"log"
	"os"

	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"github.com/Valentin-Foucher/rss-feed-aggregator/pkg/rss"
)

type Button struct {
	pressed  bool
	iterator rss.IItemsIterator
	items    []rss.IItem
}

func drawArrow(ops *op.Ops) layout.Dimensions {
	defer clip.Rect{Max: image.Pt(32, 32)}.Push(ops).Pop()
	arrow, err := os.Open("assets/images/right-arrow.png")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer arrow.Close()
	arrowData, _, err := image.Decode(arrow)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	imageOp := paint.NewImageOp(arrowData)
	imageOp.Add(ops)
	paint.PaintOp{}.Add(ops)
	return layout.Dimensions{Size: imageOp.Size()}
}
func (b *Button) Layout(gtx layout.Context) layout.Dimensions {
	arrow := drawArrow(gtx.Ops)

	for _, e := range gtx.Events(arrow) {
		if e, ok := e.(pointer.Event); ok {
			switch e.Type {
			case pointer.Press:
				if !b.pressed {
					b.items = b.iterator.Next()
					b.pressed = true
				}
			case pointer.Release:
				b.pressed = false
			}
		}
	}

	// Confine the area for pointer events.
	area := clip.Rect(image.Rect(0, 0, arrow.Size.X, arrow.Size.Y)).Push(gtx.Ops)
	pointer.InputOp{
		Tag:   arrow,
		Types: pointer.Press | pointer.Release,
	}.Add(gtx.Ops)
	area.Pop()

	return arrow
}
