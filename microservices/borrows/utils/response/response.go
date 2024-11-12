package response

import (
	"encoding/json"
	"net/http"

	"github.com/synapsis-library-management-server/microservices/borrows/utils/failure"
)

// Base is the base object of all responses
type Base struct {
	Data     interface{} `json:"data,omitempty"`
	Metadata interface{} `json:"metadata,omitempty"`
	Message  string      `json:"message,omitempty"`
}

// WithData sends a response containing a JSON object
func WithData(w http.ResponseWriter, code int, data interface{}) {
	respond(w, code, Base{Data: data})
}

// WithMetadata sends a response containing a JSON object with metadata
func WithMetadata(w http.ResponseWriter, code int, data interface{}, metadata interface{}) {
	respond(w, code, Base{Data: data, Metadata: metadata})
}

// WithMessage sends a response with a simple text message
func WithMessage(w http.ResponseWriter, code int, message string) {
	respond(w, code, Base{Message: message})
}

// WithError sends a response with an error message
func WithError(w http.ResponseWriter, err error) {
	code := failure.GetCode(err)
	errMsg := err.Error()
	respond(w, code, Base{Message: errMsg})
}

func respond(w http.ResponseWriter, code int, response Base) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}
