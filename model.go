package dragonformula

import (
	"appengine/datastore"
	"time"
)

type Question struct { // key = 自動生成
	Content           string `datastore:",noindex"`
	LargeCategoryKey  datastore.Key
	MediumCategoryKey datastore.Key
	SmallCategoryKey  datastore.Key
	Answer1Key        datastore.Key
	Answer2Key        datastore.Key
	Answer3Key        datastore.Key
	Answer4Key        datastore.Key
	Rubric            string `datastore:",noindex"`
	Percentage        float32
	Status            string        // 依頼,最終確認,承認,却下
	UserKey           datastore.Key // 作成者
}

type QuestionAnswer struct { // key = 自動採番
	Content   string `datastore:",noindex"`
	TrueFalse bool
}

type Comment struct { // key = 自動採番
	QuestionKey datastore.Key
	Status      string        // 閲覧権限
	UserKey     datastore.Key // 発言者
	CreatedAt   time.Time
}

type Category struct { // key = カテゴリ英名
	KeyName string
	Name string
	Type int
	ParentKeyName string
}

type User struct { // key = ユーザID
	Authority int
}

type UserAnswer struct { // key = ユーザID_タイムスタンプ
	UserKey     datastore.Key
	TimeStamp   string
	CategoryKey datastore.Key
	RightAnswer int
	WrongAnswer int
}

type UserAnswerDetail struct { // key = 自動採番
	UerAnswerKey datastore.Key
	UserKey      datastore.Key
	QuestionKey  datastore.Key
	CategoryKey  datastore.Key
	Corrected    bool
	TimeStamp    string
}
