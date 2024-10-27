package common

import (
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
)

type Option struct {
	models.PageInfo
	Debug bool
	//ModelType string // 新增字段，标识查询类型，比如 "advertisement" 或 "image"
}

func ComList[T any](model T, option Option) (list []T, count int64, err error) {
	DB := global.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}
	if option.Page < 1 {
		option.Page = 1
	}
	if option.PageInfo.Limit <= 0 {
		option.Limit = 10 // 设置一个默认值
	}
	offset := (option.PageInfo.Page - 1) * option.PageInfo.Limit
	if option.Sort == "" {
		option.Sort = "created_at desc"
	}

	// 构建查询条件
	query := DB.Model(&model)

	//// 如果是广告查询，加入 is_show 条件
	//if option.ModelType == "AdvertModel" {
	//	query = query.Where("is_show = ?", true)
	//}

	// 如果 model 中包含 is_show 字段且为 true/false，则根据其值进行筛选
	//if advertModel, ok := any(model).(models.AdvertModel); ok {
	//	query = query.Where("is_show = ?", advertModel.IsShow)
	//}
	//
	// 检查传入的 model 是否为 AdvertModel
	if advertModel, ok := any(model).(models.AdvertModel); ok {
		// 如果不是 admin 请求，则根据 isShow 决定是否过滤 is_show
		if advertModel.IsShow {
			query = query.Where("is_show = ?", true)
		}
	}
	//DB.Model(&model).Count(&count)
	// 获取总记录数
	if err := query.Model(&model).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	err = query.Limit(option.PageInfo.Limit).Offset(offset).Order(option.Sort).Find(&list).Error
	return list, count, err
}
