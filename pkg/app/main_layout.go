package main

import (
	"gioui.org/layout"
	"gioui.org/unit"
)

func MainLayout(gtx layout.Context, paginator *Paginator) layout.Dimensions {
	return layout.Stack{Alignment: layout.S}.Layout(gtx,
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			return Background(gtx)
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return CardListLayout(gtx, paginator.items)
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
				layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
					return drawBackButton(gtx, paginator)
				}),
				layout.Rigid(layout.Spacer{Width: unit.Dp(20)}.Layout),
				layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
					return drawForwardButton(gtx, paginator)
				}),
			)
		}),
	)
}
