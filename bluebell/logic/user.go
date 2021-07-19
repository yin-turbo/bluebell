package logic

import (
	"bluebell/dao/mysql"
	"bluebell/model"
	"bluebell/pkg/snowflake"

	"go.uber.org/zap"
)

func SignUp(p *model.ParamSignUp) (err error) {
	//1.判断用户是否存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}

	//用户不存在 生成uuid
	userID := snowflake.GenID()
	//构造一个user实体
	user := &model.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}

	//保存至数据库
	return mysql.InsertUser(user)
}

//用户登录
func Login(p *model.ParamLogin) (err error) {
	user := &model.User{
		Username: p.Username,
		Password: p.Password,
	}
	
	 
	if err = mysql.Login(user); err != nil {
		zap.L().Error("login failed", zap.Error(err))
	}

	return
}
