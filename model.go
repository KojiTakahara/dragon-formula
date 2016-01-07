package dragonformula

import (
	"time"
)

type Question struct { // key = 自動生成
	Content           string `datastore:",noindex"`
	LargeCategoryKey  string
	MediumCategoryKey string
	SmallCategoryKey  string
	Rubric            string `datastore:",noindex"`
	Percentage        float32
	Status            string // 依頼(REVIEW),最終確認,承認(APPROVED),却下
	Level             string
	UserKey           string // 作成者
	// ignored entirely by the datastore.
	Key               int64 `datastore:"-"`
	Choice1           QuestionChoice `datastore:"-"`
	Choice2           QuestionChoice `datastore:"-"`
	Choice3           QuestionChoice `datastore:"-"`
}

type QuestionChoice struct { // key = 自動採番
	Content       string `datastore:",noindex"`
	QuestionKeyId int64
	TrueFalse     bool
	// ignored entirely by the datastore.
	Key           int64 `datastore:"-"`
}

type Comment struct { // key = 自動採番
	QuestionKey string
	Status      string // 閲覧権限
	UserKey     string // 発言者
	CreatedAt   time.Time
}

type Category struct { // key = カテゴリ英名
	Key       string
	Name      string
	Type      int
	ParentKey string
	Number    float64
}

type User struct { // key = ユーザID
	Key       string
	Id        string
	Token     string
	Authority int
}

type UserAnswer struct { // key = ユーザID_タイムスタンプ
	UserKey     string
	TimeStamp   string
	CategoryKey string
	RightAnswer int // 正答数
	WrongAnswer int // 誤答数
}

type UserAnswerDetail struct { // key = ユーザID_タイムスタンプ_問番号
	UerAnswerKey string
	UserKey      string
	QuestionKey  string
	CategoryKey  string
	Corrected    bool
}
