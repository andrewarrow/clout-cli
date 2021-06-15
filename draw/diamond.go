package draw

import (
	"clout/network"
	"clout/session"

	"github.com/fogleman/gg"
)

var sizeX, sizeY float64

func DrawDiamondImage(argMap map[string]string) {
	if argMap["pie"] != "" {
		DrawPieImage()
		return
	}
	dc := gg.NewContext(600, 600)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	//im, _ := gg.LoadPNG("actor.png")
	//dc.DrawImage(im, 0, 0)

	username := argMap["username"]
	if username != "" {
		pub58 := session.UsernameToPub58(username)
		actorBytes := network.GetSingleProfilePicture(pub58)
		SavePic("actor", actorBytes)
		ResizeImage("actor.png")
	}

	dc.SetLineWidth(4)
	sizeX = float64(100)
	sizeY = float64(50)
	DrawDiamond(dc, 150.0, 100.0+sizeY+sizeY+(sizeY/2.0))
	DrawDiamond(dc, 150.0, 100.0)
	dc.SavePNG("001.png")
}

func DrawDiamond(dc *gg.Context, startX, startY float64) {
	im, _ := gg.LoadPNG("water.png")
	pattern := gg.NewSurfacePattern(im, gg.RepeatBoth)
	dc.MoveTo(startX, startY)
	dc.LineTo(startX+sizeX, startY+sizeY)
	dc.LineTo(startX+sizeX+sizeX, startY)
	dc.LineTo(startX+sizeX, startY-sizeY)
	dc.LineTo(startX, startY)
	dc.ClosePath()
	dc.SetFillStyle(pattern)
	dc.Fill()
	im, _ = gg.LoadPNG("actor.png")
	pattern = gg.NewSurfacePattern(im, gg.RepeatBoth)
	dc.MoveTo(startX, startY)
	dc.LineTo(startX+sizeX, startY+sizeY+sizeY+(sizeY/2.0))
	dc.LineTo(startX+sizeX+sizeX, startY)
	dc.LineTo(startX+sizeX, startY+sizeY)
	dc.LineTo(startX, startY)
	dc.ClosePath()
	dc.SetFillStyle(pattern)
	dc.Fill()

	dc.DrawLine(startX, startY, startX+sizeX, startY+sizeY)
	dc.DrawLine(startX+sizeX, startY+sizeY, startX+sizeX+sizeX, startY)
	dc.DrawLine(startX+sizeX+sizeX, startY, startX+sizeX, startY-sizeY)
	dc.DrawLine(startX+sizeX, startY-sizeY, startX, startY)
	dc.DrawLine(startX, startY, startX+sizeX, startY+sizeY+sizeY+(sizeY/2.0))
	dc.DrawLine(startX+sizeX, startY+sizeY+sizeY+(sizeY/2.0), startX+sizeX+sizeX, startY)
	dc.Stroke()
}
