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

func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	ServeErrorPage(w, ErrorPage{
		StatusCode: http.StatusMethodNotAllowed,
		Message:    "Method Not Allowed",
	})
}

func ForbiddenHandler(w http.ResponseWriter, r *http.Request) {
	ServeErrorPage(w, ErrorPage{
		StatusCode: http.StatusForbidden,
		Message:    "Access Forbidden",
	})
}

func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	ServeErrorPage(w, ErrorPage{
		StatusCode: http.StatusInternalServerError,
		Message:    "An Unexpected Error Occurred. Try Again Later",
	})
}
