{{define "protocol.c"}}
/**
 * @file    protocol.c
 * @brief   
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-28
 * 
 * @copyright Copyright (c) 2022  Personal
 * 
 * @par 修改日志
 * Date      |Version|Author          |Description
 * ----------|-------|----------------|--------------
 * 2022-06-28|1.0.0  |Ai-Thinker      |创建
 */
#include "config.h"

{{range $index, $model := .Models}}{{if ne $model.Type "String"}}
static {{$model.VarType}} {{$model.IdenLowCase}}_status = {{$model.DefaultValue}};{{end}}{{end}}


const McuCmdInfo CmdInfoList[] = 
{
  {{- range $index, $model := .Models}}
  {CMD_{{$model.Identifier}}, {{$model.DataType}}},{{end}}
};




/*******************************************************************************************************
 *                                  1. 以下是由爱星云平台根据物模型选择自动生成具体用户函数
 *                                ** a. 以下是由云平台自动生成APP下发数据处理函数 **
 *                                ** b. 以下是由爱星云自动生成数据上报处理函数 **
 *                                ** c. ReportAllDeviceState这个函数不可修改，由SDK自动调用上报所有状态
 * ******************************************************************************************************/

/**
 * @brief   MCU向模组上报所有的状态数据
 * 
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-29
 */
void ReportAllDeviceState(void)
{
    //#error "请在此处理可下发可上报数据及只上报数据示例,处理完成后删除该行"
    {{range $index, $model := .Models}}{{if ne $model.Type "String"}}
    //UartReport{{$model.Type}}TypeData(CMD_{{$model.Identifier}}, {{$model.IdenLowCase}}_status);{{end}}{{end}}
}

{{range $index, $model := .Models -}}
{{if and ($model.IsControl) (eq $model.Type "Bool")}}
/**
 * @brief   {{$model.Name}}状态
 * 
 * @param[in]   value      {{$model.Name}}状态数据
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-29
 */
static void Control{{$model.CaseCamel}}State(unsigned char *value)
{
    {{$model.IdenLowCase}}_status = ProGetdpidValue{{$model.Type}}(value);
    if({{$model.IdenLowCase}}_status)    
    {
        // {{$model.Name}}开状态
    }
    else
    {
        // {{$model.Name}}关闭状态
    }

    UartReport{{$model.Type}}TypeData(CMD_{{$model.Identifier}}, {{$model.IdenLowCase}}_status);
}
{{else if and ($model.IsControl) (eq $model.Type "Enum")}}
/**
 * @brief   {{$model.Name}}状态
 * 
 * @param[in]   value      {{$model.Name}}状态数据
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-29
 */
static void Control{{$model.CaseCamel}}State(unsigned char *value)
{
    {{$model.IdenLowCase}}_status = ProGetdpidValue{{$model.Type}}(value);
    switch ({{$model.IdenLowCase}}_status)
    {
        {{range $idx, $value := $model.Values}}
        case {{$value}}:
        {
        }
        break;
        {{end}}
        default:
            break;
    }

    UartReport{{$model.Type}}TypeData(CMD_{{$model.Identifier}}, {{$model.IdenLowCase}}_status);
}
{{else if and ($model.IsControl) (eq $model.Type "String")}}
/**
 * @brief   {{$model.Name}}状态
 * 
 * @param[in]   value      {{$model.Name}}状态数据
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-29
 */
static void Control{{$model.CaseCamel}}State(unsigned char *value, unsigned short DataLen)
{
    UartReport{{$model.Type}}TypeData(CMD_{{$model.Identifier}}, value, DataLen);
}
{{else if and ($model.IsControl) (eq $model.Type "Value")}}
/**
 * @brief   {{$model.Name}}状态
 * 
 * @param[in]   value      {{$model.Name}}状态数据
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-29
 */
static void Control{{$model.CaseCamel}}State(unsigned char *value)
{
    {{$model.IdenLowCase}}_status = ProGetdpidValueInt(value);
    UartReport{{$model.Type}}TypeData(CMD_{{$model.Identifier}}, {{$model.IdenLowCase}}_status);
}
{{end}}
{{- end}}

/*************************************************************************************************************
 *                              3. 以下函数主要是处理APP下发的数据
 *                      ** a. 该函数名不可修改，云平台根据不同物模型而进行自动生成处理事件 **
 *                      ** b. 云平台可根据不同物模型事件而进行自动补充调用自动生成的具体用户函数 **
 * ***********************************************************************************************************/
/**
 * @brief   处理待控制的物模型数据
 * 
 * @param[in]   cmd_id   数据id
 * @param[in]   FrameData     待控制的状态数据
 * @param[in]   data_len 数据长度
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-29
 */
void ProcessDataCmd(unsigned char cmd_id, unsigned char *FrameData, unsigned short DataLen)
{
    switch (cmd_id)
    {
        {{range $index, $model := .Models -}}
        {{if ($model.IsControl) }}
        case CMD_{{$model.Identifier}}:
        {
            // {{$model.Name}}处理函数
            {{if ne $model.Type "String"}}
            Control{{$model.CaseCamel}}State(FrameData);
            {{- else -}}
            Control{{$model.CaseCamel}}State(FrameData, DataLen);
            {{- end}}
        }
        break;
        {{end}}    
        {{- end}}
        default:
            break;
    }
}









/***********************************************************************************************************
 *                                      4. 以下代码为SDK内部调用代码
 *                              ** a. 以下代码由MCU SDK调用  **
 *                              ** b. 以下代码函数名不可修改，亦不可删除 **
 *                              ** c. 以下代码若是MCU开发者需调用，可自行补充内容，若不需要调用，可无需改动 **
 * *********************************************************************************************************/


/**
 * @brief   串口发送数据函数
 * 
 * @note    该函数需由客户适配好对于发送接口。即MCU_SDK调用该函数，即可将数据通过串口发送给模组
 * 
 * @param[in]   dataBuff        待发送的数据
 * @param[in]   dataLen         数据长度
 * @return  unsigned char            返回发送结果
 *          0：发送成功
 *          -1：发送失败
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-23
 */
char UartWriteData(unsigned char *dataBuff, unsigned char dataLen)
{
	//#error "请将MCU串口发送函数填入该函数,并删除该行"	
	return 0;
}


/**
 * @brief   根据数据Id获取序号
 * 
 * @param[in]   cmd_id   数据id
 * @return  unsigned char      返回序号
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-07-13
 */
unsigned char GetCmdIdIndex(unsigned char cmd_id)
{
    unsigned char total = (sizeof(CmdInfoList) / sizeof(CmdInfoList[0]));
    unsigned char index = 0;

    for(index = 0; index < total; index++)
    {
        if(CmdInfoList[index].cmd_id == cmd_id)
        {
            break;
        }
    }

    return index;
}

/**
 * @brief   通知MCU设备有新OTA固件可更新，可接收到可升级的版本号以及OTA固件的md5校验数据
 * 
 * @param[in]   data            版本好以及md5数据字符串
 * @param[in]   dataLen         数据长度
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2023-04-04
 */
void UpDateFirmStart(unsigned char *data, unsigned short dataLen)
{
}

/**
 * @brief    接收处理OTA固件数据，接收到前4个字节数据是OTA固件包的长度，其余字节均是OTA固件包数据
 * 
 * @param[in]   data       数据长度以及OTA固件数据
 * @param[in]   dataLen    数据长度
 *  
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2023-04-04
 */
void UpDataFireDownload(unsigned char *data, unsigned short dataLen)
{
}
{{end}}
