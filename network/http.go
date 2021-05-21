package network

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func BaseUrl() string {
	return "http://wolf-4849.wolfschedule.com:17001/"
	//return "https://bitclout.com/"
}

func DoGet(route string) string {
	agent := "agent"

	urlString := fmt.Sprintf("%s%s", BaseUrl(), route)
	request, _ := http.NewRequest("GET", urlString, nil)
	request.Header.Set("User-Agent", agent)
	request.Header.Set("Content-Type", "application/json")
	//request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", pat))
	client := &http.Client{Timeout: time.Second * 500}
	return DoHttpRead("GET", route, client, request)
}

func DoHttpRead(verb, route string, client *http.Client, request *http.Request) string {
	resp, err := client.Do(request)
	if err == nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("\n\nERROR: %d %s\n\n", resp.StatusCode, err.Error())
			return ""
		}
		if resp.StatusCode == 200 || resp.StatusCode == 201 || resp.StatusCode == 204 {
			return string(body)
		} else {
			fmt.Printf("\n\nERROR: %d %s\n\n", resp.StatusCode, string(body))
			return ""
		}
	}
	fmt.Printf("\n\nERROR: %s\n\n", err.Error())
	return ""
}

func DoPost(route string, payload []byte) string {
	body := bytes.NewBuffer(payload)
	urlString := fmt.Sprintf("%s%s", BaseUrl(), route)
	request, _ := http.NewRequest("POST", urlString, body)
	request.Header.Set("Content-Type", "application/json")
	//request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", pat))
	client := &http.Client{Timeout: time.Second * 50}

	return DoHttpRead("POST", route, client, request)
}