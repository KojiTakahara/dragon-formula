package dragonformula

import (
	"appengine"
	"appengine/datastore"
	"github.com/ajg/form"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"strconv"
	"strings"
)

func CreateDeckSheet(r render.Render, params martini.Params, w http.ResponseWriter, req *http.Request) {
	// id, _ := strconv.Atoi(params["id"])
	// var deck = GetDeckByVault(id, req)
	// if deck.MainDeck == nil {
	// 	r.JSON(400, "badRequest")
	// } else {
	// 	r.JSON(200, deck)
	// }
	r.JSON(200, "")
}

func RegistDeck(r render.Render, req *http.Request) {
	c := appengine.NewContext(req)


  qa := &QuestionAnswer{}
  qa.Content = ""
  qa.TrueFalse = true
  qaKey := datastore.NewKey(c, "", "" 0, nil)


	var formDeck FormDeck
	decoder := form.NewDecoder(req.Body)
	if err := decoder.Decode(&formDeck); err != nil {
		c.Infof("Could not decode slack message: ", err)
		r.JSON(400, err)
		return
	}
	if formDeck.DeckId == "" {
		r.JSON(200, "")
		return
	}
	deck := &Deck{}
	deck.Format = formDeck.Format
	deck.MainDeck = strings.Split(formDeck.MainDeck, ",")
	deck.HyperSpatial = strings.Split(formDeck.HyperSpatial, ",")
	id, _ := strconv.Atoi(formDeck.DeckId)
	key := datastore.NewKey(c, "Deck", "", int64(id), nil)
	key, err := datastore.Put(c, key, deck)
	if err != nil {
		c.Criticalf("%s", err)
		r.JSON(400, err)
	} else {
		r.JSON(200, deck)
	}
}