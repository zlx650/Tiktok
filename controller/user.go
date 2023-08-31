package controller

import (
	"log"
	"net/http"
	"tiktok/dao"
	"tiktok/models"
	"tiktok/service"
  "tiktok/middleware"
	"github.com/gin-gonic/gin"
	"tiktok/util"
)

var userIdSequence = int64(1)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User models.User `json:"user"`
}

type PublishListResponse struct {
	StatusCode int            `json:"status_code"`
	StatusMsg  string         `json:"status_msg,omitempty"`
	UserID     int64          `json:"user_id,omitempty"`
	Token      string         `json:"token,omitempty"`
	VideoList  []util.Video `json:"video_list,omitempty"`
}

func Register(c *gin.Context) {
	var req models.RegisterForm

	// 从 URL 查询参数中获取用户名和密码
	username := c.Query("username")
	password := c.Query("password")

	// 将查询参数绑定到请求结构体
	req.UserName = username
	req.Password = password

   userId, token, err := service.Register(req)
    if err != nil {
        log.Println("注册失败", err)
        c.JSON(http.StatusOK, UserLoginResponse{
            Response: Response{StatusCode: 2, StatusMsg: "用户已存在"},
        })
        return
    }

    c.JSON(http.StatusOK, UserLoginResponse{
        Response: Response{StatusCode: 0},
        UserId:   userId,
        Token:    token,
    })
}


func Login(c *gin.Context) {
	log.Println("Login request received")

	u := &models.LoginForm{}

	log.Println("URL:", c.Request.URL.String())
	log.Println("Params:", c.Request.URL.Query())

	// 从 URL 参数中获取用户名和密码
	username := c.Query("username")
	password := c.Query("password")

	// 将用户名和密码绑定到登录表单
	u.UserName = username
	u.Password = password

	// 在这里可以进行参数验证逻辑
	if username == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 400,
			"status_msg":  "用户名和密码不能为空",
		})
		return
	}

	user, err := service.Login(u)
	if err != nil {
		log.Println("service.Login failed", err)
		if err.Error() == dao.ErrorUserNotExit {
			c.JSON(http.StatusOK, gin.H{
				"status_code": 3,
				"status_msg":  "用户不存在",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status_code": 1,
			"status_msg":  "请求参数错误",
		})
		return
	}

	// log.Println("User found:", user)

	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"user_id":     user.UserId,
		"token":       user.Token,
		"status_msg":  "登录成功",
	})
}

func UserInfo(c *gin.Context) {
    // var u models.UserForm
  
    log.Println("URL 参数 user_id:", c.Query("user_id"))
    log.Println("URL 参数 token:", c.Query("token"))
  
    // 手动解析参数
    token := c.Query("token")
    
    // 验证 token
    claims, err := middleware.ParseToken(token)
    if err != nil {
        log.Println("Token 验证失败", err)
        c.JSON(http.StatusOK, UserResponse{
            Response: Response{StatusCode: 5, StatusMsg: "Token 验证失败"},
        })
        return
    }
    
    // 使用解析后的 claims 数据获取用户信息
    userInfo, err := service.GetUserInfo(&models.UserForm{
        UserId: claims.UserId, // 使用解析后的用户 ID
        Token:  token,
    })
    if err != nil {
        log.Println("获取用户信息失败", err)
        c.JSON(http.StatusOK, UserResponse{
            Response: Response{StatusCode: 3, StatusMsg: "用户不存在"},
        })
        return
    }

    c.JSON(http.StatusOK, UserResponse{
        Response: Response{StatusCode: 0},
        User:     *userInfo,
    })
}



func GetPublishList(c *gin.Context) {
	// 从 URL 查询参数中获取 token
	token := c.Query("token")

	// 从本地存储中获取 user_id
	userID, ok := c.Get("user_id")
	log.Println("获取用户id", userID)
	if !ok {
		c.JSON(http.StatusOK, PublishListResponse{
			StatusCode: 6,
			StatusMsg:  "无效的用户 ID",
		})
		return
	}

	// 转换 userID 为 int64 类型
	userIDInt64, ok := userID.(int64)
	if !ok {
		c.JSON(http.StatusOK, PublishListResponse{
			StatusCode: 6,
			StatusMsg:  "无效的用户 ID",
		})
		return
	}
	log.Println("获取发布列表")
	// 调用 service 获取用户发布列表
	videoList, err := service.GetPublishList(userIDInt64)
	if err != nil {
		c.JSON(http.StatusOK, PublishListResponse{
			StatusCode: 1,
			StatusMsg:  "获取发布列表失败",
		})
		return
	}


	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"user_id":    userIDInt64,
		"token":      token,
		"video_list":  videoList,
		"status_msg":  "获取成功",
	})
}