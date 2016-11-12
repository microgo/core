package entities

import (
	"master/model/postgres"
	"master/resource/helper"
)

type UserInterface interface {
	GetUserByID(int) (*postgres.User, int)
}

type User struct {
	H *helper.Helper
}

func NewUserEntity(h *helper.Helper) UserInterface {
	return &User{H: h}
}

func (e *User) GetUserByID(photoID int) (*postgres.User, int) {
	return nil, 200
}
