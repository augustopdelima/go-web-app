package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"resume-web-app/app"
	"resume-web-app/db"
	"resume-web-app/middleware"
	"resume-web-app/models/sqlite"
)

const PORT = ":8080"

//go:embed templates
var templates embed.FS

// ...

var tmpl = template.Must(template.ParseFS(templates, "templates/*.html"))

func routes(env *app.Env) http.Handler {
	router := http.NewServeMux()

	registerPageHandler := middleware.RegisterPage(tmpl)
	listResumes := middleware.ListResumes(tmpl, env)
	registerResumes := middleware.RegisterResumes(tmpl, env)
	detailResumes := middleware.DetailResume(tmpl, env)

	router.HandleFunc("GET /", middleware.SecureHeaders(listResumes))
	router.HandleFunc("GET /register", middleware.SecureHeaders(registerPageHandler))
	router.HandleFunc("POST /register", middleware.SecureHeaders(registerResumes))
	router.HandleFunc("GET /detail/{id}", middleware.SecureHeaders(detailResumes))

	return router
}

func main() {
	database := db.InitDatabase()
	defer database.Close()

	app := app.Env{
		Resume: &sqlite.ResumeModel{
			DB: database,
		},
	}

	server := http.Server{
		Addr:    PORT,
		Handler: routes(&app),
	}

	fmt.Println("Server listen in http://localhost:8080/")
	server.ListenAndServe()
}
