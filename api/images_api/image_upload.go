package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/common/res"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/plugins/huawei"
	"gvb_server/plugins/qiniu"
	"gvb_server/utils"
	"io"
	"io/fs"
	"os"
	"path"
	"strings"
)

type FileUploadResponse struct {
	FileName  string `json:"file_name"`  // 文件名
	IsSuccess bool   `json:"is_success"` // 是否上传成功
	Msg       string `json:"msg"`        // 消息
}

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

//上传单个图片,返回图片的url
func (ImagesApi) ImageUploadView(c *gin.Context) {
	//fileHeader, err := c.FormFile("image")
	//if err != nil {
	//	res.FailWithMessage(err.Error(), c)
	//	return
	//}
	//fmt.Println(fileHeader.Header)
	//fmt.Println(fileHeader.Size)
	//fmt.Println(fileHeader.Filename)
	// 上传多个图片
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	//fmt.Println(form.File)
	fileList, ok := form.File["images"]
	if !ok {
		res.FailWithMessage("不存在的文件", c)
		return
	}

	basePath := global.Config.Upload.Path
	// 判断路径是否存在
	_, err = os.ReadDir(basePath)
	if err != nil {
		// 不存在就创建
		err = os.MkdirAll(basePath, fs.ModePerm)
		if err != nil {
			global.Log.Error(err)
		}
	}

	var resList []FileUploadResponse
	for _, file := range fileList {
		fileName := file.Filename
		suffix := strings.ToLower(path.Ext(fileName))
		//fmt.Println(suffix)
		if !utils.InList(suffix, WhiteImageList) {
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       "非法文件",
			})
			continue
		}

		filePath := path.Join(basePath, file.Filename)
		// 判断大小
		//fmt.Println(filePath, float64(file.Size)/float64(1024*1024))
		size := float64(file.Size) / float64(1024*1024)
		// 超过设定大小
		if size >= float64(global.Config.Upload.Size) {
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       fmt.Sprintf("图片大小超过设定大小,当前大小为:%.2fMB,设定大小为:%dMB ", size, global.Config.Upload.Size),
			})
			continue
		}

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
		if err == nil {
			//找到了
			resList = append(resList, FileUploadResponse{
				FileName:  bannerModel.Path,
				IsSuccess: false,
				Msg:       "图片已存在",
			})
			continue
		}
		if global.Config.HuaWei.Enable {
			//var _ string
			filePath, err = huawei.UploadImage(byteData, fileName, global.Config.HuaWei.Prefix)
			if err != nil {
				global.Log.Error(err)
				continue
			}
			// 上传成功
			resList = append(resList, FileUploadResponse{
				FileName:  filePath,
				IsSuccess: true,
				Msg:       "上传华为云成功",
			})
			// 图片入库
			global.DB.Create(&models.BannerModel{
				Path:      filePath,
				Hash:      imageHash,
				Name:      fileName,
				ImageType: ctype.HuaWei,
			})
		}
		if global.Config.QiNiu.Enable {
			//var _ string
			filePath, err = qiniu.UploadImage(byteData, fileName, global.Config.QiNiu.Prefix)
			if err != nil {
				global.Log.Error(err)
				continue
			}
			// 上传成功
			resList = append(resList, FileUploadResponse{
				FileName:  filePath,
				IsSuccess: true,
				Msg:       "上传七牛云成功",
			})
			// 图片入库
			global.DB.Create(&models.BannerModel{
				Path:      filePath,
				Hash:      imageHash,
				Name:      fileName,
				ImageType: ctype.QiNiu,
			})
		}
		res.OkWithData(resList, c)

		return
		// 上传
		err = c.SaveUploadedFile(file, filePath)
		// 上传失败
		if err != nil {
			global.Log.Error(err)
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       err.Error(),
			})
			continue
		}
		// 上传成功
		resList = append(resList, FileUploadResponse{
			FileName:  filePath,
			IsSuccess: true,
			Msg:       "上传成功",
		})
		// 图片入库
		global.DB.Create(&models.BannerModel{
			Path:      filePath,
			Hash:      imageHash,
			Name:      fileName,
			ImageType: ctype.Local,
		})
		//err := c.SaveUploadedFile(file, filePath)
		//if err != nil {
		//	global.Log.Error(err)
		//	continue
		//}
		//fmt.Println(file.Header)
		//fmt.Println(file.Size)
		//fmt.Println(file.Filename)
	}
	res.OkWithData(resList, c)
}
