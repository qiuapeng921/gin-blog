package templates

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"time"
)

func InitTemplate(router *gin.Engine) {
	router.Static("/static", "./public/")
	router.StaticFile("/favicon.ico", "./public/favicon.ico")
	router.LoadHTMLGlob("templates/*")
	// 自定义方法
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatDate,
	})
}

func formatDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}
