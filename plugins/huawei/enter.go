package huawei

// 引入依赖包
import (
	"bytes"
	"errors"
	"fmt"
	obs "github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"gvb_server/config"
	"gvb_server/global"
	"time"
)

func getObsClient(q config.HuaWei) *obs.ObsClient {
	ak := global.Config.HuaWei.AccessKey
	sk := global.Config.HuaWei.SecretKey
	endPoint := global.Config.HuaWei.EndPoint
	obsClient, err := obs.New(ak, sk, endPoint, obs.WithSignature(obs.SignatureObs) /*, obs.WithSecurityToken(securityToken)*/)
	if err != nil {
		fmt.Printf("Create obsClient error, errMsg: %s", err.Error())
	}
	return obsClient

}

// UploadImage 上传图片  文件数组，前缀
func UploadImage(data []byte, imageName string, prefix string) (filePath string, err error) {
	if !global.Config.HuaWei.Enable {
		return "", errors.New("请启用华为云上传")
	}
	q := global.Config.HuaWei
	if q.AccessKey == "" || q.SecretKey == "" {
		return "", errors.New("请配置accessKey及secretKey")
	}
	if float64(len(data))/1024/1024 > q.Size {
		return "", errors.New("文件超过设定大小")
	}
	obsClient := getObsClient(q)
	input := &obs.PutObjectInput{}
	// 指定存储桶名称
	input.Bucket = "tblog"
	// 指定上传对象，此处以 example/objectname 为例。
	now := time.Now().Format("20060102150405")
	//key := fmt.Sprintf("%s/%s__%s", prefix, now, imageName)
	input.Key = fmt.Sprintf("%s/%s_%s", prefix, now, imageName)
	//fd, _ := os.Open("localfile")
	input.Body = bytes.NewReader(data)
	// 流式上传本地文件
	output, err := obsClient.PutObject(input)

	return output.ObjectUrl, err
	//if err == nil {
	//	fmt.Printf("Put object(%s) under the bucket(%s) successful!\n", input.Key, input.Bucket)
	//	fmt.Printf("StorageClass:%s, ETag:%s\n",
	//		output.StorageClass, output.ETag)
	//	return
	//}
	//fmt.Printf("Put object(%s) under the bucket(%s) fail!\n", input.Key, input.Bucket)
	//if obsError, ok := err.(obs.ObsError); ok {
	//	fmt.Println("An ObsError was found, which means your request sent to OBS was rejected with an error response.")
	//	fmt.Println(obsError.Error())
	//} else {
	//	fmt.Println("An Exception was found, which means the client encountered an internal problem when attempting to communicate with OBS, for example, the client was unable to access the network.")
	//	fmt.Println(err)
	//}
	//upToken := getToken(q)
	//cfg := getCfg(q)
	//
	//formUploader := storage.NewFormUploader(&cfg)
	//ret := storage.PutRet{}
	//putExtra := storage.PutExtra{
	//	Params: map[string]string{},
	//}
	//dataLen := int64(len(data))
	//return nil, nil
}
