/**
* @Author: yexichen
* @Date:2023/7/11-19:23
* @Desc
**/

package service

import (
	"database/sql"
	"server/dao"
	"server/model"
)

func IsUserNameExist(Username string) (bool, error) {
	User, err := dao.SelectUserByUserName(Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}
		return false, err
	}
	if User.UserName != Username {
		return true, nil
	}
	return false, nil
}

func NewUser(User model.User) error {
	return dao.InsertUser(User)
}

func IsUserNameAndMailRight(Userrname, Userrmail string) (bool, error) {
	Userr, err := dao.SelectUserByUserName(Userrname)
	if err != nil {
		return false, err
	}
	if Userr.UserName == Userrname && Userr.UserMail == Userrmail {
		return true, nil
	} else {
		return false, nil
	}
}

func GetUserInfo(username string) (model.User, error) {
	return dao.SelectUserByUserName(username)
}

func AddWinCount(userid int) error {
	return dao.UpdateWinCount(userid)
}

func SearchWinCount(userid int) int {
	return dao.SelectWinCount(userid)
}
