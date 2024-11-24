package middleware

import (
	"html"
	"html/template"
	"net/http"
	"resume-web-app/app"
	"resume-web-app/helpers"
	"resume-web-app/models"

	"github.com/gorilla/csrf"
)

func RegisterPage(tmpl *template.Template) http.HandlerFunc {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {

		csrfField := csrf.TemplateField(request)

		data := map[string]interface{}{
			"csrfField": csrfField,
		}

		tmpl.ExecuteTemplate(response, "register.html", data)
	})
}

func RegisterResumes(tmpl *template.Template, env *app.Env) http.HandlerFunc {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		err := request.ParseForm()

		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		resume := models.Resume{
			Name:       html.EscapeString(request.PostForm.Get("name")),
			Email:      html.EscapeString(request.PostForm.Get("email")),
			Cellphone:  html.EscapeString(request.PostForm.Get("cellphone")),
			WebAddress: html.EscapeString(request.PostForm.Get("web")),
			Experience: html.EscapeString(request.PostForm.Get("experience")),
		}

		if resume.Name == "" {
			http.Error(response, "Name is required", http.StatusBadRequest)
			return
		}

		if resume.Experience == "" {
			http.Error(response, "Experience is required", http.StatusBadRequest)
			return
		}

		validEmail := helpers.ValidateEmail(resume.Email)

		if !validEmail {
			http.Error(response, "Email must be valid", http.StatusBadRequest)
			return
		}

		if resume.WebAddress != "" {

			validWebAddress := helpers.ValidateUrl(resume.WebAddress)

			if !validWebAddress {
				http.Error(response, "Web Address must be valid", http.StatusBadRequest)
				return
			}
		}

		if resume.Cellphone != "" {

			validCellphone := helpers.ValidateCellphoneNumber(resume.Cellphone)

			if !validCellphone {
				http.Error(response, "Cellphone number must be valid", http.StatusBadRequest)
				return
			}

		}

		err = env.Resume.Insert(resume)

		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(response, request, "/", http.StatusMovedPermanently)
	})
}
