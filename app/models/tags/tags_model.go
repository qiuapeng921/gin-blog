package tags

import "gin-blog/helpers/pool/grom"

var (
	Table = "tags"
	Model = grom.GetOrm().Table(Table)
)