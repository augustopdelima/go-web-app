package middleware

import (
	"html/template"
	"net/http"
	"resume-web-app/app"
)

func notFound(tmpl *template.Template, response http.ResponseWriter) {
	tmpl.ExecuteTemplate(response, "not-found.html", nil)
}

func ListResumes(tmpl *template.Template, env *app.Env) http.HandlerFunc {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		path := request.URL.Path

		if path != "/" {
			notFound(tmpl, response)
			return
		}

		resumes, err := env.Resume.All()

		if err != nil {
			http.Error(response, "Internal server error", http.StatusInternalServerError)
			return
		}

		tmpl.ExecuteTemplate(response, "index.html", resumes)
	})
}
