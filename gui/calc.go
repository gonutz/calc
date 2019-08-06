package main

import (
	"github.com/gonutz/calc"
	"github.com/gonutz/w32"
	"github.com/gonutz/wui"
)

func main() {
	tahoma, _ := wui.NewFont(wui.FontDesc{Name: "Tahoma", Height: -16})

	window := wui.NewWindow()
	window.SetTitle("Calc")
	window.SetClientSize(300, 500)
	window.SetFont(tahoma)

	c := calc.NewCalculator()

	long := wui.NewLabel()
	long.SetRightAlign()
	long.SetBounds(10, 30, 280, 20)
	window.Add(long)

	short := wui.NewLabel()
	short.SetRightAlign()
	short.SetBounds(10, 50, 280, 20)
	window.Add(short)

	show := func() {
		short.SetText(c.ShortOutput())
		long.SetText(c.LongOutput())
	}

	show()

	button := func(text string, x, y int, keys ...interface{}) *wui.Button {
		tileSize := window.ClientWidth() / 4
		b := wui.NewButton()
		b.SetText(text)
		b.SetBounds(x*tileSize, window.ClientHeight()-(y+1)*tileSize, tileSize, tileSize)
		b.SetOnClick(func() {
			for _, r := range text {
				c.Input(r)
			}
			show()
		})
		window.Add(b)
		for _, key := range keys {
			if vk, ok := key.(int); ok {
				window.SetShortcut(wui.ShortcutKeys{Key: uint16(vk)}, b.OnClick())
			} else if r, ok := key.(int32); ok {
				window.SetShortcut(wui.ShortcutKeys{Rune: r}, b.OnClick())
			}
		}
		return b
	}

	button("0", 1, 0, '0', w32.VK_NUMPAD0)
	button("1", 0, 1, '1', w32.VK_NUMPAD1)
	button("2", 1, 1, '2', w32.VK_NUMPAD2)
	button("3", 2, 1, '3', w32.VK_NUMPAD3)
	button("4", 0, 2, '4', w32.VK_NUMPAD4)
	button("5", 1, 2, '5', w32.VK_NUMPAD5)
	button("6", 2, 2, '6', w32.VK_NUMPAD6)
	button("7", 0, 3, '7', w32.VK_NUMPAD7)
	button("8", 1, 3, '8', w32.VK_NUMPAD8)
	button("9", 2, 3, '9', w32.VK_NUMPAD9)
	button(".", 0, 0, ',', '.', w32.VK_OEM_PERIOD, w32.VK_DECIMAL)

	button("+", 3, 3, '+', w32.VK_ADD, w32.VK_OEM_PLUS)
	button("-", 3, 2, '-', w32.VK_SUBTRACT, w32.VK_OEM_MINUS)
	button("*", 3, 1, '*', w32.VK_MULTIPLY)
	button("/", 3, 0, '/', w32.VK_DIVIDE)

	button("=", 2, 0, w32.VK_RETURN)
	button("C", 0, 4, w32.VK_ESCAPE, 'C')

	window.Show()
}
