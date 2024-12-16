package response

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func GenericError(err error) Response {
	return Response{
		Status: StatusError,
		Error:  err.Error(),
	}
}

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func ValidationErrors(errs validator.ValidationErrors) Response {
	var errMsg []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsg = append(errMsg, fmt.Sprintf("field %s is required", err.Field()))
			fmt.Println(err.Tag())
			fmt.Println(err.Error())
		case "url":
			errMsg = append(errMsg, fmt.Sprintf("url %s is invalid", err.Value()))

		default:
			errMsg = append(errMsg, fmt.Sprintf("field is invalid"))
		}
	}

	return Response{
		Status: StatusError,
		Error:  strings.Join(errMsg, ", "),
	}
}
