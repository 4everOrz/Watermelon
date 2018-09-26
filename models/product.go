package models

import (
	"Watermelon/db/mysql"
	"fmt"
)

type Product struct {
	ProductID     int64  `gorm:"column:ProductID;primary_key;AUTO_INCREMENT;not null;"  json:"productid"`
	ProductNumber string `gorm:"column:ProductNumber" json:"productnumber"`
	ProductName   string `gorm:"column:ProductName" json:"productname"`
	Price         string `gorm:"column:Price"  json:"price"`
	CreatorID     int64  `gorm:"column:CreatorID" json:"creatorid"`
	CreateAt      string `gorm:"column:CreateAt"`
	UpdateAt      string `gorm:"column:UpdateAt"`
}

func init() {
	if !mysql.Orm.HasTable(&Product{}) {
		if err := mysql.Orm.Set("gorm:product", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Product{}).Error; err != nil {
			fmt.Println("Create table failed,error:" + err.Error())
		}
	}
}

//通过ProductID搜索
func (this *Product) GetOneByProductID(pid int64) error {
	if err := mysql.Orm.First(&this, "ProductID = ?", pid).Error; err != nil {
		return err
	}
	return nil
}

//通过CreatID搜索（sql语句）
func (this *Product) GetOnehByCreatorID(creator int64) Product {
	var product Product
	mysql.Orm.Raw("Select * from product where CreatorID = ?", creator).Scan(&product)
	return product
}

//增加新数据
func (this *Product) InsertOne() error {
	if err := mysql.Orm.Create(&this).Error; err != nil {
		return err
	}
	return nil
}

//更新某条数据
func (this *Product) UpdateOne() error {
	if err := mysql.Orm.Model(Product{}).Updates(this).Error; err != nil {
		return err
	}
	return nil
}
