package dragonformula

import (
	"appengine"
	"appengine/datastore"
	"github.com/martini-contrib/render"
	"net/http"
	"net/url"
)

/**
 カテゴリの検索
 typeとparentKeyNameで検索可能
 */
func GetCategoryList(r render.Render, req *http.Request) {
	c := appengine.NewContext(req)
	u, _ := url.Parse(req.URL.String())
	params := u.Query()
	q := datastore.NewQuery("Category")
	if len(params["type"]) != 0 {
		q = q.Filter("Type=", ToInt(params["type"][0]))
	}
	if len(params["parentKeyName"]) != 0 {
		q = q.Filter("parentKeyName=", params["parentKeyName"][0])
	}
	categories := make([]Category, 0, 10)
	_, err := q.GetAll(c, &categories)
	if err != nil {
		c.Criticalf(err.Error())
		r.JSON(400, err)
		return
	}
	r.JSON(200, categories)
}

/**
 カテゴリの登録
 */
func RegistCategory(r render.Render, req *http.Request) {
	c := appengine.NewContext(req)
	category := &Category{}
	key := datastore.NewKey(c, "Category", req.FormValue("keyName"), 0, nil)
	category.Name = req.FormValue("name")
	category.Type = ToInt(req.FormValue("type"))
	category.KeyName = key.StringID()
	category.ParentKeyName = req.FormValue("parentKeyName")
	key, err := datastore.Put(c, key, category)
	if err != nil {
		c.Criticalf("%s", err)
		r.JSON(400, err)
	} else {
		category.KeyName = key.StringID()
		r.JSON(200, category)
	}
}
