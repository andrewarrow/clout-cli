package draw

import (
	"clout/network"
	"clout/session"

	"github.com/fogleman/gg"
)

func BuyPoster() {
	dc := gg.NewContext(600, 400)
	dc.SetRGB(0, 0, 0)
	dc.Clear()
	dc.SetRGB(0, 1, 0)
	coin := "nobodygetsit"
	buyer := "buildguild"
	GetProfilePicAndEnlarge(coin)
	im, _ := gg.LoadPNG("actor.png")
	dc.DrawImage(im, 10, 20)
	GetProfilePicAndEnlarge(buyer)
	im, _ = gg.LoadPNG("actor.png")
	dc.DrawImage(im, 250, 20)
	dc.SavePNG("frames/001.png")
}

func GetProfilePicAndEnlarge(username string) {
	pub58 := session.UsernameToPub58(username)
	actorBytes := network.GetSingleProfilePicture(pub58)
	SavePic("actor", actorBytes)
	ResizeImage("actor.png")
}
