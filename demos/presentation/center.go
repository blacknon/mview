package main

import "github.com/blacknon/mview"

// Center returns a new primitive which shows the provided primitive in its
// center, given the provided primitive's size.
func Center(width, height int, p mview.Primitive) mview.Primitive {
	subFlex := mview.NewFlex()
	subFlex.SetDirection(mview.FlexRow)
	subFlex.AddItem(mview.NewBox(), 0, 1, false)
	subFlex.AddItem(p, height, 1, true)
	subFlex.AddItem(mview.NewBox(), 0, 1, false)

	flex := mview.NewFlex()
	flex.AddItem(mview.NewBox(), 0, 1, false)
	flex.AddItem(subFlex, width, 1, true)
	flex.AddItem(mview.NewBox(), 0, 1, false)

	return flex
}
