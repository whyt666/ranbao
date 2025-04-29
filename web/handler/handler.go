package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "ranbao/docs"
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

	r.Use(middleware.FormLogger())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	UserRouter := r.Group("/user")
	handler := NewHandler()

	{

		UserRouter.POST("/login", handler.Login)
		UserRouter.POST("/loginEmail", handler.LoginEmail)
		UserRouter.POST("/register", handler.Register)
		UserRouter.Use(middleware.JWTAuth()).POST("/ForgetPwd", func(context *gin.Context) {

		})
	}
	mailRouter := r.Group("/mail")
	{

		mailRouter.POST("/sendCode", handler.SendCode)
		mailRouter.POST("/verify", handler.Verify)
	}
	return r
}
