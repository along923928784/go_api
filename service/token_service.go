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
		fmt.Println("not a vowel")
	}

}

func emailLogin(c *gin.Context, acction, secret string) {
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

func VerifyToken(c *gin.Context) {
	// ctx, b := c.Get("userContext")
	// if !b {
	// 	SendResponse(c, errno.ErrForbbiden, nil)
	// }
	// ID := ctx.(*token.Context).ID
	// scope := ctx.(*token.Context).Scope
	SendResponse(c, errno.OK, nil)
}
