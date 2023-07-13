/**
* @Author: yexichen
* @Date:2023/7/11-19:23
* @Desc
**/

package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/service"
	"server/tool"
)

func login(c *gin.Context) {
	userName := c.PostForm("UserName")
	userMail := c.PostForm("UserMail")

	//check
	_, err := service.IsUserNameExist(userName)
	if err != nil {
		c.JSON(500, gin.H{
			"status": false,
			"info":   "服务器错误",
		})
		tool.Logger.Info("check username exist err",
			zap.Error(err))
		return
	}

	//可以创建新账户

	isr, errr := service.IsUserNameAndMailRight(userName, userMail)
	if errr != nil {
		c.JSON(500, gin.H{
			"status": false,
			"info":   "服务器错误",
		})
		tool.Logger.Info("check username and mail err",
			zap.Error(err))
		return
	}
	if !isr {
		c.JSON(200, gin.H{
			"code": 200,
			"err":  "姓名或邮箱错误",
		})
		return
	}

	c.JSON(200, gin.H{
		"code":     200,
		"username": userName,
	})
}
