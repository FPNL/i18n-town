package entity

type Pagination struct {
	Amount int `form:"amount,default=20"`
	Page   int `form:"page,default=1"`
	Total  int
	Order  string
}
