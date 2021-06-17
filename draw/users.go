package draw

import (
	"fmt"
	"io/ioutil"
	"math/rand"

	"github.com/fogleman/gg"
)

func DrawUserFrame(i int, list []string) {
	dc := gg.NewContext(600, 400)
	dc.SetRGB(0, 0, 0)
	dc.Clear()
	dc.SetRGB(0, 1, 0)
	x := rand.Intn(600)
	y := rand.Intn(400)
	for _, image := range list {
		im, _ := gg.LoadPNG(image)
		dc.DrawImage(im, x, y)
		x = rand.Intn(650) - 50
		y = rand.Intn(350) - 50
	}
	file := fmt.Sprintf("frames/%d.png", i)
	dc.SavePNG(file)
	files = append(files, file)
}
func UserPoster(pic string) {
	if true {
		for i := 0; i < 93; i++ {
			file := fmt.Sprintf("frames/%d.png", i)
			files = append(files, file)
		}
		MakeVideoFromImages(files)
		return
	}

	files, _ := ioutil.ReadDir(pic)
	all := []string{}
	for _, file := range files {
		all = append(all, pic+"/"+file.Name())
	}

	j := 0
	for {
		if j*1000 > 92100 {
			break
		}
		list := []string{}
		for i, file := range all {
			if i < j*1000 {
				continue
			}
			list = append(list, file)
			if i > (j*1000)+1000 {
				break
			}
		}

		DrawUserFrame(j, list)
		j++
	}
}
