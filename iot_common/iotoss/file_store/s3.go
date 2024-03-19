package file_store

import (
	"bytes"
	"cloud_platform/iot_common/iotutil"
	"fmt"
	"mime"
	"net/http"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go/aws/endpoints"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var onceApk sync.Once // guards initMime

func initMime() {
	mime.AddExtensionType(".apk", "applicaiton/vnd.Android.package-archive")
}

type S3OBS struct {
	Client     interface{}
	BucketName string
	Endpoint   string
	Region     string
}

func (e *S3OBS) Setup(endpoint, accessKeyID, accessKeySecret, BucketName string, options ...ClientOption) error {
	region := endpoints.CnNorth1RegionID
	if len(options) > 0 {
		region = iotutil.ToString(options[0]["region"])
	}
	// 创建ObsClient结构体
	sess := session.Must(session.NewSession(&aws.Config{
		//Region:      aws.String(endpoints.UsEast1RegionID),
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKeyID, accessKeySecret, ""),
	}))
	service := s3.New(sess)
	e.Client = service
	e.BucketName = BucketName
	e.Endpoint = endpoint
	return nil
}

// UpLoad 文件上传
// yourObjectName 文件路径名称，与objectKey是同一概念，表示断点续传上传文件到OSS时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg
func (e *S3OBS) UpLoad(yourObjectName string, localFile interface{}) error {

	onceApk.Do(initMime)

	fp, err := os.Open(localFile.(string))
	defer fp.Close()

	//读取文件大小
	var fileSize int64 = 0
	if err == nil {
		fi, _ := fp.Stat()
		fileSize = fi.Size()
	}
	buffer := make([]byte, fileSize)
	fp.Read(buffer)

	//获取文件类型
	//contentType, _ := GetFileContentType(fp)
	fileType := mime.TypeByExtension(path.Ext(localFile.(string)))
	if fileType == "" {
		fileType = http.DetectContentType(buffer)
	}

	service := e.Client.(*s3.S3)
	obj := &s3.PutObjectInput{
		Bucket:               aws.String(e.BucketName), // bucket名称，把自己创建的bucket名称替换到此处即可
		Key:                  aws.String(yourObjectName),
		ACL:                  aws.String("public-read"), // could be private if you want it to be access by only authorized users
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(fileSize),
		ContentType:          aws.String(fileType), //aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
		StorageClass:         aws.String("INTELLIGENT_TIERING"),
	}
	_, err = service.PutObject(obj)

	//需要授权访问
	//_, err = service.PutObjectWithContext(ctx, &s3.PutObjectInput{
	//	Bucket: aws.String(e.BucketName), //"iot-aithings-public"
	//	Key:    aws.String(yourObjectName),
	//	Body:   fp,
	//})
	if err == nil {
		fmt.Printf("RequestId:%s\n", yourObjectName)
	} else {
		fmt.Println(err.Error())
	}
	return nil
}

func (e *S3OBS) GetTempToken() (string, error) {
	return "", nil
}

func (e *S3OBS) GetUrl(key string) string {

	return ""
}

func (e *S3OBS) GetPublicUrl(key string) string {
	newKey := strings.TrimRight(key, "/")
	publicAccessURL := ""
	if strings.Index(e.Endpoint, "http://") == -1 && strings.Index(e.Endpoint, "https://") == -1 {
		publicAccessURL = fmt.Sprintf("https://%s/%s", e.Endpoint, newKey)
	} else {
		publicAccessURL = fmt.Sprintf("%s/%s", e.Endpoint, newKey)
	}
	return publicAccessURL
}
