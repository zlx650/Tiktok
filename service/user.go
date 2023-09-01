package service

import (
	"errors"
	"log"
	"tiktok/util"

	// "os/user"
	"tiktok/dao"
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

	isExistByName, err := dao.QueryUserIsExistByName(req.Username)
	if err != nil {
		return 0, "", err
	}

	// 首先检查用户是否已存在
	if isExistByName {
		return 0, "", errors.New("用户已存在")
	}

	// 用户不存在,进行注册

	// 插入User表
	u := models.User{
		Name: req.Username,
	}
	userId, err := dao.InsertUser(u)
	if err != nil {
		return 0, "", err
	}

	// 插入Account表
	req.UserId = userId
	err = dao.InsertAccount(req)
	if err != nil {
		return 0, "", err
	}

	// 生成 token
	token, err := util.CreateToken(req.UserId, req.Username, req.Password)
	if err != nil {
		return 0, "", err
	}

	// log.Println("调用了 service.register，成功返回用户 ID 和 token")
	// 在注册成功后，返回用户 ID 和 token
	return req.ID, token, nil
}

func Login(acc *models.Account) (*models.Account, string, error) {
	// 进行用户名密码验证
	dbuser, err := dao.Login(acc.Username)
	if err != nil || dbuser.UserId == 0 {
		return nil, "", err
	}

	if dbuser.Password != acc.Password {
		return nil, "", errors.New(ErrorPasswordWrong)
	}

	if err != nil {
		return nil, "", err
	}
	// log.Println("这是service里面的user2：",user)

	// TODO 这里的UserId为空，需要生成
	token, err := util.CreateToken(dbuser.UserId, dbuser.Username, dbuser.Password)
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
