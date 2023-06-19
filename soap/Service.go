package soap

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"log"
	"net/http"
)

func SendRequest(reqBody []byte, resBody any, url string) string {
	req, err := http.NewRequest("POST", url, bytes.NewReader(reqBody))
	if err != nil {
		log.Fatal("Error on creating request object. ", err.Error())
		return ""
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
		return ""
	}

	temp := res.Body

	err = xml.NewDecoder(res.Body).Decode(resBody)
	if err != nil {
		log.Fatal("Error on unmarshaling xml. ", err.Error())
		return ""
	}

	resByteData := new(bytes.Buffer)

	resByteData.ReadFrom(temp)

	resRaw := resByteData.String()

	return resRaw
}
