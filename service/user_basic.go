package service

import (
	
)

func GetUserList(c *gin.Context) {
	userList := make([]*models.UserBasic, 10)
	userList, err := models.UserBasic{}.List()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "获取失败",
		})
	}

	c.JSON(200, gin.H{
		"message": userList
	})
}