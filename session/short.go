package session

import (
	"clout/files"
	"encoding/json"
	"io/ioutil"
	"os"
)

func SaveShortMap(shortMap map[string]string) {
	existing := ReadShortMap()
	for k, v := range shortMap {
		existing[k] = v
	}
	b, _ := json.Marshal(existing)
	home := files.UserHomeDir()
	os.Mkdir(home+"/"+dir, 0700)
	path := home + "/" + dir + "/" + short
	ioutil.WriteFile(path, b, 0700)
}
func ReadShortMap() map[string]string {
	m := map[string]string{}
	asBytes := []byte(JustReadFile(short))
	if len(asBytes) == 0 {
		return m
	}

	json.Unmarshal(asBytes, &m)

	return m
}
