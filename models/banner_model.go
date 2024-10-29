package models

import (
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models/ctype"
	"os"
)

type BannerModel struct {
	MODEL
	Path      string          `json:"path"`                        // 图片路径
	Hash      string          `json:"hash"`                        // 图片的hash值，用于判断重复图片
	Name      string          `gorm:"size:200" json:"name"`        // 图片名称
	ImageType ctype.ImageType `gorm:"default:3" json:"image_type"` // 图片的类型,本地还是七牛
}

// 在同一个事务中更新数据
func (b *BannerModel) BeforeDelete(tx *gorm.DB) (err error) {
	if b.ImageType == ctype.Local {
		// 本地的图片，尝试删除文件
		if _, err := os.Stat(b.Path); err == nil {
			err = os.Remove(b.Path)
			if err != nil {
				global.Log.Error("删除图片文件失败:", err)
				return err
			}
			global.Log.Info("成功删除本地图片文件:", b.Path)
		} else if os.IsNotExist(err) {
			global.Log.Warn("文件不存在:", b.Path)
		}
	}
	return nil
}
