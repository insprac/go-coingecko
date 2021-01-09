package v3

import (
	"encoding/json"
	"errors"
	"fmt"
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

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		if err == nil {
			return newError("bad response (status=%d, body=%v)",
				resp.StatusCode,
				string(bodyBytes))
		} else {
			return newError("bad response (status=%d)", resp.StatusCode)
		}
	}

	if err != nil {
		return err
	}

	return json.Unmarshal(bodyBytes, &data)
}

func newError(message string, args ...interface{}) error {
	return errors.New(fmt.Sprintf(message, args...))
}
