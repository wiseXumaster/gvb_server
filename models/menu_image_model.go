package models

// MenuBannerModel 自定义菜单和背景图的连接表，方便排序
type MenuBannerModel struct {
	MenuID      uint        `gorm:"primaryKey" json:"menu_id"`
	BannerID    uint        `gorm:"primaryKey" json:"banner_id"`
	Sort        int         `gorm:"size:10" json:"sort"`
	MenuModel   MenuModel   `gorm:"foreignKey:MenuID;references:ID"`
	BannerModel BannerModel `gorm:"foreignKey:BannerID;references:ID"`
}
