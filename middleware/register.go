package middleware

import (
	"html"
	"html/template"
	"net/http"
	"resume-web-app/app"
	"resume-web-app/models"
)

func RegisterPage(tmpl *template.Template) http.HandlerFunc {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		tmpl.ExecuteTemplate(response, "register.html", nil)
	})
}

func escapeString(value string) string {
	return html.EscapeString(value)
}

func RegisterResumes(tmpl *template.Template, env *app.Env) http.HandlerFunc {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		err := request.ParseForm()

		if err != nil {
			http.Error(response, err.Error(), http.StatusBadRequest)
			return
		}

		resume := models.Resume{
			Name:       escapeString(request.PostForm.Get("name")),
			Email:      escapeString(request.PostForm.Get("email")),
			Cellphone:  escapeString(request.PostForm.Get("cellphone")),
			WebAddress: escapeString(request.PostForm.Get("web")),
			Experience: escapeString(request.PostForm.Get("experience")),
		}

		err = env.Resume.Insert(resume)

		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(response, request, "/", http.StatusFound)
	})
}
