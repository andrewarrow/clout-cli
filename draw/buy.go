package draw

import (
	"clout/network"
	"clout/session"
	"fmt"

	"github.com/fogleman/gg"
)

var files = []string{}
var coinX = -220
var buyerX = 600

func DrawBuyStackedChart(top []string, friendMap map[string]float64) {
	dc := gg.NewContext(430, 400)
	dc.SetRGB(0, 0, 0)
	dc.Clear()
	dc.SetRGB(0, 1, 0)

	startX := 10.0
	startY := 50.0
	sizeX := 400.0
	sizeY := 40.0
	dc.LoadFontFace("arial.ttf", 14)
	dc.DrawStringAnchored("0", 20, startY-30, 0.5, 0.5)
	dc.DrawStringAnchored("20", 90, startY-30, 0.5, 0.5)
	dc.DrawStringAnchored("40", 170, startY-30, 0.5, 0.5)
	dc.DrawStringAnchored("60", 250, startY-30, 0.5, 0.5)
	dc.DrawStringAnchored("80", 330, startY-30, 0.5, 0.5)
	dc.DrawStringAnchored("100", 400, startY-30, 0.5, 0.5)
	dc.LoadFontFace("arial.ttf", 24)
	for i, item := range top {
		if i > 4 {
			break
		}
		perX := sizeX * friendMap[item]
		dc.SetRGB(0, 0.5, 0.5)
		dc.MoveTo(startX, startY)
		dc.LineTo(startX+sizeX, startY)
		dc.LineTo(startX+sizeX, startY+sizeY)
		dc.LineTo(startX, startY+sizeY)
		dc.MoveTo(startX, startY)
		dc.ClosePath()
		dc.Fill()
		dc.SetRGB(0, 1, 0)
		dc.MoveTo(startX, startY)
		dc.LineTo(startX+perX, startY)
		dc.LineTo(startX+perX, startY+sizeY)
		dc.LineTo(startX, startY+sizeY)
		dc.MoveTo(startX, startY)
		dc.ClosePath()
		dc.Fill()
		dc.SetRGB(0, 0, 0)
		dc.DrawStringAnchored(item, 200, startY+20, 0.5, 0.5)
		startY += sizeY * 1.5
	}
	dc.SavePNG("chart.png")
}
func DrawBuyFrame(i int, usd float64) {
	dc := gg.NewContext(600, 400)
	dc.SetRGB(0, 0, 0)
	dc.Clear()
	dc.SetRGB(0, 1, 0)
	im, _ := gg.LoadPNG("coin.png")
	dc.DrawImage(im, coinX, 20)
	im, _ = gg.LoadPNG("buyer.png")
	dc.DrawImage(im, buyerX, 20)
	dc.LoadFontFace("arial.ttf", 72)
	dc.DrawStringAnchored(fmt.Sprintf("BUY $%0.2f", usd), 300, 300, 0.5, 0.5)
	file := fmt.Sprintf("frames/%d.png", i)
	dc.SavePNG(file)
	files = append(files, file)
}
func BuyPoster(coin, buyer string, usd float64) {
	GetProfilePicAndEnlarge(coin, "coin")
	GetProfilePicAndEnlarge(buyer, "buyer")

	files = []string{}
	for i := 1; i < 50; i++ {
		DrawBuyFrame(i, usd)
		coinX += 6
		buyerX -= 6
	}
	for i := 50; i < 100; i++ {
		DrawBuyFrame(i, usd)
		coinX -= 6
		buyerX += 6
	}
	MakeVideoFromImages(files)
}

func GetProfilePicAndEnlarge(username, flavor string) {
	pub58 := session.UsernameToPub58(username)
	actorBytes := network.GetSingleProfilePicture(pub58)
	SavePic(flavor, actorBytes)
	ResizeImage(flavor + ".png")
}
