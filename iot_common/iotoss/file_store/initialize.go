package file_store

import "fmt"

type OXS struct {
	// Endpoint 访问域名
	Endpoint string
	// AccessKeyID AK
	AccessKeyID string
	// AccessKeySecret AKS
	AccessKeySecret string
	// BucketName 桶名称
	BucketName string
	// 地区
	Region string
}

// Setup 配置文件存储driver
func (e *OXS) Setup(driver DriverType, options ...ClientOption) FileStoreType {
	fileStoreType := driver
	var fileStore FileStoreType
	switch fileStoreType {
	case AliYunOSS:
		fileStore = new(ALiYunOSS)
		err := fileStore.Setup(e.Endpoint, e.AccessKeyID, e.AccessKeySecret, e.BucketName)
		if err != nil {
			fmt.Println(err)
		}
		return fileStore
	case HuaweiOBS:
		fileStore = new(HuaWeiOBS)
		err := fileStore.Setup(e.Endpoint, e.AccessKeyID, e.AccessKeySecret, e.BucketName)
		if err != nil {
			fmt.Println(err)
		}
		return fileStore
	case QiNiuKodo:
		fileStore = new(QiNiuKODO)
		err := fileStore.Setup(e.Endpoint, e.AccessKeyID, e.AccessKeySecret, e.BucketName, ClientOption{
			"Expires": uint64(3600 * 1000000),
			"ApiHost": e.Endpoint,
		})
		if err != nil {
			fmt.Println(err)
		}
		return fileStore
	case AwsS3Kodo:
		fileStore = new(S3OBS)
		err := fileStore.Setup(e.Endpoint, e.AccessKeyID, e.AccessKeySecret, e.BucketName,
			map[string]interface{}{"region": e.Region})
		if err != nil {
			fmt.Println(err)
		}
		return fileStore
	}

	return nil
}
