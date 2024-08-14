// Code generated by sgen,2024-06-07 15:19:18. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package convert

import (
	"cloud_platform/iot_model/db_system/model"
	proto "cloud_platform/iot_proto/protos/protosService"                          
	"google.golang.org/protobuf/types/known/timestamppb" 
)

func SysAttachment_pb2db(src *proto.SysAttachment) *model.TSysAttachment {
	if src == nil {
		return nil
	}
	dbObj := model.TSysAttachment{   
         Id:src.Id,  
         FileName:src.FileName,  
         FileSize:src.FileSize,  
         FileType:src.FileType,  
         FileUrl:src.FileUrl,  
         FileMd5:src.FileMd5,  
         OssPlatform:src.OssPlatform,  
         SourceTable:src.SourceTable,  
         SourceRowId:src.SourceRowId,  
         Status:src.Status,  
         AllowClear:src.AllowClear,  
	}
	return &dbObj
}

func SysAttachment_db2pb(src *model.TSysAttachment) *proto.SysAttachment {
	if src == nil {
		return nil
	}
	pbObj := proto.SysAttachment{   
         Id:src.Id,  
         FileName:src.FileName,  
         FileSize:src.FileSize,  
         FileType:src.FileType,  
         FileUrl:src.FileUrl,  
         FileMd5:src.FileMd5,  
         OssPlatform:src.OssPlatform,  
         SourceTable:src.SourceTable,  
         SourceRowId:src.SourceRowId,  
         Status:src.Status,  
         AllowClear:src.AllowClear,  
         CreatedBy:src.CreatedBy,  
         CreatedAt:timestamppb.New(src.CreatedAt), 
	}
	return &pbObj
}