
package categorys

import "gin-blog/helpers/pool/grom"

var (
	Table = "categorys"
	Model = grom.GetConn().Table(Table)
)
