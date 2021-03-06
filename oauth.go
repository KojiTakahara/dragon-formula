package dragonformula

import (
	"appengine"
	"appengine/datastore"
	"appengine/urlfetch"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/clbanning/mxj"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"github.com/mrjones/oauth"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

// var TWITTER_CONSUMER_KEY = "efelQ7UNWZGge5H1p7OMRl0lR"
// var TWITTER_CALLBACK_URL = "http://stg-aqua-teacher.appspot.com/api/twitter/callback"
// var TWITTER_CONSUMER_SECRET = "51UF3rlpJWLDwVhP8lsr0fawd0y7xksmatAyln6ANtgZz3kv8s"
var TWITTER_CONSUMER_KEY = "tFc1luf519X71ihzPS6Mohibn"
var TWITTER_CALLBACK_URL = "http://aqua-teacher.appspot.com/api/twitter/callback"
var TWITTER_CONSUMER_SECRET = "KD84L4SqwdrxOiw6T2DTw4TuzcefWqdZCRCGXhxWCmj3aDSPFy"
var TWITTER_SERVER = oauth.ServiceProvider{
	RequestTokenUrl:   "http://api.twitter.com/oauth/request_token",
	AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
	AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
}
var consumer = oauth.NewConsumer(TWITTER_CONSUMER_KEY, TWITTER_CONSUMER_SECRET, TWITTER_SERVER)

// Twitter ボタンが押された時の処理
func LoginTwitter(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	oauth := NewOAuth1(c, fmt.Sprintf("http://aqua-teacher.appspot.com/"))
	result := oauth.RequestToken("https://api.twitter.com/oauth/request_token")
	oauth.Authenticate(w, r, "https://api.twitter.com/oauth/authenticate", result["oauth_token"])
}

func LoginUser(r render.Render, req *http.Request, session sessions.Session) {
	accessToken := GetAccessToken(session)
	if accessToken != nil {
		result := GetTwitterUser(req, accessToken)
		if result == nil {
			r.JSON(400, "")
		}
		r.JSON(200, result)
		return
	} else {
		r.JSON(400, "")
	}
}

func GetTwitterUser(req *http.Request, accessToken *oauth.AccessToken) mxj.Map {
	c := appengine.NewContext(req)
	consumer.HttpClient = urlfetch.Client(c)
	response, _ := consumer.Get("https://api.twitter.com/1.1/account/verify_credentials.json", nil, accessToken)
	result := make([]byte, 1024*1024)
	response.Body.Read(result)
	resultString := string(result)
	trimStr := strings.Trim(resultString, "\x00")
	if trimStr == "" {
		return nil
	}
	accountInfo, _ := mxj.NewMapJson([]byte(trimStr))
	return accountInfo
}

/**
 * Twitterからのコールバックを受け取って、セッションに詰める。ユーザの登録・更新を行う。
 **/
func CallbackTwitter(r render.Render, w http.ResponseWriter, req *http.Request, session sessions.Session) {
	c := appengine.NewContext(req)
	token := req.FormValue("oauth_token")
	verifier := req.FormValue("oauth_verifier")
	oauth := NewOAuth1(c, TWITTER_CALLBACK_URL)
	accessToken := oauth.ExchangeToken(token, verifier, "https://api.twitter.com/oauth/access_token")
	if accessToken != nil { // ログイン成功
		// セッション開始
		session.Set("accessToken", accessToken.Token)
		session.Set("accessTokenSecret", accessToken.Secret)

		c.Infof("toke")
		c.Infof(accessToken.Token)
		c.Infof("secret")
		c.Infof(accessToken.Secret)

		consumer.HttpClient = urlfetch.Client(c)
		response, _ := consumer.Get("https://api.twitter.com/1.1/account/verify_credentials.json", nil, accessToken)
		result := make([]byte, 1024*1024)
		response.Body.Read(result)
		resultString := string(result)
		trimStr := strings.Trim(resultString, "\x00")

		accountInfo, _ := mxj.NewMapJson([]byte(trimStr))

		// User登録
		if accountInfo["screen_name"] != nil {
			screenName := accountInfo["screen_name"].(string)
			user := &User{}
			user.Id = screenName
			user.Key = screenName
			user.Token = accessToken.Token
			key := datastore.NewKey(c, "User", screenName, 0, nil)
			key, err := datastore.Put(c, key, user)
			if err != nil {
				c.Criticalf("%s", err)
			} else {
				c.Infof("")
			}
		}
		// TODO トップページへリダイレクト
		r.Redirect("/")
	} else { //　ログイン失敗
		r.Redirect("/")
	}
}

/**
 * OAuth1.0aの通信を行うクラス
 * @class
 * @param {map[string]string} params oauthパラメータの配列
 * @param {appengine.Context} context コンテキスト
 */
type OAuth1 struct {
	params  map[string]string
	context appengine.Context
}

/**
 * OAuthクラスのインスタンス化
 * @function
 * @params {appengine.Context} c コンテキスト
 * @params{appengine.Context} callback コールバックURL
 * @returns {*OAuth} OAuthインスタンス
 */
func NewOAuth1(c appengine.Context, callback string) *OAuth1 {
	oauth := new(OAuth1)
	oauth.context = c
	return oauth
}

/**
 * Twitter へリクエストトークンを要求する
 * @method
 * @memberof OAuth1
 * @param {string} targetUrl リクエスト要求先のURL
 * @returns {map[string]string} リクエスト結果
 */
func (this *OAuth1) RequestToken(targetUrl string) map[string]string {
	params := make(map[string]string, 0)
	params["oauth_callback"] = TWITTER_CALLBACK_URL
	response := this.Request("POST", targetUrl, make(map[string]string), "", []string{TWITTER_CONSUMER_SECRET, ""})
	datas := strings.Split(response, "&")
	result := make(map[string]string, len(datas))
	for i := 0; i < len(datas); i++ {
		data := strings.Split(datas[i], "=")
		result[data[0]] = data[1]
	}

	return result
}

/**
 * リクエストを送信してレスポンスを受信する
 * メソッドは POST 固定
 * @method
 * @memberof OAuth1
 * @param {string} method POSTかGET
 * @param {string} targetUrl 送信先
 * @param {string} params パラメータ
 * @param {string} body リクエストボディ
 * @param {string} secret 暗号鍵（ConsumerSecret と OAuth Token Secret）
 * @returns {string} レスポンス
 */
func (this *OAuth1) Request(method string, targetUrl string, params map[string]string, body string, secret []string) string {
	oauthParams := make(map[string]string, 0)

	for key, val := range params {
		oauthParams[key] = val
	}

	oauthParams["oauth_consumer_key"] = TWITTER_CONSUMER_KEY
	oauthParams["oauth_signature_method"] = "HMAC-SHA1"
	oauthParams["oauth_version"] = "1.0"
	oauthParams["oauth_nonce"] = this.CreateNonce()
	oauthParams["oauth_timestamp"] = strconv.Itoa(int(time.Now().Unix()))
	oauthParams["oauth_signature"] = this.CreateSignature(method, targetUrl, oauthParams, secret)

	for key, val := range params {
		oauthParams[key] = val
	}

	// リクエスト送信
	httpParams := make(map[string]string, 1)
	httpParams["Authorization"] = this.CreateHeader(oauthParams)

	var response *http.Response
	if method == "GET" {
		query := make(map[string]string, 0)
		for key, val := range params {
			if key != "oauth_token" {
				query[key] = val
			}
		}
		response = Get(this.context, targetUrl, query, httpParams)
	} else {
		response = Request(this.context, method, targetUrl, httpParams, body)
	}

	// レスポンスボディの読み取り
	result := make([]byte, 1024*1024)
	response.Body.Read(result)
	resultString := string(result)
	resultString = strings.Trim(resultString, "\x00")

	return resultString
}

func GetTwitterUserById(r render.Render, params martini.Params, req *http.Request) {
	c := appengine.NewContext(req)
	oauth := NewOAuth1(c, TWITTER_CALLBACK_URL)
	url := "https://api.twitter.com/1.1/users/show.json"
	parameters := map[string]string{"screen_name": params["userId"]}
	secret := []string{TWITTER_CONSUMER_SECRET, ""}
	response := oauth.Request("GET", url, parameters, "", secret)
	user, _ := mxj.NewMapJson([]byte(response))
	r.JSON(200, user)
}

/**
 * oauth_nonce を作成する
 * @method
 * @memberof OAuth1
 * @returns {string} 作成したoauth_nonce
 */
func (this *OAuth1) CreateNonce() string {
	nonce := ""
	for i := 0; i < 4; i++ {
		nonce = strings.Join([]string{nonce, string(GetRandomizedString())}, "")
	}
	return nonce
}

/**
 * Aouthorization ヘッダを作成する
 * @method
 * @memberof OAuth1
 * @param {map[string]string} oauthParams OAuthパラメータ
 * @returns {string} ヘッダ
 */
func (this *OAuth1) CreateHeader(oauthParams map[string]string) string {
	headerParams := make([]string, 0)
	for key, val := range oauthParams {
		if key == "screen_name" {
			continue
		}
		key = url.QueryEscape(key)
		val = url.QueryEscape(val)
		set := fmt.Sprintf(`%s="%s"`, key, val)
		headerParams = append(headerParams, set)
	}
	header := strings.Join(headerParams, ", ")
	header = fmt.Sprintf("OAuth %s", header)
	return header
}

/**
 * oauth_signature を作成する
 * @method
 * @memberof OAuth1
 * @param {string} method メソッド
 * @param {string} targetUrl リクエスト送信先のURL
 * @param {map[string]string} oauthParams パラメータ
 * @param {string} secret 暗号化鍵(Consumer Secret と OAuth Token Secret)
 * @returns {string} oauth_signature
 */
func (this *OAuth1) CreateSignature(method string, targetUrl string, oauthParams map[string]string, secret []string) string {
	keys := make([]string, 0)
	for key := range oauthParams {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	params := make([]string, len(keys))
	for i := 0; i < len(keys); i++ {
		key := keys[i]
		val := oauthParams[key]
		params[i] = fmt.Sprintf("%s=%s", url.QueryEscape(key), url.QueryEscape(val))
	}
	paramString := strings.Join(params, "&")
	baseString := fmt.Sprintf("%s&%s&%s", method, url.QueryEscape(targetUrl), url.QueryEscape(paramString))

	signatureKey := fmt.Sprintf("%s&%s", url.QueryEscape(secret[0]), url.QueryEscape(secret[1]))

	hash := hmac.New(sha1.New, []byte(signatureKey))
	hash.Write([]byte(baseString))
	signature := hash.Sum(nil)
	return base64.StdEncoding.EncodeToString(signature)
}

/**
 * 認証ページヘリダイレクトする
 * @memberof OAuth1
 * @method
 * @param {http.ResponseWriter} w 応答先
 * @param {*http.Request} r リクエスト
 * @param {string} targetUrl リダイレクト先
 * @param {string} token 未認証リクエストトークン
 */
func (this *OAuth1) Authenticate(w http.ResponseWriter, r *http.Request, targetUrl string, token string) {
	to := fmt.Sprintf("?oauth_token=%s", token)
	to = strings.Join([]string{targetUrl, to}, "")
	http.Redirect(w, r, to, 302)
}

/**
 * リクエストトークンをアクセストークンに変換する
 * @memberof OAuth1
 * @method
 * @param {string} token リクエストトークン
 * @param {string} verifier 認証データ
 * @param {string} targetUrl リクエストの送信先
 * @returns {oauth.AccessToken} アクセストークン
 */
func (this *OAuth1) ExchangeToken(token string, verifier string, targetUrl string) *oauth.AccessToken {
	requestToken := &oauth.RequestToken{Token: token}
	consumer.HttpClient = urlfetch.Client(this.context)
	accessToken, err := consumer.AuthorizeToken(requestToken, verifier)
	if err != nil {
		return nil
	}
	return accessToken
}

func GetAccessToken(session sessions.Session) *oauth.AccessToken {
	token := session.Get("accessToken")
	secret := session.Get("accessTokenSecret")
	if token == nil || token == "" || secret == nil || secret == "" {
		return nil
	}
	return &oauth.AccessToken{
		Token:  token.(string),
		Secret: secret.(string),
	}
}
