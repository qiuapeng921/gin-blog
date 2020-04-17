package articles

type Entity struct {
	Id           int    `orm:"id,primary"    json:"id"`            // 主键id
	UserId       int    `orm:"user_id"       json:"user_id"`       // 用户ID
	Title        string `orm:"title"         json:"title"`         // 文章标题
	CategoryId   int    `orm:"category_id"   json:"category_id"`   // 分类id
	Content      string `orm:"content"       json:"content"`       // 博文内容
	Views        int64  `orm:"views"         json:"views"`         // 浏览量
	LikeCount    int64  `orm:"like_count"    json:"like_count"`    // 喜欢数
	CommentCount int64  `orm:"comment_count" json:"comment_count"` // 评论总数
	Status       int    `orm:"status"        json:"status"`        // 0 正常 1 删除
	CreateTime   int    `orm:"create_time"   json:"create_time"`   // 创建时间
	UpdateTime   int    `orm:"update_time"   json:"update_time"`   // 修改时间
}