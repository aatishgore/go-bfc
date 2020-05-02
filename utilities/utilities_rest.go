package utilities

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// APIRequest is help to set a parameter
type APIRequest struct {
	APIURL  string
	Method  string
	Headers []string
	Data    []byte
}

// CallAPI is call to a third part Api
func CallAPI(request APIRequest) string {

	// hit thered party api
	log.Fatalln("Call instance api on : " + request.APIURL)
	req, err := http.NewRequest(request.Method, request.APIURL, bytes.NewBuffer(request.Data))

	// checking error
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}

	req.Header.Add("Authorization", basicAuth())
	// req.BasicAuth(InstanceAPIUserName,InstanceAPIPassword)

	for _, header := range request.Headers {
		headerPart := strings.Split(header, ":")
		req.Header.Set(headerPart[0], headerPart[1])
	}

	client := &http.Client{Timeout: time.Second * 20}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}

	return string(body)
}

func basicAuth() string {
	auth := InstanceAPIUserName + ":" + InstanceAPIPassword
	return auth
}
