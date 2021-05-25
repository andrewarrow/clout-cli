package main

import (
	"clout/files"
	"encoding/json"
	"io/ioutil"
	"os"
)

func ReadCache() map[string]string {
	m := map[string]string{}
	asBytes := []byte(JustReadFile(cache))
	if len(asBytes) == 0 {
		return m
	}

	json.Unmarshal(asBytes, &m)

	return m
}

func WriteCache(m map[string]string) {
	b, _ := json.Marshal(m)
	home := files.UserHomeDir()
	os.Mkdir(home+"/"+dir, 0700)
	path := home + "/" + dir + "/" + cache
	ioutil.WriteFile(path, b, 0700)
}
