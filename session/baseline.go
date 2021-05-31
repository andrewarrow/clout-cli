package session

import (
	"clout/files"
	"encoding/json"
	"io/ioutil"
	"os"
)

func SaveBaselineNotifications(save map[string]map[string]int) {
	b, _ := json.Marshal(save)
	home := files.UserHomeDir()
	os.Mkdir(home+"/"+dir, 0700)
	path := home + "/" + dir + "/" + baseline
	ioutil.WriteFile(path, b, 0700)
}
func ReadBaseline() map[string]map[string]int {
	m := map[string]map[string]int{}
	asBytes := []byte(JustReadFile(baseline))
	if len(asBytes) == 0 {
		return m
	}

	json.Unmarshal(asBytes, &m)

	return m
}
