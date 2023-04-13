package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var ResultCodeError = "1"
var ResultCodeOk = "0"
var ResultCodeTokenAuthError = "401"
var ResultCodeSystemAuthError = "402"

func Error(c *gin.Context, msg string) {
	res := Response{Result: result(struct{ ResultCode, ResultError string }{ResultCode: ResultCodeError, ResultError: msg})}
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func TokenAuthError(c *gin.Context, msg string) {
	res := Response{Result: result(struct{ ResultCode, ResultError string }{ResultCode: ResultCodeTokenAuthError, ResultError: msg})}
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func SystemAuthError(c *gin.Context, msg string) {
	res := Response{Result: result(struct{ ResultCode, ResultError string }{ResultCode: ResultCodeSystemAuthError, ResultError: msg})}
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func OK(c *gin.Context, data interface{}) {
	res := Response{Result: result(struct{ ResultCode, ResultError string }{ResultCode: ResultCodeOk}), Data: data}
	c.AbortWithStatusJSON(http.StatusOK, res)

}
