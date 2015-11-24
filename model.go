package dragonformula

type Question struct { // key = 自動生成
	Content string `datastore:",noindex"`
	LargeCategory Key
	MediumCategory Key
	SmallCategory Key
	Answer1 string `datastore:",noindex"`
	Answer2 string `datastore:",noindex"`
	Answer3 string `datastore:",noindex"`
	Answer4 string `datastore:",noindex"`
	Rubric string `datastore:",noindex"`
	Percentage Num
	Status string // 依頼,最終確認,承認,却下
	userKey Key // 作成者
}

type Comment struct { // key = 自動採番
	QuestionKey Key
	Status string // 閲覧権限
	userKey Key // 発言者
	createdAt Date
}

type Category struct { // key = カテゴリ英名
	Name string
	Type Num
}

type User struct { // key = ユーザID
	Authority Num
}

type UserAnswer struct { // key = ユーザID_タイムスタンプ
	UserKey Key
	TimeStamp string
	CategoryKey Key
	RightAnswer Num
	WrongAnswer Num
	CategoryKey Key
}

type UserAnswerDetail struct { // key = 自動採番
	UerAnswerKey Key
	UserKey Key
	QuestionKey Key
	CategoryKey Key
	Corrected bool
	TimeStamp string
}