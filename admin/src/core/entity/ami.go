package entity

type AMI struct {
	Id       string `form:"id"`
	Nickname string `form:"nickname" binding:"required"`
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Organize string `form:"organize" binding:"required"`
}

type SimpleAMI struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
