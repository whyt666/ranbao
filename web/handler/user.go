package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"ranbao/pkg"
	"ranbao/web/middleware"
	"time"
)

// Login godoc
// @Summary 用户登录
// @Description 使用用户名和密码登录
// @Tags 用户
// @Accept json
// @Produce json
// @Param name formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {object} pkg.Result{data=map[string]string} "成功返回token"
// @Failure 400 {object} pkg.Result "参数错误"
// @Failure 401 {object} pkg.Result "登录失败"
// @Router /user/login [post]
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
		Msg: "登录成功,这是你的token",
		Data: map[string]string{
			"token": token,
		},
	})
}

// Register godoc
// @Summary 用户注册
// @Description 使用用户名和密码注册
// @Tags 用户
// @Accept json
// @Produce json
// @Param name formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {object} pkg.Result "注册成功"
// @Failure 400 {object} pkg.Result "参数错误"
// @Failure 409 {object} pkg.Result "用户已存在"
// @Router /user/register [post]
func (u *Handler) Register(ctx *gin.Context) {
	type RegisterForm struct {
		Name     string `form:"name" binding:"required"`
		Password string `form:"password" binding:"required"`
	}
	var form RegisterForm
	if err := ctx.Bind(&form); err != nil {
		return
	}

	err := u.UserService.RegisterByPassword(form.Name, form.Password)
	if err != nil {
		ctx.JSON(http.StatusOK, pkg.Result{
			Msg:  "注册失败",
			Data: err,
		})
		return
	}
	ctx.JSON(http.StatusOK, pkg.Result{
		Msg:  "注册成功",
		Data: nil,
	})
}

// LoginEmail godoc
// @Summary 邮箱登录/注册
// @Description 使用邮箱和验证码登录或注册
// @Tags 用户
// @Accept json
// @Produce json
// @Param email formData string true "邮箱"
// @Param code formData string true "验证码"
// @Success 200 {object} pkg.Result "注册成功"
// @Failure 400 {object} pkg.Result "参数错误"
// @Failure 401 {object} pkg.Result "验证码错误"
// @Failure 409 {object} pkg.Result "用户已存在"
// @Router /user/loginEmail [post]
func (u *Handler) LoginEmail(ctx *gin.Context) {
	type RegisterForm struct {
		Email string `form:"email" binding:"required"`
		Code  string `form:"code" binding:"required"`
	}
	var form RegisterForm
	if err := ctx.Bind(&form); err != nil {
		return
	}

	err := u.UserService.RegisterByEmail(form.Email, form.Code)
	if err != nil {
		ctx.JSON(http.StatusOK, pkg.Result{
			Msg:  "注册失败",
			Data: err,
		})
		return
	}
	ctx.JSON(http.StatusOK, pkg.Result{
		Msg:  "注册成功",
		Data: nil,
	})
}
