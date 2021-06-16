package draw

import (
	"clout/network"
	"clout/session"
	"fmt"

	"github.com/fogleman/gg"
)

func BuyPoster() {
	coin := "nobodygetsit"
	buyer := "buildguild"
	GetProfilePicAndEnlarge(coin, "coin")
	GetProfilePicAndEnlarge(buyer, "buyer")

	coinX := -200
	buyerX := 600
	files := []string{}
	for i := 1; i < 100; i++ {
		dc := gg.NewContext(600, 400)
		dc.SetRGB(0, 0, 0)
		dc.Clear()
		dc.SetRGB(0, 1, 0)
		im, _ := gg.LoadPNG("coin.png")
		dc.DrawImage(im, coinX, 20)
		im, _ = gg.LoadPNG("buyer.png")
		dc.DrawImage(im, buyerX, 20)
		file := fmt.Sprintf("frames/%d.png", i)
		dc.SavePNG(file)
		files = append(files, file)
		coinX += 20
		buyerX -= 20
	}
	MakeVideoFromImages(files)
}

func GetProfilePicAndEnlarge(username, flavor string) {
	pub58 := session.UsernameToPub58(username)
	actorBytes := network.GetSingleProfilePicture(pub58)
	SavePic(flavor, actorBytes)
	ResizeImage(flavor + ".png")
}
