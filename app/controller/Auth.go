package controller

import (
	"gin-blog/helpers/response"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	if c.Request.Method == "post" {
		response.Context(c).Success(c.ClientIP())
		return
	}
	response.Context(c).View("login")
}

func Register(c *gin.Context) {
	if c.Request.Method == "post" {
		response.Context(c).Success(c.ClientIP())
		return
	}
	response.Context(c).View("register")
}
