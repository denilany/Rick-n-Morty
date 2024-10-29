package customerrors

import "net/http"

type ErrorPage struct {
	StatusCode int
	Message    string
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	ServeErrorPage(w, ErrorPage{
		StatusCode: http.StatusNotFound,
		Message:    "Page Not Found",
	})
}
