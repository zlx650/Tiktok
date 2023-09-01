package service

import (
	"errors"
	"log"
	// "os/user"
	"tiktok/dao"
	"tiktok/middleware"
	"tiktok/models"
	// "gorm.io/gorm"
)

var (
	ErrorUserExit      = "用户已存在"
	ErrorUserNotExit   = "用户不存在"
	ErrorPasswordWrong = "密码错误"
	ErrorGenIDFailed   = errors.New("创建用户ID失败")
	ErrorInvalidID     = "无效的ID"
	ErrorQueryFailed   = "查询数据失败"
	ErrorInsertFailed  = errors.New("插入数据失败")
)

func Register(req models.Account) (int64, string, error) {
	log.Println("调用了 service.register")

	// 首先检查用户是否已存在
	existingUser, err := dao.FindUserByName(req.Username)
	// log.Println("打印一下这个错误",err)

	if existingUser != nil {
		return 0, "", errors.New("用户已存在")
	}

	
	// 用户不存在，进行注册
	u := models.User{
			Name: req.Username,
	}
	// 插入User表
	err = dao.InsertUser(&u)
	if err != nil {
		return 0, "", errors.New("创建失败")
	}
	
	// 插入Account表
	eu, _ := dao.FindUserByName(u.Name)
	acc := models.Account{
		UserId: eu.UserId,
		Username: eu.Username,
		Password: req.Password,
	}
	err = dao.InsertAccount(&acc)
	if err != nil {
		return 0, "", errors.New("创建失败")
	}

	
	// 生成 token
	token, err := middleware.CreateToken(acc.UserId,acc.Username, acc.Password)
	if err != nil {
		return 0, "", err
	}

	// log.Println("调用了 service.register，成功返回用户 ID 和 token")
	// 在注册成功后，返回用户 ID 和 token
	return acc.ID, token, nil
}

func Login(acc *models.Account) (*models.Account, string, error) {
	// 进行用户名密码验证
	dbuser, err := dao.Login(acc.Username)
	if dbuser.UserId == 0 {
		return nil, "", err;
	}
	
	if (dbuser.Password != acc.Password) {
		return nil, "", errors.New(ErrorPasswordWrong);
	}

	if err != nil {
		return nil, "", err
	}
	// log.Println("这是service里面的user2：",user)

	
	// TODO 这里的UserId为空，需要生成
	token, err := middleware.CreateToken(dbuser.UserId, dbuser.Username, dbuser.Password)
	if err != nil {
		return nil, "", err
	}

	return dbuser, token, nil
}

func GetUserInfo(acc *models.Account) (*models.User, error) {
	userInfo, err := dao.FindUserById(acc.UserId)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}
