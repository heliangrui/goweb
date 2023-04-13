package jwtutil

//
//import (
//	"errors"
//	"github.com/gin-gonic/gin"
//	"go-web/core/response"
//	"gopkg.in/dgrijalva/jwt-go.v3"
//	"net/http"
//	"reflect"
//	"time"
//)
//
//var Secret = []byte("app_secret:SDFXFFVDSDFEDDSCNKJYJ")
//
//// jwt过期时间, 按照实际环境设置
//const expiration = 60 * time.Minute
//
//const expirationLong = 5 * time.Hour
//
//type Token struct {
//	// 自定义字段, 可以存在用户名, 用户ID, 用户角色等等
//	Username  string
//	UserId    string
//	Role      string
//	ProjectId string
//	// jwt.StandardClaims包含了官方定义的字段
//	jwt.StandardClaims
//}
//
//// GenToken 生成token
//// param token
//// result 认证token字符 错误信息
//func GenToken(token *Token) (string, error) {
//	// 创建声明
//	token.StandardClaims = jwt.StandardClaims{
//		ExpiresAt: time.Now().Add(expiration).Unix(), // 过期时间
//		IssuedAt:  time.Now().Unix(),                 // 签发时间
//		Issuer:    "gin-jwt-demo",                    // 签发者
//		Id:        "",                                // 按需求选这个, 有些实现中, 会控制这个ID是不是在黑/白名单来判断是否还有效
//		NotBefore: 0,                                 // 生效起始时间
//		Subject:   "",                                // 主题
//	}
//	// 用指定的哈希方法创建签名对象
//	tt := jwt.NewWithClaims(jwt.SigningMethodHS256, token)
//	// 用上面的声明和签名对象签名字符串token
//	// 1. 先对Header和PayLoad进行Base64URL转换
//	// 2. Header和PayLoadBase64URL转换后的字符串用.拼接在一起
//	// 3. 用secret对拼接在一起之后的字符串进行HASH加密
//	// 4. 连在一起返回
//	return tt.SignedString(Secret)
//}
//
//// CreateCookie 创建token
//func CreateCookie(ctx *gin.Context, token string) {
//	ctx.SetCookie("hypToken", token, 0, "", "", false, false)
//}
//
//// ParseToken 验证token是否有效
//func ParseToken(tokenStr string) (*Token, error) {
//	// 第三个参数: 提供一个回调函数用于提供要选择的秘钥, 回调函数里面的token参数,是已经解析但未验证的,可以根据token里面的值做一些逻辑, 如`kid`的判断
//	token, err := jwt.ParseWithClaims(tokenStr, &Token{}, func(token *jwt.Token) (interface{}, error) {
//		return Secret, nil
//	})
//	if err != nil {
//		return nil, err
//	}
//	// 校验token
//	if claims, ok := token.Claims.(*Token); ok && token.Valid {
//		return claims, nil
//	}
//	return nil, errors.New("invalid token")
//}
//
//// IsLogin 判断是否登录
//func IsLogin(ctx *gin.Context) bool {
//	cookies := ctx.Request.Cookies()
//	var cookie http.Cookie
//	for _, item := range cookies {
//		if item.Name == "hypToken" {
//			cookie = *item
//		}
//	}
//	if reflect.DeepEqual(cookie, http.Cookie{}) || cookie.Value == "" {
//		return false
//	}
//
//	value := cookie.Value
//	_, err := ParseToken(value)
//	if err != nil {
//		return false
//	}
//	return true
//}
//
//// JWTAuthMiddleware 认证解析
//func JWTAuthMiddleware() func(ctx *gin.Context) {
//	return func(ctx *gin.Context) {
//		// 根据实际情况取TOKEN, 这里从request header取
//		cookies := ctx.Request.Cookies()
//		var cookie http.Cookie
//		for _, item := range cookies {
//			if item.Name == "hypToken" {
//				cookie = *item
//			}
//		}
//		if reflect.DeepEqual(cookie, http.Cookie{}) || cookie.Value == "" {
//			response.Error(ctx, "用户未登录！")
//			return
//		}
//		value := cookie.Value
//		token, err := ParseToken(value)
//		if err != nil {
//			//var validationError *jwt.ValidationError
//			//is := errors.As(err, &validationError)
//			//if is && (time.Now().Unix()-token.ExpiresAt) < int64(expirationLong) {
//			//	genToken, err := GenToken(token)
//			//	if err != nil {
//			//		response.Error(ctx, "用户解析错误！")
//			//	}
//			//	CreateCookie(ctx, genToken)
//			//	appendToken(ctx, token)
//			//	return
//			//}
//			response.Error(ctx, "用户登录失效！")
//			return
//		}
//		appendToken(ctx, token)
//	}
//}
//
//func appendToken(ctx *gin.Context, token *Token) {
//	// 此处已经通过了, 可以把Claims中的有效信息拿出来放入上下文使用
//	ctx.Set("username", token.Username)
//	ctx.Set("userid", token.UserId)
//	ctx.Set("role", token.Role)
//	ctx.Set("projectId", token.ProjectId)
//	ctx.Next()
//}
//
//// GetToken 获取token
//func GetToken(ctx *gin.Context) *Token {
//	token := Token{}
//	token.UserId = ctx.GetString("username")
//	token.Role = ctx.GetString("role")
//	token.Username = ctx.GetString("username")
//	token.ProjectId = ctx.GetString("projectId")
//	return &token
//}
