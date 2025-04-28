package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ranbao/pkg"
	"ranbao/service"
	"ranbao/web/middleware"
)

var _ service.IMailServiceV1 = (*service.MailService)(nil)

type Handler struct {
	UserService service.IUserService
	MailService service.IMailServiceV1
}

func NewHandler() *Handler {
	return &Handler{
		UserService: &service.UserService{},
		MailService: &service.MailService{},
	}
}

func InitRouter() *gin.Engine {
	r := gin.Default()
	//发送邮箱验证码

	UserRouter := r.Group("/user")
	handler := NewHandler()

	{
		UserRouter.POST("/login", handler.Login)
		UserRouter.POST("/register", handler.Register)
		UserRouter.Use(middleware.JWTAuth()).POST("/ForgetPwd", func(context *gin.Context) {

		})
	}
	mailRouter := r.Group("/mail")
	{
		mailRouter.POST("/sendCode", handler.SendCode)
		mailRouter.GET("/Verify", handler.Verify)
	}
	return r
}

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
