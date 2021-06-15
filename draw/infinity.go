package draw

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/fogleman/gg"
)

func MakeVideoFromImages(files []string) {
	/*

		ffmpeg -f image2 -framerate 1 -i img_%02d.png -filter_complex "drawtext=enable='between(t,0,1)':text='word1':fontsize=24:fontcolor=white:x=w-tw:y=0,drawtext=enable='between(t,0,0.9)':text='word2':fontsize=24:fontcolor=white:x=w-tw:y=0" out.gif
	*/
	//b, err := exec.Command("ffmpeg", "-y", "-framerate", "1/2", "-i", "frames/%03d.png", "-loop", "0", "output.gif").CombinedOutput()
	params := []string{"-o", "output.gif", "-r", "0.25", "-Q", "100"}
	b, err := exec.Command("gifski", append(params, files...)...).CombinedOutput()
	fmt.Println(string(b), err)
}
func DrawInfinities() {
	one := 1
	three := 3
	nine := 9
	files := []string{}
	for i := 1; i < 20; i++ {
		dc := gg.NewContext(600, 400)
		dc.SetRGB(0, 0, 0)
		dc.Clear()
		dc.SetRGB(0, 1, 0)
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
	dc.DrawStringAnchored(fmt.Sprintf("%d", one), 100, 100, 0.5, 0.5)
	dc.DrawStringAnchored(fmt.Sprintf("%d", three), 100+200, 100, 0.5, 0.5)
	dc.DrawStringAnchored(fmt.Sprintf("%d", nine), 100+400, 100, 0.5, 0.5)

	lines := []string{}
	lines = AsciiByteAddition(lines, fmt.Sprintf("%d", one))
	dc.LoadFontFace(font, 14)
	for i, line := range lines {
		dc.DrawStringAnchored(line, 100, float64(150+(i*25)), 0.5, 0.5)
	}
	dc.DrawStringAnchored("1,2,4,8,7,5", 100, float64(150+(len(lines)*25)), 0.5, 0.5)

	lines = []string{}
	lines = AsciiByteAddition(lines, fmt.Sprintf("%d", three))
	dc.LoadFontFace(font, 14)
	for i, line := range lines {
		dc.DrawStringAnchored(line, 100+200, float64(150+(i*25)), 0.5, 0.5)
	}
	dc.DrawStringAnchored("3,6", 100+200, float64(150+(len(lines)*25)), 0.5, 0.5)

	lines = []string{}
	lines = AsciiByteAddition(lines, fmt.Sprintf("%d", nine))
	dc.LoadFontFace(font, 14)
	for i, line := range lines {
		dc.DrawStringAnchored(line, 100+400, float64(150+(i*25)), 0.5, 0.5)
	}
	dc.DrawStringAnchored("9", 100+400, float64(150+(len(lines)*25)), 0.5, 0.5)

	dc.SavePNG("frames/" + filename)
}
func AsciiByteAddition(lines []string, a string) []string {

	sum := byte(0)
	buff := []string{}
	for i := range a {

		word := a[i : i+1]
		t, _ := strconv.Atoi(word)
		buff = append(buff, fmt.Sprintf("%d", t))

		sum += byte(t)
	}
	strSum := fmt.Sprintf("%d", sum)
	lines = append(lines, strings.Join(buff, "+")+"="+strSum)
	if len(strSum) > 1 {
		return AsciiByteAddition(lines, strSum)
	}
	lines = append(lines, strSum)
	return lines
}
