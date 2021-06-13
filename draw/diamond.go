package draw

import "github.com/fogleman/gg"

var sizeX, sizeY float64

func DrawDiamondImage() {
	dc := gg.NewContext(600, 600)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)

	dc.SetLineWidth(2)
	sizeX = float64(100)
	sizeY = float64(50)
	DrawDiamond(dc, 25.0, 100.0+sizeY+sizeY+sizeY)
	DrawDiamond(dc, 25.0, 100.0)
	dc.Stroke()
	dc.SavePNG("001.png")
}

func DrawDiamond(dc *gg.Context, startX, startY float64) {
	dc.DrawLine(startX, startY, startX+sizeX, startY+sizeY)
	dc.DrawLine(startX+sizeX, startY+sizeY, startX+sizeX+sizeX, startY)
	dc.DrawLine(startX+sizeX+sizeX, startY, startX+sizeX, startY-sizeY)
	dc.DrawLine(startX+sizeX, startY-sizeY, startX, startY)
	dc.DrawLine(startX, startY, startX+sizeX, startY+sizeY+sizeY+sizeY)
	dc.DrawLine(startX+sizeX, startY+sizeY+sizeY+sizeY, startX+sizeX+sizeX, startY)
}
