/**
* @Author: yexichen
* @Date:2023/7/11-19:23
* @Desc
**/

package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type EnterRoomResp struct {
	Play  bool   `json:"play"`
	Enter bool   `json:"enter"`
	Info  string `json:"info"`
	Num   int    `json:"num"`
}

func enterroom(c *gin.Context) {
	roomId := c.PostForm("roomId")
	if roomcount[roomId] >= 2 {
		c.JSON(200, gin.H{
			"status": true,
		})
	} else {
		c.JSON(200, gin.H{
			"status": false,
		})
	}
	fmt.Println(roomcount[roomId])
}

func checkroom(c *gin.Context) {
	roomId := c.PostForm("roomId")
	if roomcount[roomId] >= 2 {
		c.JSON(200, gin.H{
			"status": true,
		})
	} else {
		c.JSON(200, gin.H{
			"status": false,
		})
	}
	fmt.Println(roomcount[roomId])
}
