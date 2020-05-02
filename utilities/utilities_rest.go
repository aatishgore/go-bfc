package utilities

import (
	"bytes"
	"fmt"
	"net/http"
)

// APIRequest is help to set a parameter
type APIRequest struct {
	APIURL  string
	Method  string
	Headers string
	Data    []byte
}

// CallAPI is call to a third part Api
func CallAPI(request APIRequest) string {
	var data string

	// hit thered party api
	response, err := http.Post(request.APIURL, request.Headers, bytes.NewBuffer(request.Data))

	// checking error
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}

	fmt.Printf("%T", response.Body)

	return data
}
