package main

import (
	"fmt"
	"html/template"
	"net/http"
	"resume-web-app/db"
	"resume-web-app/models"
	"resume-web-app/models/sqlite"
)

const PORT = ":8080"

type app struct {
	resume *sqlite.ResumeModel
}

var tmpl = template.Must(template.ParseFiles("templates/index.go.tmpl", "templates/detail.go.tmpl", "templates/register.go.tmpl", "templates/not-found.go.tmpl"))

func (app *app) listResumes(response http.ResponseWriter, request *http.Request) {

	resumes, err := app.resume.All()

	if err != nil {
		http.Error(response, "Internal server error", http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(response, "index.go.tmpl", resumes)
}

func (app *app) registerPage(response http.ResponseWriter, request *http.Request) {
	tmpl.ExecuteTemplate(response, "register.go.tmpl", nil)
}

func (app *app) registerResumes(response http.ResponseWriter, request *http.Request) {

	err := request.ParseForm()
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	resume := models.Resume{
		Name:       request.PostForm.Get("name"),
		Email:      request.PostForm.Get("email"),
		Cellphone:  request.PostForm.Get("cellphone"),
		WebAddress: request.PostForm.Get("web"),
		Experience: request.PostForm.Get("experience"),
	}

	err = app.resume.Insert(resume)

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(response, request, "/", http.StatusFound)

}

func (app *app) detailResume(response http.ResponseWriter, request *http.Request) {
	id := request.PathValue("id")

	resume, err := app.resume.SelectOne(id)

	if err != nil {
		tmpl.ExecuteTemplate(response, "not-found.go.tmpl", nil)
		return
	}

	tmpl.ExecuteTemplate(response, "detail.go.tmpl", resume)

}

func (app *app) routes() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("GET /", app.listResumes)
	router.HandleFunc("GET /register", app.registerPage)
	router.HandleFunc("POST /register", app.registerResumes)
	router.HandleFunc("GET /detail/{id}", app.detailResume)

	return router
}

func main() {
	database := db.InitDatabase()

	app := app{
		resume: &sqlite.ResumeModel{
			DB: database,
		},
	}

	server := http.Server{
		Addr:    PORT,
		Handler: app.routes(),
	}

	fmt.Println("Server listen in http://localhost:8080/")
	server.ListenAndServe()
}
