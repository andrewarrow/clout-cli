package draw

import "github.com/fogleman/gg"

func DrawDiamond() {
	dc := gg.NewContext(600, 600)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)

	dc.SetLineWidth(2)
	dc.DrawRectangle(50, 50, 100, 100)
	dc.Stroke()
	dc.SavePNG("001.png")
}
