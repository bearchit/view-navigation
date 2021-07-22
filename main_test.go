package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApp(t *testing.T) {
	vc := Navigation{}

	vc.PushScreen(NewScreen("Home"))
	vc.PushScreen(NewScreen("Detail"))

	purchaseLayer := NewScreen("PurchaseLayer")
	vc.PushScreen(purchaseLayer)

	purchaseLayer.PushView(NewView("Step #1"))

	t.Log(vc.Stack())

	var screen Screen

	screen = vc.PopScreen()
	t.Log(screen)
	assert.Equal(t, "PurchaseLayer", screen.ID())

	screen = vc.PopScreen()
	t.Log(screen)
	assert.Equal(t, "Detail", screen.ID())

	screen = vc.PopScreen()
	t.Log(screen)
	assert.Equal(t, "Home", screen.ID())
}
