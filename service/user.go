package service

import (
	"errors"
	"log"
	"tiktok/dao"
	"tiktok/middleware"
	"tiktok/models"
	// "gorm.io/gorm"
)

func Register(req models.Account) (int64, string, error) {
	log.Println("调用了 service.register")

	// 首先检查用户是否已存在
	existingUser, err := dao.FindUserByName(req.Username)
	// log.Println("打印一下这个错误",err)

	if existingUser != nil {
		return 0, "", errors.New("用户已存在")
	}
	log.Println("111")
	// 用户不存在，进行注册
	acc := models.Account{
		Username: req.Username,
		Password: req.Password,
	}

	err = dao.InsertRegisterForm(&acc)
	if err != nil {
		return 0, "", errors.New("创建失败")
	} else {
		u := models.User{
			Name: req.Username,
		}
		errs := dao.InsertUser(&u)
		if errs != nil {
			return 0, "", errors.New("创建失败")
		}
	}

	
	// 生成 token
	token, err := middleware.CreateToken(acc.ID,acc.Username, acc.Password)
	if err != nil {
		return 0, "", err
	}

	// log.Println("调用了 service.register，成功返回用户 ID 和 token")
	// 在注册成功后，返回用户 ID 和 token
	return acc.ID, token, nil
}

func Login(acc *models.Account) (*models.Account, string, error) {
	// 进行用户名密码验证
	user, err := dao.Login(acc.Username)
	if err != nil {
		return nil, "", err
	}
	// log.Println("这是service里面的user2：",user)
	
	// TODO 这里的UserId为空，需要生成
	token, err := middleware.CreateToken(user.UserId, user.Username, user.Password)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func GetUserInfo(acc *models.Account) (*models.User, error) {
	userInfo, err := dao.FindUserById(acc.UserId)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}
