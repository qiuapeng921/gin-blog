package system

import (
	"gopkg.in/russross/blackfriday.v1"
)

func StringToMarkDown(input []byte) []byte {
	return blackfriday.MarkdownBasic(input)
}
