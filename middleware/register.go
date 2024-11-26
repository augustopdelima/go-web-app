package middleware

import (
	"html"
	"html/template"
	"net/http"
	"resume-web-app/app"
	"resume-web-app/helpers"
	"resume-web-app/models"
	"strings"

	"github.com/gorilla/csrf"
)

func sanitizeInput(input string) string {
	removedWhiteSpace := strings.TrimSpace(input)

	return html.EscapeString(removedWhiteSpace)
}

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
			ShowErrorPage(tmpl, response, "Erro interno no servidor", http.StatusInternalServerError)
			return
		}

		resume := models.Resume{
			Name:       sanitizeInput(request.PostForm.Get("name")),
			Email:      sanitizeInput(request.PostForm.Get("email")),
			Cellphone:  sanitizeInput(request.PostForm.Get("cellphone")),
			WebAddress: sanitizeInput(request.PostForm.Get("web")),
			Experience: sanitizeInput(request.PostForm.Get("experience")),
		}

		if resume.Name == "" {
			ShowErrorPage(tmpl, response, "Nome é necessário", http.StatusBadRequest)
			return
		}

		if resume.Experience == "" {
			ShowErrorPage(tmpl, response, "Experiência é necessária", http.StatusBadRequest)
			return
		}

		validEmail := helpers.ValidateEmail(resume.Email)

		if !validEmail {
			ShowErrorPage(tmpl, response, "Email deve ser válido", http.StatusBadRequest)
			return
		}

		if resume.WebAddress != "" {

			validWebAddress := helpers.ValidateUrl(resume.WebAddress)

			if !validWebAddress {
				ShowErrorPage(tmpl, response, "Endereço web deve ser válido", http.StatusBadRequest)
				return
			}
		}

		if resume.Cellphone != "" {

			validCellphone := helpers.ValidateCellphoneNumber(resume.Cellphone)

			if !validCellphone {
				ShowErrorPage(tmpl, response, "Telefone deve ser válido", http.StatusBadRequest)
				return
			}

		}

		err = env.Resume.Insert(resume)

		if err != nil {
			ShowErrorPage(tmpl, response, "Não foi possível cadastrar o currículo", http.StatusInternalServerError)
			return
		}

		http.Redirect(response, request, "/", http.StatusMovedPermanently)
	})
}
