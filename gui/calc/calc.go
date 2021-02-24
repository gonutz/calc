package main

import (
	"github.com/gonutz/calc"
	"github.com/gonutz/wui/v2"
)

func main() {
	tahoma, _ := wui.NewFont(wui.FontDesc{Name: "Tahoma", Height: -16})

	window := wui.NewWindow()
	window.SetTitle("Calc")
	window.SetInnerSize(300, 500)
	window.SetFont(tahoma)

	c := calc.NewCalculator()

	long := wui.NewLabel()
	long.SetAlignment(wui.AlignRight)
	long.SetBounds(10, 30, 280, 20)
	window.Add(long)

	short := wui.NewLabel()
	short.SetAlignment(wui.AlignRight)
	short.SetBounds(10, 50, 280, 20)
	window.Add(short)

	show := func() {
		short.SetText(c.ShortOutput())
		long.SetText(c.LongOutput())
	}

	show()

	button := func(text string, x, y int, keys ...wui.Key) *wui.Button {
		tileSize := window.InnerWidth() / 4
		b := wui.NewButton()
		b.SetText(text)
		b.SetBounds(x*tileSize, window.InnerHeight()-(y+1)*tileSize, tileSize, tileSize)
		b.SetOnClick(func() {
			for _, r := range text {
				c.Input(r)
			}
			show()
		})
		window.Add(b)
		for _, key := range keys {
			window.SetShortcut(b.OnClick(), key)
		}
		return b
	}

	button("0", 1, 0, '0', wui.KeyNum0)
	button("1", 0, 1, '1', wui.KeyNum1)
	button("2", 1, 1, '2', wui.KeyNum2)
	button("3", 2, 1, '3', wui.KeyNum3)
	button("4", 0, 2, '4', wui.KeyNum4)
	button("5", 1, 2, '5', wui.KeyNum5)
	button("6", 2, 2, '6', wui.KeyNum6)
	button("7", 0, 3, '7', wui.KeyNum7)
	button("8", 1, 3, '8', wui.KeyNum8)
	button("9", 2, 3, '9', wui.KeyNum9)
	button(".", 0, 0, ',', '.', wui.KeyOEMPeriod, wui.KeyDecimal)

	button("+", 3, 3, '+', wui.KeyAdd, wui.KeyOEMPlus)
	button("-", 3, 2, '-', wui.KeySubtract, wui.KeyOEMMinus)
	button("*", 3, 1, '*', wui.KeyMultiply)
	button("/", 3, 0, '/', wui.KeyDivide)

	button("=", 2, 0, wui.KeyReturn)
	button("C", 0, 4, wui.KeyEscape, 'C')
	button("N", 1, 4, 'N').SetText("+/-")

	window.Show()
}
