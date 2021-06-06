package main

import (
	"clout/files"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func HandleYoutube() {
	id := argMap["id"]
	fmt.Println(id)
	home := files.UserHomeDir()
	dir := "clout-cli-youtube"
	path := home + "/" + dir
	os.Mkdir(path, 0700)
	cmd := exec.Command("youtube-dl", "--output",
		path+"/%(id)s.%(ext)s",
		"--recode-video", "mp4", id)

	//var out bytes.Buffer
	//cmd.Stdout = &out

	go cmd.Run()

	for {
		b, _ := exec.Command("ls", "-lh", path).CombinedOutput()
		s := string(b)
		fmt.Println(s)
		fmt.Println("")
		time.Sleep(time.Second * 5)
	}

}
