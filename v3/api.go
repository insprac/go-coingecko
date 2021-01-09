package v3

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const baseURL = "https://api.coingecko.com/api/v3"

func get(path string, data interface{}) error {
	resp, err := http.Get(baseURL + path)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("Bad response status code")
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	return json.Unmarshal(bodyBytes, &data)
}
