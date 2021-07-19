package mysql

import (
	"bluebell/model"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
)

//将每一步数据库的操作封装成函数
//待logic层调用业务需求

//检查用户是否已存在

var secret = "liwenzhou.com"

func CheckUserExist(username string) (err error) {
	//定义一条sql
	sqlStr := `select count(user_id) from user where username = ?`
	var count int64
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("user exiting")
	}

	return

}

//向数据库中加入一条数据
func InsertUser(user *model.User) (err error) {
	//对密码加密 返回加密后的密文
	user.Password = EncryptPassword(user.Password)
	//执行sql语句入库
	sqlStr := `insert into user(user_id,username,password) values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

//encryptPassword 密码加密
func EncryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))

}

func Login(user *model.User) (err error) {
	oPassword := user.Password //获取用户登录密码
	//查询数据库
	sqlStr := `select user_id,username,password from user where username =?`
	err = db.Get(user, sqlStr, user.Username)

	//查询为空 用户不存在
	if err == sql.ErrNoRows {
		return errors.New("用户名不存在")
	}
	if err != nil {
		return
	}

	password := EncryptPassword(oPassword)
	if password != oPassword {
		return errors.New("用户名密码错误")
	}
	return

}
