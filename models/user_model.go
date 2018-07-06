package models

//"Watermelon/common/db/mysql"
import (
	"Watermelon/db/mysql"
	"fmt"
)

type User struct {
	UserID    int64  `gorm:"column:UserID;primary_key;AUTO_INCREMENT;not null;"`
	Name      string `gorm:"column:Name"`
	Age       int    `gorm:"column:Age"`
	CreatedAt string `gorm:"column:CreatedAt"`
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

//条件查询
func (user User) GetEntity() (User, int64) {
	var count int64
	var entity = User{}
	mysql.Orm.Where(&User{UserID: user.UserID}).Find(&entity).Count(&count)
	return entity, count
}

//更新一条记录
func (user User) Update() {

	mysql.Orm.Model(User{}).Updates(User{UserID: 1})

}

//返回表名
func (user User) TableName() string {
	return "User"
}
