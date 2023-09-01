package dao

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"tiktok/models"
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

// 根据用户Id查找用户是否存在
func FindUserById(user_id int64) (*models.User, error) {

	var user *models.User
	err := DB.Where("id = ?", user_id).First(&user).Error
	if err != nil {
		log.Println(err.Error())
	}
	return user, err
}

func QueryUserIsExistByName(username string) (bool, error) {
	var user models.Account

	if err := DB.Where("username = ?", username).First(&user).Error; err != nil {
		return false, err
	}

	// user存在
	if user.ID != 0 {
		return true, nil
	}

	// user不存在
	return false, nil

}

// 根据用户名查找用户是否存在
func FindUserByName(username string) (*models.Account, error) {
	var user models.Account

	query := "SELECT * FROM account WHERE username = ?"
	result := DB.Raw(query, username).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Println("User not found:", username)
			return nil, errors.New(ErrorUserNotExit)
		}
		log.Println("Error querying database:", result.Error)
		return nil, result.Error
	}

	return &user, nil
}

// 添加新用户
func InsertUser(user models.User) (int64, error) {
	if err := DB.Create(&user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}

func InsertAccount(user models.Account) error {
	err := DB.Create(&user).Error
	return err
}

func Login(username string) (*models.Account, error) {
	var user models.Account
	result := DB.Where(&models.Account{Username: username}).Find(&user)
	if result.RowsAffected > 0 {
		return &user, nil
	}
	return nil, errors.New(ErrorUserNotExit)
}
