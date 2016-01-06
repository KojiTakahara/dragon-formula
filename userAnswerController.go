package dragonformula

import (
	"appengine"
	"appengine/datastore"
	"time"
	"github.com/martini-contrib/render"
	"net/http"
	"strconv"
)

/**
userKey
categoryKey
rightAnswer
wrongAnswer
question1
category1
corrected1
...
**/
func RegistUserAnswer(r render.Render, req *http.Request) {
	c := appengine.NewContext(req)
	unix := time.Now().Unix()
	timestamp := strconv.FormatInt(unix, 10)
	userAnswer := &UserAnswer{}
	userKey := req.FormValue("userKey")
	keyStr := userKey + "_" + timestamp
	userAnswer.UserKey = userKey
	userAnswer.TimeStamp = timestamp
	userAnswer.CategoryKey = req.FormValue("categoryKey")
	userAnswer.RightAnswer = ToInt(req.FormValue("rightAnswer"))
	userAnswer.WrongAnswer = ToInt(req.FormValue("wrongAnswer"))
	key := datastore.NewKey(c, "UserAnswer", keyStr, 0, nil)
	_, err := datastore.Put(c, key, userAnswer)
	if err != nil {
		c.Criticalf("%s", err)
		r.JSON(400, err)
	} else {
		for i := 1; i <= 15; i++ {
			RegistUserAnswerDetail(req, key.StringID(), i)	
		} 
		r.JSON(200, userAnswer)
	}
}

func RegistUserAnswerDetail(req *http.Request, k string, i int) {
	c := appengine.NewContext(req)
	detail := &UserAnswerDetail{}
	index := strconv.Itoa(i)
	key := datastore.NewKey(c, "UserAnswerDetail", k + "_" + index, 0, nil)
	detail.UerAnswerKey = k
	detail.UserKey = req.FormValue("userKey")
	detail.QuestionKey = req.FormValue("question" + index)
	detail.CategoryKey = req.FormValue("category" + index)
	detail.Corrected, _ = strconv.ParseBool(req.FormValue("corrected" + index))
	_, err := datastore.Put(c, key, detail)
	if err != nil {
		c.Criticalf("%s", err)
	}
}
