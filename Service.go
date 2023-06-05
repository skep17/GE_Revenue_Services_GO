package main

import (
	"bytes"
	"crypto/tls"
	"log"
	"net/http"
)

var url string

func initServicePath(path string) {
	url = path
}

func sendRequest(ser []byte) *http.Response {
	req, err := http.NewRequest("POST", url, bytes.NewReader(ser))
	if err != nil {
		log.Fatal("Error on creating request object. ", err.Error())
		return nil
	}
	req.Header.Set("Content-type", "text/xml")

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Error on dispatching request. ", err.Error())
		return nil
	}

	return res
}
