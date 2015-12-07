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
	Status            string // 依頼,最終確認,承認,却下
	UserKey           string // 作成者
	// ignored entirely by the datastore.
	Choice1 QuestionChoice `datastore:"-"`
	Choice2 QuestionChoice `datastore:"-"`
	Choice3 QuestionChoice `datastore:"-"`
	Choice4 QuestionChoice `datastore:"-"`
}

type QuestionChoice struct { // key = 自動採番
	Content       string `datastore:",noindex"`
	QuestionKeyId int64
	TrueFalse     bool
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
	RightAnswer int
	WrongAnswer int
}

type UserAnswerDetail struct { // key = 自動採番
	UerAnswerKey string
	UserKey      string
	QuestionKey  string
	CategoryKey  string
	Corrected    bool
	TimeStamp    string
}
