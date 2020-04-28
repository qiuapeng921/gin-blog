package request

type IndexRequest struct {
	Id    int `form:"id"`
	Page  int `form:"page"`
	Limit int `form:"limit"`
}