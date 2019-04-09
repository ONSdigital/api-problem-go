// Package problem provides a quick standard http package construct for returning
// API Problem Details objects to clients as specified in RFC7807 - https://tools.ietf.org/html/rfc7807
package problem

import (
	"encoding/json"
	"log"
	"net/http"
)

// Details represents a Problem Details object as specified in RFC7807
type Details struct {
	Type   string `json:"type,omitempty"`   // Link to a resource for the problem
	Title  string `json:"title,omitempty"`  // Short description of the issue
	Status int    `json:"status"`           // The http status code
	Detail string `json:"detail,omitempty"` // Further human-readable detail
}

// WriteResponse writes an API problem response report to the given ResponseWriter.
// If for some reason it fails to marshal the json response, it returns a 500
// internal error.
func WriteResponse(details Details, rw http.ResponseWriter) {
	pr, err := json.Marshal(&details)
	if err != nil {
		log.Printf(`event="Error writing problem response" error="%v"`, err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	if details.Status == 0 {
		details.Status = http.StatusInternalServerError
	}
	rw.Header().Set("Content-Type", "application/problem+json")
	rw.Header().Set("Content-Language", "en")
	rw.WriteHeader(details.Status)
	if _, err = rw.Write(pr); err != nil {
		log.Printf(`event="Error writing problem response" error="%v"`, err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}
