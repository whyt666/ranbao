package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"ranbao/pkg"
	"time"
)

func (h *Handler) SendCode(ctx *gin.Context) {
	type SendCode struct {
		Mail string `json:"mail"`
	}
	var form SendCode
	if err := ctx.Bind(&form); err != nil {
		return
	}

	context1, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := h.MailService.SendMail(context1, form.Mail)
	if err != nil {
		ctx.JSON(http.StatusOK, pkg.Result{
			Msg:  "发送验证码失败",
			Data: err,
		})
		return
	}
	ctx.JSON(http.StatusOK, pkg.Result{
		Msg:  "发送成功",
		Data: nil,
	})
}

func (h *Handler) Verify(ctx *gin.Context) {
	type SendCode struct {
		Mail string `json:"mail"`
		Code string `json:"code"`
	}
	var form SendCode
	if err := ctx.Bind(&form); err != nil {
		return
	}

	context1, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := h.MailService.Verify(context1, form.Mail, form.Code)
	if err != nil {
		ctx.JSON(http.StatusOK, pkg.Result{
			Msg:  "",
			Data: err,
		})
		return
	}
	ctx.JSON(http.StatusOK, pkg.Result{
		Msg:  "验证码正确",
		Data: nil,
	})
}
