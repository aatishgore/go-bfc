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

	response := Response{status: status, data: data}

	jsonResponseData, err := json.Marshal(response)

	// if unable to encode into a json then send default error json
	if err != nil {
		jsonResponseData, _ = json.Marshal(Response{status: false, data: ""})
	}

	return string(jsonResponseData)
}
