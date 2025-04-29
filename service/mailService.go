package service

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gopkg.in/gomail.v2"
	"math/rand"
	"ranbao/service/cache"
	"time"
)

var (
	ErrCodeVaild = errors.New("验证码无效")
)

type IMailServiceV1 interface {
	SendMail(ctx context.Context, mail string) error
	Verify(ctx context.Context, mail string, code string) error
}

type MailService struct {
}

const (
	QQMailSMTPCode = "muxammcfejavebei"
	QQMailSender   = "@qq.com"
	QQMailTitle    = "验证码"
	SMTPAdr        = "smtp.qq.com"
	SMTPPort       = 25
	MailListSize   = 2048

	SecretId    = "AKIDSjlgKhpgFXYIkvpRmqc8MK1ScU5PSGo4"
	SecreKey    = "1lDOJxKQFn44nFPItDH4ZnkGWUCdmeFZ"
	SmsSdkAppId = "1400696970"
	TemplateId  = "1452825"
)

// SendMail QQ邮箱验证码
func (s *MailService) SendMail(ctx context.Context, mail string) error {

	//产生六位数验证码
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))

	//将验证码记录到redis中
	err := cache.CacheMailCtl.Set(ctx, mail, vcode)
	if err != nil {
		return err
	}

	//发送的内容
	html := fmt.Sprintf(`<div>
        <div>
            尊敬的用户，您好！
        </div>
        <div style="padding: 8px 40px 8px 50px;">
            <p>您的验证码为: %s </p>
			<p>验证码有效期为3分钟左右，请尽快使用!!!</p>
        </div>
        <div>
            <p>此邮箱为系统邮箱，请勿回复。</p>
        </div>
    </div>`, vcode)

	m := gomail.NewMessage()

	m.SetHeader(`From`, "whyt858@qq.com")
	m.SetHeader(`To`, mail)
	m.SetHeader(`Subject`, QQMailTitle)
	m.SetBody(`text/html`, html)

	err = gomail.NewDialer(SMTPAdr, 25, "whyt858@qq.com", "muxammcfejavebei").DialAndSend(m)

	if err != nil {
		zap.S().Panicf("Send Email Fail, %s", err.Error())
		return err
	}

	zap.S().Info("Send Email Success")
	return nil
}

func (s *MailService) Verify(ctx context.Context, mail string, code string) error {
	//先查询
	res, err := cache.CacheMailCtl.Get(ctx, mail)
	if err != nil {
		zap.S().Info(err)
		return err
	}
	//后对比
	if res != code {
		return ErrCodeVaild
	}

	return nil
}
