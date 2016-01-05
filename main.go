package dragonformula

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
)

var m *martini.Martini

func init() {
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Get("/api/question", GetQuestionList)
	m.Post("/api/question", RegistQuestion)
	m.Get("/api/category", GetCategoryList)
	m.Post("/api/category", RegistCategory)
	m.Get("/api/loginUser", LoginUser)
	m.Get("/api/twitter/login", LoginTwitter)
	m.Get("/api/twitter/callback", CallbackTwitter)
	m.Get("/api/twitter/user/:userId", GetTwitterUserById)
	http.ListenAndServe(":8080", m)
	http.Handle("/", m)
}
