package controller

import (
	"github.com/gin-gonic/gin"
	"go-web/core/controllers"
	"go-web/core/tokenAuth"
	"go-web/core/validator"
	"go-web/model/vo"
	"go-web/service"
	"net/http"
)

// LoginController 登录模块
type LoginController struct {
	controllers.Controller
}

// Login ide登录
func (l LoginController) Login(c *gin.Context) {
	from := vo.LoginFrom{}
	err := l.MakeContext(c).BindQueryParam(&from).Error
	if err != nil {
		translate := validator.Translate(err, &from)
		l.Err(translate)
		return
	}
	//判断是否登录
	login, _ := tokenAuth.IsLogin(c)
	if login {
		l.Err("用户已登录！请退出重新登录！")
		return
	}
	loginService := service.LoginService{}
	l.MakeService(&loginService.Service)

	token, err := loginService.TenantLogin(from)
	if err != nil {
		l.Err(err.Error())
		return
	}
	tokenAuth.CreateCookie(l.Context, token)
	l.Logger.Info(token)
	l.OK("登录成功！")
}

func (l LoginController) Logout(ctx *gin.Context) {
	l.MakeContext(ctx)
	cookies := ctx.Request.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == "hypToken" {
			cookie.Value = ""
			http.SetCookie(ctx.Writer, &http.Cookie{
				Name:     "hypToken",
				Value:    "",
				MaxAge:   0,
				Path:     "/",
				Domain:   "",
				Secure:   false,
				HttpOnly: true,
			})
		}
	}

	l.OK("退出登录成功")

}
