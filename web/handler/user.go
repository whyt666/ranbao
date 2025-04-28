package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"ranbao/pkg"
	"ranbao/web/middleware"
	"time"
)

func (u *Handler) Login(ctx *gin.Context) {
	type LoginForm struct {
		Name     string `form:"name" binding:"required"`
		Password string `form:"password" binding:"required"`
	}
	var form LoginForm
	if err := ctx.Bind(&form); err != nil {
		return
	}

	user, err := u.UserService.LoginByPassword(form.Name, form.Password)
	if err != nil {
		ctx.JSON(http.StatusOK, pkg.Result{
			Msg:  "登录失败",
			Data: err,
		})
		return
	}

	//发放token
	token, err := middleware.NewJWT().CreateToken(middleware.CustomClaims{
		ID:       user.ID,
		NickName: user.Name,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + 60*60*24*30, //30天过期
			Issuer:    "why",
		},
	})
	if err != nil {
		ctx.JSON(http.StatusOK, pkg.Result{
			Msg:  "发放token失败",
			Data: err,
		})
		return
	}
	ctx.JSON(http.StatusOK, pkg.Result{
		Msg: "登录成功，这是你的token",
		Data: map[string]string{
			"token": token,
		},
	})
}
