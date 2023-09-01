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
	ErrorUserExit      = errors.New("用户已存在")
	ErrorUserNotExit   = "用户不存在"
	ErrorPasswordWrong = errors.New("密码错误")
	ErrorGenIDFailed   = errors.New("创建用户ID失败")
	ErrorInvalidID     = "无效的ID"
	ErrorQueryFailed   = "查询数据失败"
	ErrorInsertFailed  = errors.New("插入数据失败")
)

func Register(req models.Account) (int64, string, error) {
	log.Println("调用了 service.register")

	isExistByName, err := dao.QueryAccountIsExistByName(req.Username)
	if err != nil {
		return 0, "", err
	}

	// 首先检查用户是否已存在
	if isExistByName {
		return 0, "", ErrorUserExit
	}

	// 用户不存在,进行注册

	// 插入User表
	u := models.User{
		ID:              0,
		Name:            req.Username,
		FollowCount:     0,
		FollowerCount:   0,
		IsFollow:        false,
		Avatar:          "https://www.helloimg.com/images/2023/08/25/oiG7nT.jpg",
		BackgroundImage: "https://www.helloimg.com/images/2023/08/25/oiGQUq.png",
		Signature:       "这家伙很懒，什么也没写",
		TotalFavorited:  "100",
		WorkCount:       0,
		FavoriteCount:   0,
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
	return req.UserId, token, nil
}

func Login(acc *models.Account) (*models.Account, string, error) {

	// 进行用户名和密码验证
	dbuser, err := LoginValidation(acc)
	if err != nil {
		return nil, "", err
	}

	// 创建token
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

func LoginValidation(acc *models.Account) (*models.Account, error) {

	// 从数据库中查询account信息
	dbuser, err := dao.QueryAccountByName(acc.Username)

	// 判断是否存在账号
	if errors.Is(err, dao.ErrorUserNotExit) {
		return nil, err
	}

	// 判断账号密码是否正确
	if acc.Password != dbuser.Password {
		return nil, ErrorPasswordWrong
	}

	return dbuser, nil

}
