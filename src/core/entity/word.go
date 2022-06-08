package entity

type Word struct {
	Id   int    `form:"id"`
	Tag  string `form:"tag" binding:"required"`
	Lang string `form:"lang" binding:"required"`
	Word string `form:"word" binding:"required"`
}
