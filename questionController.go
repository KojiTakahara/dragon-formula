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
		question.Choice1 = choices[0].Content
		question.Choice2 = choices[1].Content
		question.Choice3 = choices[2].Content
		question.Choice4 = choices[3].Content
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
	question := &Question{}
	key := datastore.NewKey(c, "Question", req.FormValue("key"), 0, nil)
	question.Content = req.FormValue("content")
	question.Status = "REVIEW"
	question.LargeCategoryKey = req.FormValue("largeCategoryKey")
	question.MediumCategoryKey = req.FormValue("mediumCategoryKey")
	question.SmallCategoryKey = req.FormValue("smallCategoryKey")
	// question.UserKey
	key, err := datastore.Put(c, key, question)
	if err != nil {
		c.Criticalf("%s", err)
		r.JSON(400, err)
	} else {
		RegistQuestionChoice(r, req, key.StringID())
		r.JSON(200, question)
	}
}

func RegistQuestionChoice(r render.Render, req *http.Request, keyName string) {
	c := appengine.NewContext(req)
	choice := &QuestionChoice{}
	key := datastore.NewKey(c, "QuestionChoice", req.FormValue("key"), 0, nil)
	choice.Content = req.FormValue("content")
	choice.TrueFalse = false //ToBool req.FormValue("trueFalse")
	choice.QuestionKey = keyName
	key, err := datastore.Put(c, key, choice)
	if err == nil {
		//r.JSON(200, question)
	}
}
