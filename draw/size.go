package draw

import (
	"io/ioutil"
	"os"
	"os/exec"
)

func ResizeImage(filename string) {
	exec.Command("magick", "convert", filename, "-resize", "50%", filename).Output()
}
func SavePic(flavor string, data []byte) {
	os.Remove(flavor + ".webp")
	ioutil.WriteFile(flavor+".webp", data, 0755)
	exec.Command("convert", flavor+".webp", flavor+".png").Output()
}
