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
	LoginByEmail(email string, code string) (model.User, error)
	LoginByPassword(name string, password string) (model.User, error)
}

type UserService struct {
}

func (u *UserService) LoginByEmail(email string, code string) (model.User, error) {

	//校验验证码是否正确
	err := MailSrvCtl.Verify(context.Background(), email, code)
	if err != nil {
		return model.User{}, ErrCodeVaild
	}
	//查询是否已注册
	if ok := dao.IsUserExistByEmail(email); ok {
		//已注册过且验证码正确，直接登录
		user, err1 := dao.FindUserByEmail(email)
		if err1 != nil {
			return model.User{}, err
		}
		return user, nil
	} else {
		user := model.User{
			Name:  "用户" + strings.Split(email, "@")[0],
			Email: email,
		}

		err = dao.InsertUser(user)
		if err != nil {
			return model.User{}, err
		}
		return user, nil
	}

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
