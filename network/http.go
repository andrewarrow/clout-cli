package network

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func BaseUrl() string {
	//return "http://hacknode.io/"
	return "https://bitclout.com/"
}

func DoTest404(word string) string {
	jsonString := `{"PublicKeyBase58Check":"","Username":"%s"}`
	sendString := fmt.Sprintf(jsonString, word)
	body := bytes.NewBuffer([]byte(sendString))
	urlString := fmt.Sprintf("%s%s", BaseUrl(), "api/v0/get-single-profile")
	request, _ := http.NewRequest("POST", urlString, body)
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: time.Second * 500}
	resp, err := client.Do(request)
	if err == nil {
		defer resp.Body.Close()
		if resp.StatusCode == 404 {
			return "404"
		} else {
			return "200"
		}
	}
	fmt.Printf("\n\nERROR: %s\n\n", err.Error())
	return ""
}
func DoGetWithPat(pat, url string) string {
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("bearer %s", pat))
	client := &http.Client{Timeout: time.Second * 500}
	return DoHttpRead("GET", "", client, request)
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
			text := string(body)
			if strings.Contains(text, "RuleErrorFollowEntryAlreadyExists") == false {
				fmt.Printf("\n\nERROR: %d %s\n\n", resp.StatusCode, string(body))
			}
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
	client := &http.Client{Timeout: time.Second * 50}

	return DoHttpRead("POST", route, client, request)
}
func DoPostMultipart(route, ct string, payload []byte) string {
	body := bytes.NewBuffer(payload)
	urlString := fmt.Sprintf("%s%s", BaseUrl(), route)
	request, _ := http.NewRequest("POST", urlString, body)
	request.Header.Set("Content-Type", ct)
	client := &http.Client{Timeout: time.Second * 50}

	return DoHttpRead("POST", route, client, request)
}
