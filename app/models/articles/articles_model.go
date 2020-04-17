package articles

import "gin-blog/helpers/pool/grom"

var (
	Model = grom.GetConn().Table("articles")
)
