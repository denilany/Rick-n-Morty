package customerrors

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type ErrorPage struct {
	StatusCode int
	Message    string
}

func ServeErrorPage(w http.ResponseWriter, errorPage ErrorPage) {
	tmplPath := filepath.Join("templates", "error-page.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "An Unexpected Error Occurred. Try Again Later", http.StatusInternalServerError)
		log.Printf("Error parsing template %s: %v\n", tmplPath, err)
		return
	}
	w.WriteHeader(errorPage.StatusCode)

	if err := tmpl.Execute(w, errorPage); err != nil {
		http.Error(w, "An Unexpected Error Occurred. Try Again Later", http.StatusInternalServerError)
		log.Printf("Error executing template: %v\n", err)
	}
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	ServeErrorPage(w, ErrorPage{
		StatusCode: http.StatusNotFound,
		Message:    "Not Found",
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

func BadRequestHandler(w http.ResponseWriter, r *http.Request) {
	ServeErrorPage(w, ErrorPage{
		StatusCode: http.StatusBadRequest,
		Message:    "Bad Request",
	})
}
