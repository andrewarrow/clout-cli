package draw

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func ResizeImage(filename string) {
	exec.Command("convert", filename, "-resize", "200%", filename).Output()
}
func SavePic(flavor string, data []byte) {
	os.Remove(flavor + ".webp")
	ioutil.WriteFile(flavor+".webp", data, 0755)
	exec.Command("convert", flavor+".webp", flavor+".png").Output()
}
func SavePicWithPath(imageKind, path, flavor string, data []byte) {
	os.Remove(path + "/" + flavor + "." + imageKind)
	fmt.Println(path + "/" + flavor + "." + imageKind)
	ioutil.WriteFile(path+"/"+flavor+"."+imageKind, data, 0755)
	b, _ := exec.Command("convert", path+"/"+flavor+"."+imageKind, path+"/"+flavor+".png").CombinedOutput()
	fmt.Println(string(b))

}
