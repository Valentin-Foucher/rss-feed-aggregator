package main

import (
	"fmt"
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

type Paginator struct {
	left  Button
	right Button

	iterator rss.IItemsIterator
	items    []rss.IItem
}

type Button struct {
	pressed bool
}

func drawForwardArrow(ops *op.Ops) layout.Dimensions {
	return drawArrow(ops, "right")
}

func drawBackArrow(ops *op.Ops) layout.Dimensions {
	return drawArrow(ops, "left")
}

func drawArrow(ops *op.Ops, direction string) layout.Dimensions {
	defer clip.Rect{Max: image.Pt(32, 32)}.Push(ops).Pop()
	arrow, err := os.Open(fmt.Sprintf("assets/images/%s-arrow.png", direction))
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

func drawForwardButton(gtx layout.Context, paginator *Paginator) layout.Dimensions {
	arrow := drawForwardArrow(gtx.Ops)
	b := paginator.right

	for _, e := range gtx.Events(1) {
		if e, ok := e.(pointer.Event); ok {
			switch e.Type {
			case pointer.Press:
				if !b.pressed {
					if paginator.iterator.HasNext() {
						paginator.items = paginator.iterator.Next()
					}
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
		Tag:   1,
		Types: pointer.Press | pointer.Release,
	}.Add(gtx.Ops)
	area.Pop()

	return arrow
}

func drawBackButton(gtx layout.Context, paginator *Paginator) layout.Dimensions {
	arrow := drawBackArrow(gtx.Ops)
	b := paginator.left

	for _, e := range gtx.Events(2) {
		if e, ok := e.(pointer.Event); ok {
			switch e.Type {
			case pointer.Press:
				if !b.pressed {
					if paginator.iterator.HasPrevious() {
						paginator.items = paginator.iterator.Previous()
					}
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
		Tag:   2,
		Types: pointer.Press | pointer.Release,
	}.Add(gtx.Ops)
	area.Pop()

	return arrow
}

func newPaginator(iter rss.IItemsIterator) *Paginator {
	paginator := new(Paginator)
	paginator.iterator = iter
	paginator.left = Button{pressed: false}
	paginator.right = Button{pressed: false}
	paginator.items = iter.Next()
	return paginator
}
