package utils

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Post(url string, param ...string) (string, error) {
	var reqBody io.Reader
	if len(param) > 0 {
		reqBody = strings.NewReader(param[0])
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		log.Println(err)
		return "", err
	}
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(body), nil
}

func Get(url string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return "", err
	}
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(body), nil
}
