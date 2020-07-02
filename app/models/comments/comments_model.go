package comments

import (
	"gin-blog/helpers/pool/grom"
)

var (
	Table = "comments"
	Model = grom.GetOrm().Table(Table)
)
