package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/golang/gddo/httputil/header"

	logger "github.com/sirupsen/logrus"
)

type malformedRequest struct {
	status       int
	errorMessage string
}

func (mr *malformedRequest) Error() string {
	return mr.errorMessage
}

// Request decoder and validation
func decodeRequest(w http.ResponseWriter, request *http.Request, dataStruct interface{}) {
	err := decodeJSONBody(w, request, dataStruct)
	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.errorMessage, mr.status)
		} else {
			log.Println(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func decodeToken(w http.ResponseWriter, request *http.Request) string {
	if request.Header.Get("Access-Token") != "" {
		value, _ := header.ParseValueAndParams(request.Header, "Access-Token")
		return value
	}
	return ""
}

/**
/ JSON Request Body Decoder
**/

func decodeJSONBody(w http.ResponseWriter, request *http.Request, dataStruct interface{}) error {

	enableCors(&w)

	if request.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(request.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			return &malformedRequest{status: http.StatusUnsupportedMediaType, errorMessage: msg}
		}
	}

	request.Body = http.MaxBytesReader(w, request.Body, 1048576)

	dec := json.NewDecoder(request.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&dataStruct)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			logger.Error(msg)
			return &malformedRequest{status: http.StatusBadRequest, errorMessage: msg}

		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := fmt.Sprintf("Request body contains badly-formed JSON")
			logger.Error(msg)
			return &malformedRequest{status: http.StatusBadRequest, errorMessage: msg}

		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			logger.Error(msg)
			return &malformedRequest{status: http.StatusBadRequest, errorMessage: msg}

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			logger.Error(msg)
			return &malformedRequest{status: http.StatusBadRequest, errorMessage: msg}

		case errors.Is(err, io.EOF):
			msg := "Request body must not be empty"
			logger.Error(msg)
			return &malformedRequest{status: http.StatusBadRequest, errorMessage: msg}

		case err.Error() == "http: request body too large":
			msg := "Request body must not be larger than 1MB"
			logger.Error(msg)
			return &malformedRequest{status: http.StatusRequestEntityTooLarge, errorMessage: msg}

		default:
			return err
		}
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		msg := "Request body must only contain a single JSON object"
		logger.Error(msg)
		return &malformedRequest{status: http.StatusBadRequest, errorMessage: msg}
	}

	return nil
}

/**
/ Response Encoder
**/

func EncodeResponse(w http.ResponseWriter, r *http.Request, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}
