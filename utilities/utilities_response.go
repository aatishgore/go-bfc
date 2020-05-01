package utilities

import (
	"encoding/json"
)

// Response is ...
type Response struct {
	status bool
	data   string
}

// ResponseEncode is ...
func ResponseEncode(status bool, data string) string {
	jsonResponseData, err := json.Marshal(&Response{status: status, data: data})

	// if unable to encode into a json then send default error json
	if err != nil {
		jsonResponseData, _ = json.Marshal(&Response{status: false, data: ""})
	}

	println(jsonResponseData)

	return string(jsonResponseData)
}
