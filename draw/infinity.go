package draw

import (
	"fmt"
	"os/exec"

	"github.com/fogleman/gg"
)

func MakeVideoFromImages(files []string) {
	/*

		ffmpeg -f image2 -framerate 1 -i img_%02d.png -filter_complex "drawtext=enable='between(t,0,1)':text='word1':fontsize=24:fontcolor=white:x=w-tw:y=0,drawtext=enable='between(t,0,0.9)':text='word2':fontsize=24:fontcolor=white:x=w-tw:y=0" out.gif
	*/
	//b, err := exec.Command("ffmpeg", "-y", "-framerate", "1/2", "-i", "frames/%03d.png", "-loop", "0", "output.gif").CombinedOutput()
	params := []string{"-o", "output.gif", "-r", "1", "-Q", "100"}
	b, err := exec.Command("gifski", append(params, files...)...).CombinedOutput()
	fmt.Println(string(b), err)
}
func DrawInfinities() {
	one := 1
	three := 3
	nine := 9
	files := []string{}
	for i := 1; i < 10; i++ {
		dc := gg.NewContext(600, 400)
		dc.SetRGB(1, 1, 1)
		dc.Clear()
		dc.SetRGB(0, 0, 0)
		file := fmt.Sprintf("%03d.png", i)
		DrawInfinityFrame(dc, file, one, three, nine)
		files = append(files, "frames/"+file)
		one = one * 2
		three = three * 2
		nine = nine * 2
	}
	MakeVideoFromImages(files)
}
func DrawInfinityFrame(dc *gg.Context, filename string, one, three, nine int) {
	font := "arial.ttf"
	dc.LoadFontFace(font, 18)
	dc.DrawStringAnchored(fmt.Sprintf("%d", one), 50, 100, 0.5, 0.5)
	dc.DrawStringAnchored(fmt.Sprintf("%d", three), 50+200, 100, 0.5, 0.5)
	dc.DrawStringAnchored(fmt.Sprintf("%d", nine), 50+400, 100, 0.5, 0.5)
	dc.SavePNG("frames/" + filename)
}
