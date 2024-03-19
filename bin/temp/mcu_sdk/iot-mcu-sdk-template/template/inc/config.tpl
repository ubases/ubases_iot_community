{{define "config.h"}}
/**
 * @file    config.h
 * @brief   
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-22
 * 
 * @copyright Copyright (c) 2022  Personal
 * 
 * @par 修改日志
 * Date      |Version|Author          |Description
 * ----------|-------|----------------|--------------
 * 2022-06-22|1.0.0  |Ai-Thinker      |创建
 */
#ifndef __CONFIG_H
#define __CONFIG_H
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
// #include <stdint.h>
// #include <stdbool.h>

#include "mcu_api.h"
#include "system.h"
#include "protocol.h"




/* 开发平台创建产品后生成的产品ID，用户可根据创建不同产品而获得不同的产品ID，可在该处进行修改*/
#define  PRODUCT_ID                         "{{.ProductKey}}"

/* 产品标识,该标识是用户在创建不同产品时将自定义的产品标识 */
#define  PRODUCT_FLAG                       "{{.WifiFlag}}"


//============================ 空净物模型数据定义 ==================================================

{{range $index, $model := .Models}}
//{{$model.Name}}({{$model.RwFlagDesc}})
//备注:
#define CMD_{{$model.Identifier}}              {{$model.Dpid}}{{end}}

#endif
{{end}}
