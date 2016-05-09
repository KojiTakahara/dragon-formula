package dragonformula

import (
	"appengine"
	"appengine/datastore"
	"github.com/martini-contrib/render"
	"net/http"
	"net/url"
  "strconv"
)

/**
 * 検索
 */
func GetQuestionAnnotationList(r render.Render, req *http.Request) {
  c := appengine.NewContext(req)
	u, _ := url.Parse(req.URL.String())
	p := u.Query()
	q := datastore.NewQuery("QuestionAnnotation")
  if len(p["questionKeyId"]) != 0 {
		q = q.Filter("QuestionKeyId=", ToInt64(p["questionKeyId"][0]))
	}
  qaList := make([]QuestionAnnotation, 0, 10)
	keys, err := q.GetAll(c, &qaList)
	if err != nil {
		c.Criticalf(err.Error())
		r.JSON(400, err)
		return
	}
	for i := range qaList {
		qaList[i].Key = keys[i].IntID()
	}
  r.JSON(200, qaList)
}

/**
 * 登録
 */
func RegistQuestionAnnotation(r render.Render, req *http.Request) {
	c := appengine.NewContext(req)
	a := &QuestionAnnotation{}
	key := datastore.NewKey(c, "QuestionAnnotation", req.FormValue("Key"), 0, nil)
  a.QuestionKeyId = ToInt64(req.FormValue("QuestionKeyId"))
  a.CardName = req.FormValue("CardName")
  a.Annotation = req.FormValue("Annotation")
  resultkey, err := datastore.Put(c, key, a)
  if err != nil {
		c.Criticalf("%s", err)
		r.JSON(400, err)
	} else {
		c.Infof("%d", resultkey.IntID())
		r.JSON(200, a)
	}
}

/**
 * 更新
 */
func UpdateQuestionAnnotation(r render.Render, req *http.Request) {
	c := appengine.NewContext(req)
	id, _ := strconv.Atoi(req.FormValue("Key"))
	key := datastore.NewKey(c, "QuestionAnnotation", "", int64(id), nil)
	var qa QuestionAnnotation
	if err := datastore.Get(c, key, &qa); err != nil {
		c.Criticalf(err.Error())
	}
	qa.Annotation = req.FormValue("Annotation")
  qa.CardName = req.FormValue("CardName")
    _, err := datastore.Put(c, key, &qa)
	if err != nil {
		c.Criticalf("%s", err)
		r.JSON(400, err)
	} else {
		r.JSON(200, qa)
	}
}

/**
 * 削除
 */
func DeleteQuestionAnnotation(r render.Render, req *http.Request) {
	c := appengine.NewContext(req)
	id, _ := strconv.Atoi(req.FormValue("Key"))
	key := datastore.NewKey(c, "QuestionAnnotation", "", int64(id), nil)
	err := datastore.Delete(c, key)
	if err != nil {
		r.JSON(400, "削除に失敗しました")
		return
	}
	r.JSON(200, "成功")
}