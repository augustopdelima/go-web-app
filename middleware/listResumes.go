package middleware

import (
	"html/template"
	"net/http"
	"resume-web-app/app"
)

func ListResumes(tmpl *template.Template, env *app.Env) http.HandlerFunc {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		path := request.URL.Path

		if path != "/" {
			response.WriteHeader(http.StatusNotFound)
			tmpl.ExecuteTemplate(response, "not-found.html", nil)
			return
		}

		resumes, err := env.Resume.All()

		if err != nil {
			ShowErrorPage(tmpl, response, "Erro interno no servidor", http.StatusInternalServerError)
			return
		}

		tmpl.ExecuteTemplate(response, "index.html", resumes)
	})
}
