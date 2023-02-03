package handler

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"tiktok.service/biz/dal/mysql"
	"tiktok.service/biz/model"
	utils2 "tiktok.service/biz/utils"
)

// Register 注册用户事务
func Register(ctx context.Context, c *app.RequestContext) {
	//临时结构体
	var registerStruct struct {
		Username string `form:"user_name" json:"user_name" query:"username" vd:"(len($) > 0 && len($) < 128); msg:'Illegal format'"`
		Password string `form:"password" json:"password" query:"password" vd:"(len($) > 0 && len($) < 128); msg:'Illegal format'"`
	}

	//传入数据绑定至临时结构体
	if err := c.BindAndValidate(&registerStruct); err != nil {
		c.JSON(http.StatusOK, utils.H{
			"message": err.Error(),
			"code":    http.StatusBadRequest,
		})
		return
	}

	//在sql中查找用户名
	users, err := mysql.FindUserByName(registerStruct.Username)
	if err != nil {
		c.JSON(http.StatusOK, utils.H{
			"message": err.Error(),
			"code":    http.StatusBadRequest,
		})
		return
	}

	//用户名称重复则报错
	if len(users) != 0 {
		c.JSON(http.StatusOK, utils.H{
			"message": "user already exists",
			"code":    http.StatusBadRequest,
		})
		return
	}

	//创建用户
	if err = mysql.CreateUser([]*model.User{
		{
			UserName: registerStruct.Username,
			Password: utils2.MD5(registerStruct.Password),
		},
	}); err != nil {
		c.JSON(http.StatusOK, utils.H{
			"message": err.Error(),
			"code":    http.StatusBadRequest,
		})
		return
	}

	c.JSON(http.StatusOK, utils.H{
		"message": "success",
		"code":    http.StatusOK,
	})
}
