package comments

import (
	"gin-blog/helpers/pool/grom"
)

var (
	Table = "comments"
	Model = grom.GetConn().Table(Table)
)
