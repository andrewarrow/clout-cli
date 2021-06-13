package draw

import "github.com/fogleman/gg"

func DrawDiamond() {
	dc := gg.NewContext(600, 600)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)

	dc.SetLineWidth(2)
	startX := float64(25)
	startY := float64(100)
	//sizeX := float64(100)
	//sizeY := float64(50)
	size := float64(50)
	dc.DrawLine(startX, startY, startX+size, startY+size)
	dc.DrawLine(startX+size, startY+size, startX+size+size, startY)
	dc.DrawLine(startX+size+size, startY, startX+size, startY-size)
	dc.DrawLine(startX+size, startY-size, startX, startY)

	dc.DrawLine(startX, startY, startX+size, startY+size+size+size)
	dc.DrawLine(startX+size, startY+size+size+size, startX+size+size, startY)
	dc.Stroke()
	dc.SavePNG("001.png")
}
