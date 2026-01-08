package helpers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/api-service/models"
)

func GetIntQueryParam(r *http.Request, param string) (int, error) {
	query := r.URL.Query()
	value := query.Get(param)
	if value == "" {
		return 0, errors.New("query parameter not found")
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	return intValue, nil
}

func GetBoolQueryParam(r *http.Request, param string) (bool, error) {
	query := r.URL.Query()
	value := query.Get(param)
	if value == "" {
		return false, errors.New("query parameter not found")
	}
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return false, err
	}
	return boolValue, nil
}

func ParseJSONBody(w http.ResponseWriter, r *http.Request, target interface{}) error {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "invalid content type", http.StatusBadRequest)
		return errors.New("invalid content type")
	}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(target)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	return nil
}

func HandleError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	log.Println(err)
}

func SplitAndTrim(s string, sep string) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		return strings.ContainsRune(sep, r)
	})
}

func ValidateUser(user *models.User) error {
	if user.Username == "" || user.Email == "" {
		return errors.New("username and email are required")
	}
	return nil
}