package dragonformula

type FormQuestion struct {
	Content     string `form:"content"`
	Answer1     string `form:"answer1"`
	Answer2     string `form:"answer2"`
	Answer3     string `form:"answer3"`
}