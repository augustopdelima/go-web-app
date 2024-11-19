package middleware

import (
	"html/template"
	"net/http"
	"resume-web-app/app"
)

func DetailResume(tmpl *template.Template, env *app.Env) http.HandlerFunc {
	return http.HandlerFunc(
		func(response http.ResponseWriter, request *http.Request) {
			id := request.PathValue("id")

			resume, err := env.Resume.SelectOne(id)

			if err != nil {
				tmpl.ExecuteTemplate(response, "not-found.html", nil)
				return
			}

			tmpl.ExecuteTemplate(response, "detail.html", resume)
		},
	)
}
