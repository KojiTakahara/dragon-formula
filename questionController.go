package dragonformula

import (
	"appengine"
	"appengine/datastore"
	"fmt"
	"github.com/martini-contrib/render"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
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
		q = q.Filter("Status=", p["status"][0])
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
		q := datastore.NewQuery("QuestionChoice")
		q = q.Filter("QuestionKeyId =", keys[i].IntID())
		choices := make([]QuestionChoice, 0, 3)
		choicesKeys, err := q.GetAll(c, &choices)
		if err != nil {
			c.Criticalf(err.Error())
			r.JSON(400, err)
			return
		}
		if len(choices) != 3 {
			r.JSON(400, fmt.Sprintf("size error. choices = %d", len(choices)))
			return
		}
		for i := 0; i < len(choices); i++ {
			choices[i].Key = choicesKeys[i].IntID()
		}
		shuffleQuestionChoice(choices)
		questions[i].Choice1 = choices[0]
		questions[i].Choice2 = choices[1]
		questions[i].Choice3 = choices[2]
		questions[i].Key = keys[i].IntID()
	}
	shuffleQuestion(questions)
	if len(p["limit"]) != 0 {
		limit := ToInt(p["limit"][0])
		questions = questions[0:limit]
	}
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
	question.UserKey = req.FormValue("userKey")
	// question.UserKey
	resultkey, err := datastore.Put(c, key, question)
	if err != nil {
		c.Criticalf("%s", err)
		r.JSON(400, err)
	} else {
		c.Infof("%d", resultkey.IntID())
		RegistQuestionChoice(req, resultkey.IntID(), "1")
		RegistQuestionChoice(req, resultkey.IntID(), "2")
		RegistQuestionChoice(req, resultkey.IntID(), "3")
		r.JSON(200, question)
	}
}

/**
問題ステータスの更新
**/
func UpdateQuestionStatus(r render.Render, req *http.Request) {
	c := appengine.NewContext(req)
	id, _ := strconv.Atoi(req.FormValue("key"))
	key := datastore.NewKey(c, "Question", "", int64(id), nil)
	var question Question
	if err := datastore.Get(c, key, &question); err != nil {
		c.Criticalf(err.Error())
	}
	question.Status = req.FormValue("status")
	_, err := datastore.Put(c, key, question)
	if err != nil {
		c.Criticalf("%s", err)
		r.JSON(400, err)
	} else {
		r.JSON(200, question)
	}
}

func RegistQuestionChoice(req *http.Request, k int64, num string) {
	c := appengine.NewContext(req)
	choice := &QuestionChoice{}
	key := datastore.NewKey(c, "QuestionChoice", req.FormValue("choice"+num+"Key"), 0, nil)
	choice.Content = req.FormValue("choice" + num + "Content")
	choice.TrueFalse, _ = strconv.ParseBool(req.FormValue("choice" + num + "Bool"))
	choice.QuestionKeyId = k
	_, err := datastore.Put(c, key, choice)
	if err != nil {
		c.Criticalf("%s", err)
	}
}
