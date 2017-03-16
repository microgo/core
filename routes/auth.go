package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"master/constant"
	"master/model/response"
	"net/http"
	"strconv"
)

type LoginForm struct {
	Email    string
	Password string
}

const (
	DefaultCookieTime = 3600 * 24 * 365 * 2
)

func ApplyAuthRoutes(app *gin.Engine, router *Router) {
	group := app.Group("/feels-core/auth")
	group.POST("/login", router.Login)
}

func (r *Router) Login(c *gin.Context) {
	loginForm := LoginForm{}
	err := c.Bind(&loginForm)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}
	user, code := authEntity.UserAuthentication(loginForm.Email, loginForm.Password)
	if code != 200 {
		c.AbortWithStatus(code)
		return
	}
	_, err = r.SetAuthorizationToken("auth", strconv.Itoa(user.ID), "/", c.Writer)
	if err != nil {
		c.JSON(500, nil)
		return
	}
	res := response.Response{
		Code: code,
		Data: user,
	}
	c.JSON(code, res)
}

func (r *Router) Logout(c *gin.Context) {
	r.ClearCookie(c.Writer, "auth", "/")
	res := response.Response{
		Code:    200,
		Message: "Ok",
	}
	c.JSON(200, res)
}

func (r *Router) SetAuthorizationToken(cookieName, value, path string, w http.ResponseWriter) (string, error) {
	hashKey := []byte(constant.AppHashKey)
	blockKey := []byte(constant.AppBlockKey)
	secureCookie := securecookie.New(hashKey, blockKey)
	secureCookie.MaxAge(DefaultCookieTime)
	if encoded, err := secureCookie.Encode(cookieName, value); err == nil {
		cookie := &http.Cookie{
			Name:     cookieName,
			Value:    encoded,
			Path:     path,
			MaxAge:   DefaultCookieTime,
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
		return encoded, nil
	} else {
		return "", err
	}
}

func (r *Router) ClearCookie(w http.ResponseWriter, cookieName, cookiePath string) {
	ignoredContent := "Work Hard, Dream Big"
	cookie := fmt.Sprintf("%s=%s; path=%s; expires=Thu, 01 Jan 1970 00:00:00 GMT", cookieName, ignoredContent, cookiePath)
	r.SetHeader(w, "Set-Cookie", cookie, true)
}

func (r *Router) SetHeader(w http.ResponseWriter, hdr, val string, unique bool) {
	if unique {
		w.Header().Set(hdr, val)
	} else {
		w.Header().Add(hdr, val)
	}
}
