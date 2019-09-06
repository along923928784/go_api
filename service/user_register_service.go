package service

import (
	"github.com/gin-gonic/gin"
	"jiyue.im/model"
	"jiyue.im/pkg/errno"
)

// @Summary 邮箱方式注册用户
// @Description Add a new user
// @Tags USER
// @Accept  json
// @Produce  json
// @Param service body model.RegisterRequest true "邮箱方式注册用户"
// @Success 200 {object} model.CreateResponse
// @Router /v1/user/register [post]
func UserRegister(c *gin.Context) {
	var r model.RegisterRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, err)
		return
	}

	u := model.UserModel{RegisterRequest: model.RegisterRequest{
		NickName: r.NickName,
		Email:    r.Email,
		Password: r.Password,
	},
	}

	// Validate the data.
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	// // Encrypt the user password.
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}
	// // Insert the user to the database.
	if err := u.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := model.CreateResponse{
		ID:        u.Id,
		NickeName: r.NickName,
	}

	SendResponse(c, nil, rsp)
}
