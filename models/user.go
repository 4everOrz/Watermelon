package models

import (
	"Watermelon/db/mysql"
	"fmt"
)

type User struct {
	UserID    int64  `gorm:"column:UserID;primary_key;AUTO_INCREMENT;not null;"  json:"userid"`
	Name      string `gorm:"column:Name"  json:"name"`
	LoginName string `gorm:"column:LoginName"  json:"loginname"`
	Age       int    `gorm:"column:Age"  json:"age"`
	Email     string `gorm:"column:Email" json:"email"`
	Password  string `gorm:"column:Password" json:"password"`
	Token     string `gorm:"column:Token" json:"token"`
	CreatedAt string `gorm:"column:CreatedAt" json:"creatat"`
	UpdateAt  string `gorm:"column:UpdateAt" json:"updateat"`
}

func init() {
	if !mysql.Orm.HasTable(&User{}) {
		if err := mysql.Orm.Set("gorm:user", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&User{}).Error; err != nil {
			fmt.Println("Create table failed,error:" + err.Error())
		}
	}
}

//添加一条记录
func (user *User) Create() error {
	if err := mysql.Orm.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

//删除一条记录
func (user User) Delete() error {
	if err := mysql.Orm.Where(&User{UserID: 1}).Delete(User{}).Error; err != nil {
		return err
	}
	return nil
}

//条件查询ID
func (user User) GetEntityByID() (User, error) {
	var entity = User{}
	err := mysql.Orm.Where(&User{UserID: user.UserID}).Find(&entity).Error
	return entity, err
}

//条件查询LoginName
func (user User) GetEntityByLoginName() (User, error) {
	var entity = User{}
	err := mysql.Orm.Where(&User{LoginName: user.LoginName}).Find(&entity).Error
	return entity, err
}

//记录是否存在
func (user User) EntityExist() (bool, error) {
	var count int64
	var allrignt bool
	err := mysql.Orm.Where("LoginName= ?", user.LoginName).Find(&User{}).Count(&count).Error
	if count == 0 {
		allrignt = true
	} else {
		allrignt = false
	}
	return allrignt, err
}

//更新一条记录
func (user User) UpdateUser() error {
	err := mysql.Orm.Model(User{}).Updates(user).Error
	return err
}

//返回表名
func (user User) TableName() string {
	return "User"
}
