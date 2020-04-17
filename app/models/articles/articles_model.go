package articles

import "gin-blog/helpers/pool/grom"

var (
	Table = "articles"
	Model = grom.GetConn().Table(Table)
)
