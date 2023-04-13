package service

import (
	"go-web/core/errors"
	"go-web/core/services"
	"go-web/core/tokenAuth"
	"go-web/model/vo"
	"go-web/repository"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	services.Service
}

// TenantLogin
// param from 表单认证信息
// result token 认证加密字符 err 返回错误信息
func (l LoginService) TenantLogin(from vo.LoginFrom) (token string, err error) {

	tenantRepository := repository.EcityosConfigTenantRepository{}
	l.MakeRepository(&tenantRepository.Repository)
	//获取当前用户
	data, err := tenantRepository.QueryTenantByUserName(from.Username)
	if err != nil {
		return "", &errors.CommonError{Msg: "用户名或密码错误！"}
	}
	//验证信息
	err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(from.Password))
	if err != nil {
		return "", &errors.CommonError{Msg: "用户名或密码错误！"}
	}
	j := tokenAuth.Token{UserID: data.UserId, LoginType: tokenAuth.LoginTypeTenement, ProjectID: ""}
	//生成token
	genToken, err := tokenAuth.GenToken(&j)
	if err != nil {
		return "", &errors.CommonError{Msg: "用户信息生成错误！"}
	}
	return genToken, err
}
