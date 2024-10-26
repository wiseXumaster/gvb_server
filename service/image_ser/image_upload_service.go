package image_ser

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/plugins/huawei"
	"gvb_server/plugins/qiniu"
	"gvb_server/utils"
	"io"
	"mime/multipart"
	"path"
	"strings"
)

var (
	WhiteImageList = []string{
		".jpg",
		".png",
		".jpeg",
		".ico",
		".tiff",
		".gif",
		".svg",
		".webp",
	}
)

type FileUploadResponse struct {
	FileName  string `json:"file_name"`  // 文件名
	IsSuccess bool   `json:"is_success"` // 是否上传成功
	Msg       string `json:"msg"`        // 消息
}

// 图片上传的方法
func (ImageService) ImageUploadService(file *multipart.FileHeader) (res FileUploadResponse) {
	fileName := file.Filename

	//本地路径
	basePath := global.Config.Upload.Path
	//拼接路径
	filePath := path.Join(basePath, file.Filename)

	//上传的名字是路径
	res.FileName = filePath

	//默认失败
	res.IsSuccess = false

	//分割后缀
	suffix := strings.ToLower(path.Ext(fileName))

	// 文件白名单判断
	if !utils.InList(suffix, WhiteImageList) {
		res.Msg = "非法文件"
		return res
	}

	// 判断大小
	size := float64(file.Size) / float64(1024*1024)
	// 超过设定大小
	if size >= float64(global.Config.Upload.Size) {
		res.Msg = fmt.Sprintf("图片大小超过设定大小,当前大小为:%.2fMB,设定大小为:%dMB ", size, global.Config.Upload.Size)
		return res
	}

	//读取文件内容,计算md5
	fileObj, err := file.Open()
	if err != nil {
		global.Log.Error(err)
	}
	byteData, err := io.ReadAll(fileObj)
	//md5String := utils.Md5(byteData)
	imageHash := utils.Md5(byteData)
	//fmt.Println(md5String)
	//去数据库中查询这个图片是否存在
	var bannerModel models.BannerModel
	err = global.DB.Take(&bannerModel, "hash = ?", imageHash).Error
	// 空说明没有报错,也就是图片已存在
	if err == nil {
		//找到了
		res.FileName = bannerModel.Path
		res.Msg = "图片已存在"
		return res
	}
	fileType := ctype.Local
	if global.Config.HuaWei.Enable {
		//var _ string
		filePath, err = huawei.UploadImage(byteData, fileName, global.Config.HuaWei.Prefix)
		if err != nil {
			global.Log.Error(err)
			res.Msg = err.Error()
			return res
		}
		// 上传成功
		res.FileName = filePath
		res.Msg = "上传华为云成功"
		fileType = ctype.HuaWei
	}
	if global.Config.QiNiu.Enable {
		//var _ string
		filePath, err = qiniu.UploadImage(byteData, fileName, global.Config.QiNiu.Prefix)
		if err != nil {
			global.Log.Error(err)
			res.Msg = err.Error()
			return res
		}
		// 上传成功
		res.FileName = filePath
		res.Msg = "上传七牛云成功"
		fileType = ctype.QiNiu
	}
	// 图片入库
	err = global.DB.Create(&models.BannerModel{
		Path:      filePath,
		Hash:      imageHash,
		Name:      fileName,
		ImageType: fileType,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.Msg = err.Error()
		return res
	}
	res.IsSuccess = true
	return res
}
