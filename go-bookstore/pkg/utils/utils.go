package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(request *http.Request, store interface{}) {
	if body, err := ioutil.ReadAll(request.Body); err == nil {
		if err := json.Unmarshal([]byte(body), store); err != nil {
			return
		}
	}
}
