package draw

import (
	"fmt"
	"io/ioutil"
)

func UserPoster(pic string) {

	files, _ := ioutil.ReadDir(pic)

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
