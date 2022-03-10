package utils

import (
	"encoding/json"
	"net/http"
)

const successStatusText string = "success"
const errorStatusText string = "error"

type GenericSuccessResponse struct {
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type GenericErrorResponse struct {
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
	Error      string `json:"error"`
}

func JSONSuccess(w http.ResponseWriter, statusCode int, message string) error {
	return JSONWrite(w, statusCode, &GenericSuccessResponse{
		Status:     successStatusText,
		StatusCode: statusCode,
		Message:    message,
	})
}

func JSONError(w http.ResponseWriter, statusCode int, err string) error {
	return JSONWrite(w, statusCode, &GenericErrorResponse{
		Status:     errorStatusText,
		StatusCode: statusCode,
		Error:      err,
	})
}

func JSONWrite(w http.ResponseWriter, statusCode int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)

	byteData, err := json.Marshal(data)
	if err != nil {
		JSONError(w, http.StatusInternalServerError, err.Error())
		return err
	}
	w.Write(byteData)
	return nil
}
