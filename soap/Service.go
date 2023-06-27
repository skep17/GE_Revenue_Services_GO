package soap

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"io"
	"log"
	"net/http"
)

type SoapEnvelope struct {
	XMLName xml.Name `xml:"soapenv:Envelope"`
	SoapEnv string   `xml:"xmlns:soapenv,attr"`
	Tem     string   `xml:"xmlns:tem,attr"`
	Body    struct {
		RequestStruct any
	} `xml:"soapenv:Body"`
}

type RawData struct {
	Request  string
	Response string
}

func makeBody(reqBody any) ([]byte, error) {
	env := &SoapEnvelope{SoapEnv: "http://schemas.xmlsoap.org/soap/envelope/", Tem: "http://tempuri.org/", Body: struct{ RequestStruct any }{RequestStruct: reqBody}}
	output, err := xml.MarshalIndent(env, "", "  ")

	return output, err
}

func readResponse(data *io.ReadCloser) ([]byte, error) {
	result, err := io.ReadAll(*data)
	*data = io.NopCloser(bytes.NewBuffer(result))

	return result, err
}

func SendRequest(reqBody, resBody any, url string) RawData {
	var result RawData

	rawRequest, err := makeBody(reqBody)
	if err != nil {
		log.Fatal("Error on encoding request body. ", err.Error())
	}

	result.Request = string(rawRequest)

	req, err := http.NewRequest("POST", url, bytes.NewReader(rawRequest))
	if err != nil {
		log.Fatal("Error on creating request object. ", err.Error())
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
	}

	rawResponse, err := readResponse(&res.Body)

	result.Response = string(rawResponse)

	err = xml.NewDecoder(res.Body).Decode(resBody)
	if err != nil {
		log.Fatal("Error on unmarshaling xml. ", err.Error())
	}

	return result
}
