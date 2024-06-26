package model

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Token     string
	TokenDate time.Time // Token时间
	NickName  string    // 昵称
	Phone     string    // 电话
	Flag      string    // 启用标志(1-启用 0-停用)
}

// UserRefreshManagerToken
//
//	@Description:	修改指定用户的token数据
//	@param			token	数据格式	<token_value:timestamp>
//	@return			err
func UserRefreshManagerToken(db *gorm.DB, userId int64, token string) (err error) {
	res := db.Model(&User{}).Where("id = ?", userId).Update("manager_token", token)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("res.RowsAffected == 0")
	}
	return nil
}

// UserSelectIdByToken token查询用户数据 token = "HASH"
func UserSelectIdByToken(db *gorm.DB, token string) (user User, err error) {
	err = db.Table("users").
		Where("token LIKE ? AND flag = ?", token+":%", "1").Take(&user).Error
	return
}

// UserRefreshToken
//
//	@Description:	修改指定用户的token数据
//	@param			token	数据格式	<token_value:timestamp>
//	@return			err
func UserRefreshToken(db *gorm.DB, userId int64, token string) (err error) {
	res := db.Model(&User{}).Where("id = ?", userId).Update("token", token)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("res.RowsAffected == 0")
	}
	return nil
}

func UserSelectIdByManagerToken(db *gorm.DB, token string) (user User, err error) {
	err = db.Table("users").
		Where("manager_token LIKE ? AND flag = ?", token+":%", "1").Take(&user).Error
	return
}
