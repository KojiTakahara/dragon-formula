package dragonformula

import (
	"appengine"
	"appengine/datastore"
	"github.com/martini-contrib/render"
	"net/http"
	// "net/url"
)

func GetCategoryList(r render.Render, req *http.Request) {
	c := appengine.NewContext(req)
	// u, _ := url.Parse(req.URL.String())
	// params := u.Query()
	q := datastore.NewQuery("Category")
	categories := make([]Category, 0, 10)
	_, err := q.GetAll(c, &categories)
	if err != nil {
		c.Criticalf(err.Error())
		r.JSON(400, err)
		return
	}
	r.JSON(200, categories)
}

func RegistCategory(r render.Render, req *http.Request) {
	c := appengine.NewContext(req)
	category := &Category{}
	category.Name = req.FormValue("name")
	category.Type = ToInt(req.FormValue("type"))
	key := datastore.NewKey(c, "Category", req.FormValue("keyName"), 0, nil)
	key, err := datastore.Put(c, key, category)
	if err != nil {
		c.Criticalf("%s", err)
		r.JSON(400, err)
	} else {
		r.JSON(200, category)
	}
}
