package draw

import (
	"fmt"
	"strings"

	"github.com/fogleman/gg"
)

func DrawPeaceVideo() {
	lines := strings.Split(text, "\n")
	j := 0
	total := 0
	for i := 1; i < 2309; i += 10 {
		total++
		file := fmt.Sprintf("/Users/andrewarrow/youtube/foo/%04d.png", i)
		ResizeImageBy(file, "50")
		files = append(files, file)
		dc := gg.NewContext(960, 540)
		im, _ := gg.LoadPNG(file)
		dc.DrawImage(im, 0, 0)
		dc.LoadFontFace("arial.ttf", 48)
		dc.SetRGB(1, 1, 1)
		dc.DrawStringAnchored(lines[j], 482, 402, 0.5, 0.5)
		dc.SetRGB(0, 1, 0)
		dc.DrawStringAnchored(lines[j], 480, 400, 0.5, 0.5)
		dc.SavePNG(file)
		fmt.Println(i, total, total%10, j, lines[j])
		if total%10 == 0 && total != 0 && j != len(lines)-1 {
			j++
		}
	}
	for i := 1; i < 2309; i += 10 {
		file := fmt.Sprintf("/Users/andrewarrow/youtube/foo/%04d.png", i)
		files = append(files, file)
	}
	MakeVideoFromImages(files)
}

var text = `But peace is what you want most of the time.
That's interesting. 
You can convert peace to happiness
anytime you want.
If you're a peaceful person, 
anything you do will be a happy activity. 
And by the way, being on social media, 
engaging in politics, will not bring you peace.
There is nothing less peaceful.
Right.
In today's day and age?
The way we think you get peace is by 
resolving all your external problems. 
But there is unlimited external problems. 
So the only way to actually get peace on 
the inside by giving up this idea of problems.
Who thinks you can get peace by resolving 
external problems other than politicians?
Everybody.
Yeah?
That's what everybody struggling to do, right? 
Why are you trying to make money? 
To solve all your money problems.`
