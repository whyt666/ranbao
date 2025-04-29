package service

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"ranbao/service/dao"
	"ranbao/service/model"
	"strings"
)

var (
	UserNotFind    = errors.New("用户未找到")
	UserRegistered = errors.New("用户已注册")
	PasswordErr    = errors.New("密码错误")
)

type IUserService interface {
	RegisterByPassword(name string, password string) error
	RegisterByEmail(email string, code string) error
	LoginByPassword(name string, password string) (model.User, error)
}

type UserService struct {
}

func (u *UserService) RegisterByEmail(email string, code string) error {

	//查询是否已注册
	if ok := dao.IsUserExistByEmail(email); ok {
		return UserRegistered
	}
	//校验验证码是否正确
	err := MailSrvCtl.Verify(context.Background(), email, code)
	if err != nil {
		return ErrCodeVaild
	}

	user := model.User{
		Name:  "用户" + strings.Split(email, "@")[0],
		Email: email,
	}

	err = dao.InsertUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) RegisterByPassword(name string, password string) error {

	//查询是否已注册
	if ok := dao.IsUserExistByName(name); ok {
		return UserRegistered
	}

	//用加密后的密码，创建记录
	user := model.User{
		Name: name,
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user.Password = string(hash)
	err := dao.InsertUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserService) LoginByPassword(name string, password string) (model.User, error) {
	if ok := dao.IsUserExistByName(name); !ok {
		return model.User{}, UserNotFind
	}
	//拿到用户
	user, err := dao.FindUserByName(name)
	if err != nil {
		return model.User{}, err
	}
	//验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return model.User{}, PasswordErr
	}

	return user, nil
}

func ForgetPWD() {

}
