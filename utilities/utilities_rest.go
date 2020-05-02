package utilities

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strings"
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

	req.Header.Add("Authorization", "Basic "+basicAuth())
	// req.BasicAuth(InstanceAPIUserName,InstanceAPIPassword)

	for _, header := range request.Headers {
		headerPart := strings.Split(header, ":")
		req.Header.Set(headerPart[0], headerPart[1])
	}

	// data, err := ioutil.ReadAll(response.Body)
	data := ""
	return string(data)
}

func basicAuth() string {
	auth := InstanceAPIUserName + ":" + InstanceAPIPassword
	return auth
}
