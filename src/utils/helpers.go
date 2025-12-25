package helpers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/api-service/models"
)

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func ValidateRequest(r *http.Request, expectedMethods []string) error {
	method := r.Method
	for _, m := range expectedMethods {
		if strings.EqualFold(method, m) {
			return nil
		}
	}
	return errors.New("invalid request method")
}

func GetRequestBody(r *http.Request, target interface{}) error {
	if r.Header.Get("Content-Type") != "application/json" {
		return errors.New("invalid content type")
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(target)
	if err != nil {
		log.Printf("error decoding request body: %v", err)
		return err
	}
	return nil
}

func ParseModelFromRequest(r *http.Request) (*models.RequestModel, error) {
	var model models.RequestModel
	err := GetRequestBody(r, &model)
	if err != nil {
		return nil, err
	}
	return &model, nil
}