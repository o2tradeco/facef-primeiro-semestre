package main

import (
    "net/http"
	"fmt"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

type App struct {
	router *chi.Mux
}

func NewApp() *App {
	return &App{}
}

func (app *App) initRouter() {
	app.router = chi.NewRouter()
}

func main() {
	app := NewApp()
	app.initRouter()
	fmt.Printf("%T\n", app.router)
	app.initRoutes()
	fmt.Printf("%T\n", app.router)
	app.initHttpServer()
}

func (app *App) initRoutes() {
	app.router.Use(middleware.Logger)
	app.router.Get("/cartao", app.Handler)
	app.router.Get("/style", app.StaticRoutes)
}

func (app *App) initHttpServer() {
	fmt.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", app.router)
}

func (app *App) Handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web-es/cartao.html")
}

func (app *App) StaticRoutes(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web-es/css/stylesheet.css")
}
