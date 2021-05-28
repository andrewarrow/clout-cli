package session

import (
	"clout/files"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

func BackupSecrets(phrase string) {
	shhh := JustReadFile(file)

	val := encrypt([]byte(shhh), phrase)
	val64 := base64.StdEncoding.EncodeToString(val)

	home := files.UserHomeDir()
	os.Mkdir(home+"/"+dir, 0700)
	path := home + "/" + dir + "/" + backup
	ioutil.WriteFile(path, []byte(val64), 0700)

	fmt.Println("Backup writen to", path)
}
