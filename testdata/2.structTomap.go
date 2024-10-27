package main

import (
	"fmt"
	"github.com/fatih/structs"
	"gvb_server/models"
)

type AdvertModel struct {
	models.MODEL `structs:"-"`
	Title        string `json:"title" structs:"title"`     // 显示的标题
	Href         string `json:"href" structs:"href"`       // 跳转链接
	Images       string `json:"images" structs:"images"`   // 图片
	IsShow       bool   `json:"is_show" structs:"is_show"` // 是否展示
}

func main() {
	u1 := AdvertModel{
		Title:  "老t博客5",
		Href:   "http://sly9bn5nh.hn-bkt.clouddn.com/gvb/20241026152740__QQ%E5%9B%BE%E7%89%8720240908215741.png",
		Images: "http://sly9bn5nh.hn-bkt.clouddn.com/gvb/20241026154525_1231231233.jpg",
		IsShow: false,
	}
	m3 := structs.Map(u1)
	fmt.Println(m3)
}
