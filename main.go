package dragonformula

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"net/http"
)

var m *martini.Martini
var SESSION_KEY = "teacher_session"

func init() {
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Use(sessions.Sessions(SESSION_KEY, sessions.NewCookieStore([]byte(SESSION_KEY))))
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
