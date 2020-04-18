package comments

type Entity struct {
	Id         int    `orm:"id,primary"  json:"id"`          // ID
	UserId     int    `orm:"user_id"     json:"user_id"`     // 发表用户ID
	ArticleId  int    `orm:"article_id"  json:"article_id"`  // 评论博文ID
	Content    string `orm:"content"     json:"content"`     // 评论内容
	ParentId   int64  `orm:"parent_id"   json:"parent_id"`   // 父评论ID
	CreateTime int    `orm:"create_time" json:"create_time"` // 创建时间
}
