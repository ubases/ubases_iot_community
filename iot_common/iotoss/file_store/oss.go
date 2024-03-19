package file_store

import (
	"fmt"
	"log"
	"strings"

	"github.com/qiniu/go-sdk/v7/storage"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type ALiYunOSS struct {
	Client     interface{}
	BucketName string
	Endpoint   string
}

//Setup 装载
//endpoint sss
func (e *ALiYunOSS) Setup(endpoint, accessKeyID, accessKeySecret, BucketName string, options ...ClientOption) error {
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		log.Println("Error:", err)
		return err
	}
	e.Client = client
	e.BucketName = BucketName
	e.Endpoint = endpoint
	return nil
}

// UpLoad 文件上传
func (e *ALiYunOSS) UpLoad(yourObjectName string, localFile interface{}) error {
	// 获取存储空间。
	bucket, err := e.Client.(*oss.Client).Bucket(e.BucketName)
	if err != nil {
		log.Println("Error:", err)
		return err
	}
	// 设置分片大小为100 KB，指定分片上传并发数为3，并开启断点续传上传。
	// 其中<yourObjectName>与objectKey是同一概念，表示断点续传上传文件到OSS时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	// "LocalFile"为filePath，100*1024为partSize。
	err = bucket.UploadFile(yourObjectName, localFile.(string), 100*1024, oss.Routines(3), oss.Checkpoint(true, ""))
	if err != nil {
		log.Println("Error:", err)
		return err
	}
	return nil
}

func (e *ALiYunOSS) GetTempToken() (string, error) {
	return "", nil
}

func (e *ALiYunOSS) GetUrl(key string) string {

	return ""
}

func (e *ALiYunOSS) GetPublicUrl(key string) string {
	domain := fmt.Sprintf("%s.%s", e.BucketName, e.Endpoint)
	if strings.Index(e.Endpoint, "http://") == -1 && strings.Index(e.Endpoint, "https://") == -1 {
		domain = fmt.Sprintf("https://%s.%s", e.BucketName, e.Endpoint)
	}
	publicAccessURL := storage.MakePublicURL(domain, key)
	return publicAccessURL
}
