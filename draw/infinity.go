package draw

import (
	"fmt"
	"os/exec"

	"github.com/fogleman/gg"
)

func MakeVideoFromImages() {
	b, err := exec.Command("ffmpeg", "-y", "-framerate", "1/2", "-i", "frames/%03d.png", "-loop", "0", "output.gif").CombinedOutput()
	fmt.Println(string(b), err)
}
func DrawInfinities() {
	one := 1
	three := 3
	nine := 9
	for i := 1; i < 10; i++ {
		dc := gg.NewContext(600, 400)
		dc.SetRGB(1, 1, 1)
		dc.Clear()
		dc.SetRGB(0, 0, 0)
		DrawInfinityFrame(dc, fmt.Sprintf("%03d.png", i), one, three, nine)
		one = one * 2
		three = three * 2
		nine = nine * 2
	}
	MakeVideoFromImages()
}
func DrawInfinityFrame(dc *gg.Context, filename string, one, three, nine int) {
	font := "arial.ttf"
	dc.LoadFontFace(font, 18)
	dc.DrawStringAnchored(fmt.Sprintf("%d", one), 50, 100, 0.5, 0.5)
	dc.DrawStringAnchored(fmt.Sprintf("%d", three), 50+200, 100, 0.5, 0.5)
	dc.DrawStringAnchored(fmt.Sprintf("%d", nine), 50+400, 100, 0.5, 0.5)
	dc.SavePNG("frames/" + filename)
}
