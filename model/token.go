package model

import validator "gopkg.in/go-playground/validator.v9"

type TokenRequest struct {
	Type    int32  `json:"type"`
	Account string `json:"account" binding:"required" validate:"min=1,max=32"`
	Secret  string `json:"secret"`
}

type Token struct {
	Token string `json:"token"`
}

func (t *TokenRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(t)
}
