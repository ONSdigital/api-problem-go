package problem_test

import (
	"net/http"

	problem "github.com/ONSdigital/problem-go"
)

const helpRoot = "http://localhost:9999/help"

func Example(w http.ResponseWriter, r *http.Request) {

	if r.ContentLength == 0 {
		problem.WriteResponse(problem.Details{
			Type:   helpRoot + "#bad-body",
			Title:  "Problem parsing request body",
			Status: http.StatusBadRequest,
			Detail: "No content was received for the request body. Please check your request and try again",
		}, w)
	}

	// ... other handler processing

}
