package templates

import (
	"gin-blog/app/models/categorys"
	"gin-blog/helpers/pool/grom"
	"gin-blog/helpers/system"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"html/template"
	"path/filepath"
)

func InitTemplate(router *gin.Engine) {
	router.Static("/static", "./public/assets")
	router.StaticFile("/favicon.ico", "./public/favicon.ico")
	router.HTMLRender = loadTemplates("./templates")
}

//自定义函数
func FuncMap() template.FuncMap {
	return template.FuncMap{
		"formatDate":     formatDate,
		"markdownToHtml": markdownToHtml,
		"getCategory":    getCategory,
	}
}

//多模板（模板继承）
func loadTemplates(templatesDir string) multitemplate.Renderer {
	renderer := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
	if err != nil {
		panic(err.Error())
	}

	pages, pagesError := filepath.Glob(templatesDir + "/pages/*.html")
	if pagesError != nil {
		panic(pagesError.Error())
	}

	for _, page := range pages {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, page)
		renderer.AddFromFilesFuncs(filepath.Base(page), FuncMap(), files...)
	}
	return renderer
}

// 时间戳转时间格式
func formatDate(timestamp int64) string {
	return system.FormatDate(timestamp)
}

func markdownToHtml(markdown string) string {
	return system.MarkDownToHTML(markdown)
}

func getCategory() (category []categorys.Entity) {
	grom.GetConn().Find(&category)
	return category
}