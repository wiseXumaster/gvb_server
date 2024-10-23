package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/common/res"
	"gvb_server/global"
	"io/fs"
	"os"
	"path"
)

type FileUploadResponse struct {
	FileName  string `json:"file_name"`  // 文件名
	IsSuccess bool   `json:"is_success"` // 是否上传成功
	Msg       string `json:"msg"`        // 消息
}

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
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
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
		// 上传
		err := c.SaveUploadedFile(file, filePath)
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
