package httputils

import (
	"encoding/json"
	"net/http"
)

// WriteResponse is about writing response and set header to application/json.
func WriteResponse(w http.ResponseWriter, code int, data interface{}) {
	// set header and return the response
	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

// WriteHandlerError takes in Error and ResponseWriter
// and optionally logs and writes it to the ResponseWriter.
func WriteHandlerError(header http.Header, err *Error, w http.ResponseWriter) {
	if err.HTTPStatusCode == http.StatusInternalServerError {
		err.ErrorMessage = "an error occurred"
	}
	// We need to create a new protobuf error in order to
	// get a valid proto.Message.
	pbErr := Error{
		HTTPStatusCode: err.HTTPStatusCode,
		ErrorMessage:   err.ErrorMessage,
		ErrorMessages:  err.ErrorMessages,
	}

	WriteResponse(w, err.HTTPStatusCode, &pbErr)
}
