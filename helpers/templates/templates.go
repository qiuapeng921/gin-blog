package templates

import (
	"fmt"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"html/template"
	"path/filepath"
	"time"
)

func InitTemplate(router *gin.Engine) {
	router.Static("/static", "./public/assets")
	router.StaticFile("/favicon.ico", "./public/favicon.ico")
	//router.LoadHTMLGlob("templates/*/**")
	router.HTMLRender = loadTemplates("./templates")
	// 自定义方法
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatDate,
	})
}

//多模板（模板继承）
func loadTemplates(templatesDir string) multitemplate.Renderer {
	renderer := multitemplate.NewRenderer()

	articleLayouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
	if err != nil {
		panic(err.Error())
	}

	pages, pagesError := filepath.Glob(templatesDir + "/pages/*.html")
	if pagesError != nil {
		panic(pagesError.Error())
	}

	for _, page := range pages {
		layoutCopy := make([]string, len(articleLayouts))
		copy(layoutCopy, articleLayouts)
		files := append(layoutCopy, page)
		renderer.AddFromFiles(filepath.Base(page), files...)
	}
	return renderer
}

// 时间戳转时间格式
func formatDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}
