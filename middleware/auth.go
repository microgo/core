package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"master/constant"
	"strconv"
)

type Auth struct {
	SecureCookie *securecookie.SecureCookie
}

func (a *Auth) GetCurrentUserId(c *gin.Context) (int, error) {
	cookie, err := c.Request.Cookie("auth")
	if err != nil {
		return -1, err
	}
	value := ""
	err = a.SecureCookie.Decode("auth", cookie.Value, &value)
	if err != nil {
		return -1, err
	}
	userId, err := strconv.Atoi(value)
	if err != nil {
		return -1, err
	}
	return userId, err
}

func AuthInit() gin.HandlerFunc {
	hashKey := []byte(constant.AppHashKey)
	blockKey := []byte(constant.AppBlockKey)
	securecookie := securecookie.New(hashKey, blockKey)
	securecookie.MaxAge(3600 * 24 * 30)
	auth := Auth{
		SecureCookie: securecookie,
	}
	return func(c *gin.Context) {
		userId, err := auth.GetCurrentUserId(c)
		if err != nil {
			c.AbortWithStatus(401)
			return
		}
		if userId < 1 {
			c.AbortWithStatus(401)
			return
		}
		c.Set("userID", userId)
		c.Next()
	}
}
