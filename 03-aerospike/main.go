package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

type user struct {
	Username string
	Password []byte
	ID       string
}

var tpl *template.Template

func main() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))

	f, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	log.SetOutput(f)

	router := httprouter.New()
	router.GET("/", index)
	router.GET("/login", loginPage)
	router.POST("/login", login)
	router.GET("/logout", logout)
	router.GET("/create", createPage)
	router.POST("/create", create)
	err = http.ListenAndServe(":9000", router)
	if err != nil {
		panic(err)
	}
}

func index(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
}

func loginPage(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
}

func login(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
}

func logout(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
}

func createPage(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
}

func create(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
}
