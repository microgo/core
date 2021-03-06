package entities

import (
	"master/model/postgres"
	"master/resource/helper"
)

type AuthInterface interface {
	UserAuthentication(string, string) (*postgres.User, int)
}

type Auth struct {
	H *helper.Helper
}

func NewAuthEntity(h *helper.Helper) AuthInterface {
	return &Auth{H: h}
}

func (e *Auth) UserAuthentication(email, password string) (*postgres.User, int) {
	return nil, 404
}
