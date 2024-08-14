package commonGlobal

import (
	"cloud_platform/iot_app_api_service/config"
	"cloud_platform/iot_app_api_service/rpc"
	"cloud_platform/iot_common/iotstruct"
	"cloud_platform/iot_common/iotutil"
	proto "cloud_platform/iot_proto/protos/protosService"
	"context"
)

func SaveAttachmentRecord(dir string, files ...iotstruct.FileResponse) []iotstruct.FileResponse {
	if files == nil {
		return files
	}
	go func() {
		defer iotutil.PanicHandler("saveAttachmentRecord", files)
		for _, f := range files {
			rpc.ClientSysAttachmentService.Create(context.Background(), &proto.SysAttachment{
				Id:          iotutil.GetNextSeqInt64(),
				FileName:    f.Name,
				FileSize:    int32(f.Size),
				FileType:    f.Type,
				FileUrl:     f.FullPath,
				FileMd5:     f.Key,
				OssPlatform: config.Global.Oss.UseOss,
				SourceTable: dir,
				SourceRowId: "",
				Status:      1, //默认为1
				AllowClear:  1,
			})
		}
	}()
	return files
}

func SetAttachmentStatus(sourceTable, sourceRowId string, filepath ...string) {
	if filepath == nil && len(filepath) == 0 {
		return
	}
	go func() {
		defer iotutil.PanicHandler("SetAttachmentStatus", filepath)
		for _, f := range filepath {
			if f == "" {
				continue
			}
			rpc.ClientSysAttachmentService.Update(context.Background(), &proto.SysAttachment{
				FileUrl:     f,
				Status:      2, //更新上传成功状态
				SourceTable: sourceTable,
				SourceRowId: sourceRowId,
			})
		}
	}()
}
