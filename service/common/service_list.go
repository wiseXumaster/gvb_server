package common

import (
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
)

type Option struct {
	models.PageInfo
	Debug bool
}

func ComList[T any](model T, option Option) (list []T, count int64, err error) {
	DB := global.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}
	if option.PageInfo.Page < 1 {
		option.PageInfo.Page = 1
	}
	if option.PageInfo.Limit <= 0 {
		option.PageInfo.Limit = 10 // 设置一个默认值
	}
	offset := (option.PageInfo.Page - 1) * option.PageInfo.Limit

	DB.Model(&model).Count(&count)
	err = DB.Limit(option.PageInfo.Limit).Offset(offset).Find(&list).Error

	//var imageList []models.BannerModel

	//count := global.DB.Find(&imageList).RowsAffected
	// 使用 Count 来获取总数
	//global.DB.Model(&models.BannerModel{}).Count(&count)

	//if offset < 0 {
	//	offset = 0
	//}

	return list, count, err
}
