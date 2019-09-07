package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"jiyue.im/model"
	"jiyue.im/pkg/auth"
	"jiyue.im/pkg/enum"
	"jiyue.im/pkg/errno"
	"jiyue.im/pkg/token"
)

// @Summary 获取token
// @Tags User
// @Accept  json
// @Produce  json
// @Param token body model.TokenRequest true "token"
// @Success 200 {object} model.Token "{"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1Njc3Njk4MDYsImlhdCI6MTU2Nzc2MjYwNiwiaWQiOjEsIm5iZiI6MTU2Nzc2MjYwNiwic2NvcGUiOjh9.-VlL6oAa8mMD2Wd0Os1in1V5T9sdcwv6OCupihZKZNY"}"
// @Router /v1/user/token  [post]
func GetToken(c *gin.Context) {
	var t model.TokenRequest

	if err := c.Bind(&t); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	if err := t.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	switch enum.LoginType(t.Type) {
	case enum.USER_EMAIL:
		emailLogin(c, t.Account, t.Secret)
	case enum.USER_MINI_PROGRAM:
		codeToToken(c, t.Account)
	case enum.USER_MOBILE:
		fmt.Println(t.Type)
	default:
		SendResponse(c, errno.ErrType, nil)
	}

}

// @Summary 验证Token合法性
// @Tags User
// @Accept  json
// @Produce  json
// @Param x-auth-token header string  true "x-auth-token"
// @Success 200 {object} service.Response "OK"
// @Router /v1/user/token/verify [post]
func VerifyToken(c *gin.Context) {
	// ctx, b := c.Get("userContext")
	// if !b {
	// 	SendResponse(c, errno.ErrForbbiden, nil)
	// }
	// ID := ctx.(*token.Context).ID
	// scope := ctx.(*token.Context).Scope
	SendResponse(c, errno.OK, "OK")
}

func emailLogin(c *gin.Context, acction, secret string) {
	if secret == "" {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	d, err := model.GetUser(acction)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	if err := auth.Compare(d.Password, secret); err != nil {
		SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}

	t, err := token.Sign(c, token.Context{ID: d.Id, Scope: auth.USER}, "")
	if err != nil {
		SendResponse(c, errno.ErrToken, nil)
		return
	}
	SendResponse(c, nil, model.Token{Token: t})

}
