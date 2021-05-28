package session

import (
	"clout/files"
	"fmt"
	"os"
)

func Logout() {
	m := ReadAccounts()
	if len(m) == 0 {
		return
	}
	if len(m) == 1 {
		home := files.UserHomeDir()
		path := home + "/" + dir + "/" + file
		os.Remove(path)
		fmt.Println("Secret removed.")
		fmt.Println("")
		return
	}
	username := JustReadFile(selected)
	if username == "" {
		fmt.Println("Please run `clout account [username]` to select account first.")
		return
	}
	delete(m, username)
	WriteAccounts(m)
	m = ReadAccounts()
	for username, _ = range m {
		WriteSelected(username)
		break
	}
}
