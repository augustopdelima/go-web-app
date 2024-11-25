package middleware

import (
	"html/template"
	"net/http"
)

type ErrorData struct {
	ErrorMessage string
}

func ShowErrorPage(tmpl *template.Template, response http.ResponseWriter, message string, statusCode int) {
	error := ErrorData{
		ErrorMessage: message,
	}
	response.WriteHeader(statusCode)
	tmpl.ExecuteTemplate(response, "error.html", error)
}

func ShowError(tmpl *template.Template, message string) http.HandlerFunc {
	return http.HandlerFunc(func(response http.ResponseWriter, r *http.Request) {
		ShowErrorPage(tmpl, response, message, http.StatusBadRequest)
	})

}
