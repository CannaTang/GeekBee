package api

import (
	"GeekBee5/api/middleware"
	"GeekBee5/dao"
	"GeekBee5/model"
	"GeekBee5/utils"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

func register(c *gin.Context) {
	form := model.User{}
	if err := c.ShouldBind(&form); err != nil {
		fmt.Println(err)
		utils.RespSuccess(c, "verification failed")
		return
	}
	// 传入用户名和密码
	username := c.PostForm("username")
	password := c.PostForm("password")
	phone := c.PostForm("phone")

	// 验证用户名是否重复
	flag := dao.SelectUser(username)
	fmt.Println(flag)
	if flag {
		// 以 JSON 格式返回信息
		utils.RespFail(c, "user already exists")
		return
	}

	dao.AddUser(username, password, phone)
	// 以 JSON 格式返回信息
	utils.RespSuccess(c, "add user successful")
}

func login(c *gin.Context) {
	if err := c.ShouldBind(&model.User{}); err != nil {
		utils.RespFail(c, "verification failed")
		return
	}
	// 传入用户名和密码
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 验证用户名是否存在
	flag := dao.SelectUser(username)
	// 不存在则退出
	if !flag {
		// 以 JSON 格式返回信息
		utils.RespFail(c, "user doesn't exists")
		return
	}

	// 查找正确的密码
	selectPassword := dao.SelectPasswordFromUsername(username)
	// 若不正确则传出错误
	if selectPassword != password {
		// 以 JSON 格式返回信息
		utils.RespFail(c, "wrong password")
		return
	}

	// 正确则登录成功
	// 创建一个我们自己的声明
	claim := model.MyClaims{
		Username: username, // 自定义字段
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * 1).Unix(), // 过期时间
			Issuer:    "Tzz",                                  // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	tokenString, _ := token.SignedString(middleware.Secret)
	utils.RespSuccess(c, tokenString)
}

func getUsernameFromToken(c *gin.Context) {
	username, _ := c.Get("username")
	utils.RespSuccess(c, username.(string))
}

func modifyUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	newName := c.PostForm("newName")
	newPassword := c.PostForm("newPassword")
	//检验用户名存在
	flag := dao.SelectUser(username)
	if !flag {
		utils.RespFail(c, "user doesn't exists")
		return
	}
	//检验密码正确
	correctPassword := dao.SelectPasswordFromUsername(username)
	if password != correctPassword {
		utils.RespFail(c, "wrong password")
		return
	}
	//检验目标用户名
	flag = dao.SelectUser(newName)
	if flag {
		utils.RespFail(c, "username already exists")
		return
	}
	dao.DelUser(username)
	dao.AddUser(newName, newPassword, dao.PhoneNum[username])
	dao.SaveAll()
	utils.RespSuccess(c, "Modify User successfully")
}

func FindPassword(c *gin.Context) {
	username := c.PostForm("username")
	//检验用户名存在
	flag := dao.SelectUser(username)
	if !flag {
		utils.RespFail(c, "user doesn't exists")
		return
	}
	phone := c.PostForm("phone")
	if dao.SelectPhoneFromUsername(username, phone) {
		password := dao.SelectPasswordFromUsername(username)
		utils.RespSuccess(c, "your password is "+password)
	} else {
		utils.RespFail(c, "wrong phone number")
	}
}

func MakeComment(c *gin.Context) {
	username := c.PostForm("username")
	comment := c.PostForm("comment")
	dao.AddComment(username, comment)
	utils.RespSuccess(c, "Add comment successfully")
}
