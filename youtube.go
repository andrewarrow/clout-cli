package main

import (
	"clout/files"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func SetupYoutubeDirectory() string {
	home := files.UserHomeDir()
	dir := "clout-cli-youtube"
	path := home + "/" + dir
	os.Mkdir(path, 0700)
	return path
}
func HandleYoutube() {
	id := argMap["id"]
	fmt.Println(id)

	action := argMap["action"]

	if action == "" {
		action = "download"
	}

	if action == "download" {
		DownloadYoutube(id)
	} else if action == "cut" {
		CutUpFile(id)
	}

}

func CutUpFile(id string) {
	path := SetupYoutubeDirectory()
	seconds := 60
	for {
		name := fmt.Sprintf("%s/%s_%05d_%05d.mp4", path, id, seconds-60, seconds)
		cmd := exec.Command("ffmpeg", "-i", path+"/"+id+".mp4",
			"-ss", fmt.Sprintf("%d", seconds-60),
			"-t", fmt.Sprintf("%d", 60),
			"-c", "copy",
			name)

		fmt.Println(name)
		cmd.Run()
		fi, _ := os.Stat(name)
		size := fi.Size()
		fmt.Println(size)
		if size < 1000 {
			break
		}

		seconds += 50
	}
}

func DownloadYoutube(id string) {
	path := SetupYoutubeDirectory()

	cmd := exec.Command("youtube-dl", "--output",
		path+"/%(id)s.%(ext)s",
		"--recode-video", "mp4", id)

	go cmd.Run()

	for {
		PrintDirectoryInfo(path)
		time.Sleep(time.Second * 5)
	}
}

func PrintDirectoryInfo(path string) {
	b, _ := exec.Command("ls", "-lh", path).CombinedOutput()
	s := string(b)
	fmt.Println(s)
	fmt.Println("")
}
