package dragonformula

import (
	"appengine"
	"appengine/datastore"
	"github.com/martini-contrib/render"
	"net/http"
)

/**
 * ユーザの一覧取得
 */
func GetUserList(r render.Render, req *http.Request) {
	c := appengine.NewContext(req)
	q := datastore.NewQuery("User")
	q = q.Order("Key")
	categories := make([]Category, 0, 10)
	_, err := q.GetAll(c, &categories)
	if err != nil {
		c.Criticalf(err.Error())
		r.JSON(400, err)
		return
	}
	r.JSON(200, categories)
}