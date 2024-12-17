package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string      `json:"status"`
	Error  string      `json:"error"`
	Data   interface{} `json:"data"`
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func SuccessResponse(data interface{}) Response {
	return Response{
		Status: StatusOK,
		Error:  "",
		Data:   data,
	}
}

func GenericError(err error) Response {
	return Response{
		Status: StatusError,
		Error:  err.Error(),
	}
}

func WriteJson(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func ValidationErrors(errs validator.ValidationErrors) Response {
	var errMsg []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsg = append(errMsg, fmt.Sprintf("field %s is required", err.Field()))
			fmt.Println(err.Tag())
			fmt.Println(err.Error())
		case "base64":
			errMsg = append(errMsg, fmt.Sprintf("field %s is invalid", err.Field()))

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
