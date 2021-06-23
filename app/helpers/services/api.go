package services

import (
	"api-boilerplate/app/helpers"
	"api-boilerplate/domain"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func LicenseAPI(method, endpoint, payload, bearer string) (response domain.Response) {
	validation := make(map[string]interface{})

	req, err := http.NewRequest(method, os.Getenv("API_LICENSE")+endpoint, strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	if len(bearer) > 0 {
		req.Header.Set("Authorization", bearer)
	}
	if err != nil {
		return helpers.ErrorResponse(400, err.Error(), err, validation)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return helpers.ErrorResponse(400, err.Error(), err, validation)
	}

	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return helpers.ErrorResponse(400, err.Error(), err, validation)
	}

	err = json.Unmarshal(respByte, &response)
	if err != nil {
		return helpers.ErrorResponse(400, err.Error(), err, validation)
	}

	return response
}
