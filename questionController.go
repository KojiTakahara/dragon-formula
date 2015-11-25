package dragonformula

import (
	"appengine"
	"appengine/datastore"
	"github.com/martini-contrib/render"
	"math/rand"
	"net/http"
	"net/url"
)

/**
問題選択肢の検索
*/
func GetQuestionList(r render.Render, req *http.Request) {
	c := appengine.NewContext(req)
	u, _ := url.Parse(req.URL.String())
	p := u.Query()
	q := datastore.NewQuery("Question")
	if len(p["status"]) != 0 { // ステータスフィルター
		q = q.Filter("Status", p["status"][0])
	}
	if len(p["largeCategoryKey"]) != 0 { // 大項目フィルター
		q = q.Filter("LargeCategoryKey=", p["largeCategoryKey"][0])
	}
	questions := make([]Question, 0, 10)
	keys, err := q.GetAll(c, &questions)
	if err != nil {
		c.Criticalf(err.Error())
		r.JSON(400, err)
		return
	}
	for i := 0; i < len(questions); i++ {
		question := questions[i]
		q := datastore.NewQuery("QuestionChoice")
		q = q.Filter("QuestionKey=", keys[i].StringID())
		choices := make([]QuestionChoice, 0, 4)
		_, err := q.GetAll(c, &choices)
		if err != nil {
			c.Criticalf(err.Error())
			r.JSON(400, err)
			return
		}
		shuffleQuestionChoice(choices)

	}
	shuffleQuestion(questions)
	r.JSON(200, questions)
}

func shuffleQuestion(data []Question) {
	n := len(data)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
}

func shuffleQuestionChoice(data []QuestionChoice) {
	n := len(data)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
}

/**
問題の登録
*/
func RegistQuestion(r render.Render, req *http.Request) {
	c := appengine.NewContext(req)
	category := &Question{}
	key := datastore.NewKey(c, "Question", req.FormValue("key"), 0, nil)
	// category.Name = req.FormValue("name")
	// category.Type = ToInt(req.FormValue("type"))
	// category.Key = key.StringID()
	// category.ParentKey = req.FormValue("parentKey")
	key, err := datastore.Put(c, key, category)
	if err != nil {
		c.Criticalf("%s", err)
		r.JSON(400, err)
	} else {
		//category.Key = key.StringID()
		r.JSON(200, category)
	}
}
