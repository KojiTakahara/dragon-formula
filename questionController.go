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
	if len(p["userKey"]) != 0 { // userKeyフィルター
		q = q.Filter("UserKey=", p["userKey"][0])
	}
	questions := make([]Question, 0, 10)
	keys, err := q.GetAll(c, &questions)
	if err != nil {
		c.Criticalf(err.Error())
		r.JSON(400, err)
		return
	}
	for i := 0; i < len(questions); i++ {
		// Get Choice
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
			// TODO ○×に合わせて修正する
			c.Errorf(fmt.Sprintf("size error. choices = %d", len(choices)))
			// r.JSON(400, fmt.Sprintf("size error. choices = %d", len(choices)))
			// return
		}
		for i := 0; i < len(choices); i++ {
			choices[i].Key = choicesKeys[i].IntID()
		}
		if questions[i].LargeCategoryKey != "rule_3" {
			shuffleQuestionChoice(choices)
		}
		if 3 <= len(choices) {
			questions[i].Choice3 = choices[2]
		}
		if 2 <= len(choices) {
			questions[i].Choice2 = choices[1]
		}
		if 1 <= len(choices) {
			questions[i].Choice1 = choices[0]
		}

		// Get Annotations
		q = datastore.NewQuery("QuestionAnnotation")
		if len(p["questionKeyId"]) != 0 {
			q = q.Filter("QuestionKeyId=", keys[i].IntID())
		}
		qaList := make([]QuestionAnnotation, 0, 10)
		annotationkeys, err := q.GetAll(c, &qaList)
		if err != nil {
			c.Criticalf(err.Error())
		}
		for i := range qaList {
			qaList[i].Key = annotationkeys[i].IntID()
		}
		questions[i].Annotations = qaList

		questions[i].Key = keys[i].IntID()
	}
	shuffleQuestion(questions)
	if len(p["limit"]) != 0 {
		limit := ToInt(p["limit"][0])
		c.Infof(fmt.Sprintf("limit = %d", limit))
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
	question.Status = req.FormValue("status")
	question.LargeCategoryKey = req.FormValue("largeCategoryKey")
	question.MediumCategoryKey = req.FormValue("mediumCategoryKey")
	question.SmallCategoryKey = req.FormValue("smallCategoryKey")
	question.Rubric = req.FormValue("rubric")
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
更新
**/
func UpdateQuestion(r render.Render, req *http.Request) {
	c := appengine.NewContext(req)
	id, _ := strconv.Atoi(req.FormValue("Key"))
	key := datastore.NewKey(c, "Question", "", int64(id), nil)
	var question Question
	if err := datastore.Get(c, key, &question); err != nil {
		c.Criticalf(err.Error())
	}
	question.Content = req.FormValue("Content")
	question.LargeCategoryKey = req.FormValue("LargeCategoryKey")
	question.MediumCategoryKey = req.FormValue("MediumCategoryKey")
	question.SmallCategoryKey = req.FormValue("SmallCategoryKey")
	question.Rubric = req.FormValue("Rubric")
	question.Percentage, _ = strconv.ParseFloat(req.FormValue("Percentage"), 64)
	question.Status = req.FormValue("Status")
	question.Level = req.FormValue("Level")
	question.UserKey = req.FormValue("UserKey")
	_, err := datastore.Put(c, key, &question)
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
