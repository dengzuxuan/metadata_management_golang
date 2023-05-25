package utils

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"mime/multipart"
)

func handleError(err error) {
	fmt.Println("Error:", err)
}
func UploadOss(objectname string, file *multipart.FileHeader) error {
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	endpoint := "https://oss-cn-beijing.aliyuncs.com"
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。\

	accessKeyId := "LTAI5tCq3h3XaPtSV4nTTAkL"
	accessKeySecret := "8s4Za6MP73MuOLcdQOgsB0zqcMWrj8"
	// yourBucketName填写存储空间名称。
	bucketName := "metadata-manager"
	//// yourObjectName填写Object完整路径，完整路径不包含Bucket名称。
	//objectName := objectname
	//// yourLocalFileName填写本地文件的完整路径。
	//localFileName := "yourLocalFileName"
	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		return err
	}
	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return err
	}
	fd, _ := file.Open()
	// 上传文件。
	err = bucket.PutObject(objectname, fd)
	if err != nil {
		return err
	}
	return nil
}
