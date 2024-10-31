package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/common/res"
	"gvb_server/global"
	"gvb_server/service"
	"gvb_server/service/image_ser"
	"io/fs"
	"os"
)

// ImageUploadView 上传图片
// @Tags 图片管理
// @Summary 上传单个或多个图片
// @Description 上传图片并返回上传成功的图片 URL
// @Accept multipart/form-data
// @Param images formData file true "上传的图片文件，支持多个文件"
// @Produce json
// @Success 200 {object} res.Response{data=[]image_ser.FileUploadResponse} "图片上传成功，返回图片 URL 列表"
// @Failure 400 {object} res.Response{msg=string} "请求解析错误或文件不存在"
// @Router /api/images [post]
func (ImagesApi) ImageUploadView(c *gin.Context) {
	// 上传多个图片
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println(form.File)
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

	var resList []image_ser.FileUploadResponse
	for _, file := range fileList {
		// 上传文件
		serviceRes := service.ServiceApp.ImageService.ImageUploadService(file)

		if !serviceRes.IsSuccess {
			resList = append(resList, serviceRes)
			continue
		}

		if global.Config.QiNiu.Enable || global.Config.HuaWei.Enable {
			resList = append(resList, serviceRes)
			continue
		}

		if !global.Config.QiNiu.Enable && !global.Config.HuaWei.Enable {
			//本地还得保存
			global.Log.Info(serviceRes.FileName)
			// 上传
			err = c.SaveUploadedFile(file, serviceRes.FileName)
			// 上传失败
			if err != nil {
				global.Log.Error(err)
				serviceRes.Msg = err.Error()
				serviceRes.IsSuccess = false
				resList = append(resList, serviceRes)
				continue
			}
		}
	}
	res.OkWithData(resList, c)
}
