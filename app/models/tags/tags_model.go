package tags

import "gin-blog/helpers/pool/grom"

var (
	Table = "tags"
	Model = grom.GetConn().Table(Table)
)