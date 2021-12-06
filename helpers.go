package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func getDayInput(day int) (string, error) {
	cookie, err := getCookie()
	if err != nil {
		return "", nil
	}
	url := fmt.Sprintf("https://adventofcode.com/2021/day/%d/input", day)
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Cookie", cookie)
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	responseBody := string(body)
	return responseBody, nil
}

func getCookie() (string, error) {
	content, err := ioutil.ReadFile("cookie.txt")
	if err != nil {
		return "", err
	}
	return string(content), nil
}
