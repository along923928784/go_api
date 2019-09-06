package model

import (
	validator "gopkg.in/go-playground/validator.v9"
	"jiyue.im/pkg/auth"
)

// User represents a registered user.
type UserModel struct {
	BaseModel
	RegisterRequest
	OpenId string `json:"openid" gorm:"column:openid"`
}

type RegisterRequest struct {
	NickName string `json:"nickname" gorm:"column:nickname;unique_index:hash_idx" binding:"required" validate:"min=1,max=32"`
	Email    string `json:"email" gorm:"column:email" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password" binding:"required" validate:"min=5,max=128"`
}

type CreateResponse struct {
	ID        uint64 `json:"id"`
	NickeName string `json:"nickname"`
}

func (u *UserModel) TableName() string {
	return "tb_users"
}

func (u *UserModel) Create() error {
	return DB.Create(&u).Error
}

func DeleteUser(id uint64) error {
	user := UserModel{}
	user.BaseModel.Id = id
	return DB.Delete(&user).Error
}

// Update updates an user account information.
func (u *UserModel) Update() error {
	return DB.Save(u).Error
}

// GetUser gets an user by the user identifier.
func GetUser(email string) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Where("email = ?", email).First(&u)
	return u, d.Error
}

func GetUserByOpenId(openid string) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Where("openid = ?", openid).First(&u)
	return u, d.Error
}

func (u *UserModel) RegisterByOpenId() error {
	return DB.Create(&u).Error
}

func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

// Validate the fields.
func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
