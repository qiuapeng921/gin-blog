package request

type ArticleRequest struct {
	Title   string `form:"title" binding:"required"`
	Content string `form:"content" binding:"required"`
}