package session

import (
	"clout/files"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

func SecretsFromBackup(phrase string) {
	enc := JustReadFile(backup)
	decodedBytes, _ := base64.StdEncoding.DecodeString(enc)
	shhh := decrypt(decodedBytes, phrase)

	home := files.UserHomeDir()
	os.Mkdir(home+"/"+dir, 0700)
	path := home + "/" + dir + "/" + file
	ioutil.WriteFile(path, shhh, 0700)
}

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
