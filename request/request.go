package request

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetFormValue(r *http.Request, key string) string {
	r.ParseForm()
	return r.Form.Get(key)
}

func GetFormValues(r *http.Request, key string) []string {
	r.ParseForm()
	return r.Form[key]
}

func GetFormValueBool(r *http.Request, key string) bool {
	r.ParseForm()
	boolS := strings.ToLower(r.Form.Get(key))
	if boolS == "true" ||
		boolS == "t" ||
		boolS == "yes" ||
		boolS == "y" {
		return true
	}

	return false
}

func TryDecodeBody(w http.ResponseWriter, r *http.Request, target interface{}) ([]byte, error) {
	if r.Body == nil {
		return []byte{}, errors.New("empty body request")
	}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		return b, err
	}

	if b == nil {
		return b, errors.New("empty body request")
	}

	if err := json.Unmarshal(b, target); err != nil {
		return b, err
	}

	return b, nil
}
