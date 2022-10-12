package common

import (
	"Gin_GoBlog/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Success 成功响应的信息
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, model.ResMsg{
		Code: "1",
		Msg:  "success",
		Data: data,
	})
}

// Error 错误响应的信息
func Error(c *gin.Context, data interface{}) {
	c.JSON(http.StatusBadRequest, model.ResMsg{
		Code: "0",
		Msg:  "error",
		Data: data,
	})
}
