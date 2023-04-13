package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/valyala/fastjson"
	"go-web/core/response"
	"go-web/core/services"
	"go.uber.org/zap"
)

type Controller struct {
	Context *gin.Context
	Logger  *zap.Logger
	//Orm     *gorm.DB
	Error error
}

func (e *Controller) AddError(err error) error {
	if e.Error == nil {
		e.Error = err
	} else if err != nil {
		e.Logger.Error(e.Error.Error())
		e.Error = err
	}
	return e.Error
}

func (e *Controller) MakeService(c *services.Service) *Controller {
	c.Logger = e.Logger
	return e
}

// MakeContext 设置http上下文
func (e *Controller) MakeContext(c *gin.Context) *Controller {
	e.Context = c
	e.Logger = zap.L()
	//e.Logger = api.GetRequestLogger(c)
	return e
}

func (e *Controller) BindQueryParam(d interface{}) *Controller {
	var err error
	err = e.Context.ShouldBindQuery(d)
	e.AddError(err)
	return e
}
func (e *Controller) BindQueryPathParam(d interface{}) *Controller {
	var err error
	err = e.Context.ShouldBindUri(d)
	e.AddError(err)
	return e
}

func (e *Controller) Bind(d interface{}, bindings ...binding.Binding) *Controller {
	var err error
	if len(bindings) == 0 {
		bindings = append(bindings, binding.JSON, nil)
	}
	for i := range bindings {
		switch bindings[i] {
		case binding.JSON:
			err = e.Context.ShouldBindWith(d, binding.JSON)
		case binding.XML:
			err = e.Context.ShouldBindWith(d, binding.XML)
		case binding.Form:
			err = e.Context.ShouldBindWith(d, binding.Form)
		case binding.Query:
			err = e.Context.ShouldBindWith(d, binding.Query)
		case binding.FormPost:
			err = e.Context.ShouldBindWith(d, binding.FormPost)
		case binding.FormMultipart:
			err = e.Context.ShouldBindWith(d, binding.FormMultipart)
		case binding.ProtoBuf:
			err = e.Context.ShouldBindWith(d, binding.ProtoBuf)
		case binding.MsgPack:
			err = e.Context.ShouldBindWith(d, binding.MsgPack)
		case binding.YAML:
			err = e.Context.ShouldBindWith(d, binding.YAML)
		case binding.Header:
			err = e.Context.ShouldBindWith(d, binding.Header)
		default:
			err = e.Context.ShouldBindUri(d)
		}
		if err != nil {
			e.AddError(err)
		}
	}
	return e
}

func (e *Controller) Err(msg string) {
	response.Error(e.Context, msg)
}

func (e *Controller) OK(data interface{}) {
	response.OK(e.Context, data)
}

func (e *Controller) OKFastJsonParse(parse *fastjson.Value) {
	var v map[string]interface{}
	json.Unmarshal([]byte(parse.String()), &v)
	response.OK(e.Context, v)
}
