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


static BOOL switch_status = FALSE;
static unsigned char mode_status = 0;
static unsigned char speed_status = 0;
static unsigned long  temp_num = 0;
static unsigned long  humidity_num = 0;
static unsigned long  pm25_num = 0;
static unsigned long tvoc_num = 0;
static unsigned long eco2_num = 0;
static unsigned char airquality_status = 0;
static unsigned long  filter_num = 0;
static unsigned long  filterdays_num =0;
static BOOL filter_reset_status =FALSE;
static unsigned char faultCode_status = 0;
static unsigned char faultType_status = 0;
static unsigned long  totalTime_num = 0;
static unsigned long  totalPm_num = 0;
static unsigned char tmep_unit_convert = 0;
static unsigned long  methanal_num = 0;
static BOOL anion_status = FALSE;
static BOOL uv_status = FALSE;
static BOOL  lock_status = FALSE;
static unsigned char light_status = 0;




const McuCmdInfo CmdInfoList[] = 
{
  {CMD_LIGHT, DATA_TYPE_ENUM},
  {CMD_LOCK, DATA_TYPE_BOOL},
  {CMD_UV, DATA_TYPE_BOOL},
  {CMD_ANION, DATA_TYPE_BOOL},
  {CMD_METHANAL, DATA_TYPE_VALUE},
  {CMD_TEMP_UNIT_CONVERT, DATA_TYPE_ENUM},
  {CMD_TOTAL_PM, DATA_TYPE_VALUE},
  {CMD_TOTAL_TIME, DATA_TYPE_VALUE},
  {CMD_FAULT_TYPE, DATA_TYPE_ENUM},
  {CMD_FAULT_CODE, DATA_TYPE_ENUM},
  {CMD_FILTER_RESET, DATA_TYPE_BOOL},
  {CMD_FILTER_DAYS, DATA_TYPE_VALUE},
  {CMD_FILTER, DATA_TYPE_VALUE},
  {CMD_AIR_QUALITY, DATA_TYPE_ENUM},
  {CMD_ECO2, DATA_TYPE_VALUE},
  {CMD_TVOC, DATA_TYPE_VALUE},
  {CMD_PM25, DATA_TYPE_VALUE},
  {CMD_HUMIDITY, DATA_TYPE_VALUE},
  {CMD_TEMP, DATA_TYPE_VALUE},
  {CMD_SPEED, DATA_TYPE_ENUM},
  {CMD_MODE, DATA_TYPE_ENUM},
  {CMD_SWITCH, DATA_TYPE_BOOL},
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
    
    UartReportEnumTypeData(CMD_LIGHT, light_status);
    UartReportBoolTypeData(CMD_LOCK, lock_status);
    UartReportBoolTypeData(CMD_UV, uv_status);
    UartReportBoolTypeData(CMD_ANION, anion_status);
    UartReportValueTypeData(CMD_METHANAL, methanal_num);
    UartReportEnumTypeData(CMD_TEMP_UNIT_CONVERT, tmep_unit_convert);
    UartReportValueTypeData(CMD_TOTAL_PM, totalPm_num);
    UartReportValueTypeData(CMD_TOTAL_TIME, totalTime_num);
    UartReportEnumTypeData(CMD_FAULT_TYPE, faultType_status);
    UartReportEnumTypeData(CMD_FAULT_CODE, faultCode_status);
    UartReportBoolTypeData(CMD_FILTER_RESET, filter_reset_status);
    UartReportValueTypeData(CMD_FILTER_DAYS, filterdays_num);
    UartReportValueTypeData(CMD_FILTER, filter_num);
    UartReportEnumTypeData(CMD_AIR_QUALITY, airquality_status);
    UartReportValueTypeData(CMD_ECO2, eco2_num);
    UartReportValueTypeData(CMD_TVOC, tvoc_num);
    UartReportValueTypeData(CMD_PM25, pm25_num);
    UartReportValueTypeData(CMD_HUMIDITY, humidity_num);
    UartReportValueTypeData(CMD_TEMP, temp_num);
    UartReportEnumTypeData(CMD_SPEED, speed_status);
    UartReportEnumTypeData(CMD_MODE, mode_status);
    UartReportBoolTypeData(CMD_SWITCH, switch_status);
}



/**
 * @brief   提示灯状态
 * 
 * @param[in]   value      提示灯状态数据
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-29
 */
static void ControlLightState(unsigned char *value)
{
    light_status = ProGetdpidValueEnum(value);
    switch (light_status)
    {
        
        case 0:
        {
        }
        break;
        
        case 1:
        {
        }
        break;
        
        case 2:
        {
        }
        break;
        
        default:
            break;
    }

    UartReportEnumTypeData(CMD_LIGHT, light_status);
}



/**
 * @brief   童锁状态
 * 
 * @param[in]   value      童锁状态数据
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-29
 */
static void ControlLockState(unsigned char *value)
{
    lock_status = ProGetdpidValueBool(value);
    if(lock_status)
    {
        // 童锁开状态
    }
    else
    {
        // 童锁关闭状态
    }

    UartReportBoolTypeData(CMD_LOCK, lock_status);
}



/**
 * @brief   UV杀菌状态
 * 
 * @param[in]   value      UV杀菌状态数据
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-29
 */
static void ControlUvState(unsigned char *value)
{
    uv_status = ProGetdpidValueBool(value);
    if(uv_status)
    {
        // UV杀菌开状态
    }
    else
    {
        // UV杀菌关闭状态
    }

    UartReportBoolTypeData(CMD_UV, uv_status);
}



/**
 * @brief   负离子状态
 * 
 * @param[in]   value      负离子状态数据
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-29
 */
static void ControlAnionState(unsigned char *value)
{
    anion_status = ProGetdpidValueBool(value);
    if(anion_status)
    {
        // 负离子开状态
    }
    else
    {
        // 负离子关闭状态
    }

    UartReportBoolTypeData(CMD_ANION, anion_status);
}





/**
 * @brief   温标切换状态
 * 
 * @param[in]   value      温标切换状态数据
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-29
 */
static void ControlTempUnitConvertState(unsigned char *value)
{
    tmep_unit_convert = ProGetdpidValueEnum(value);
    switch (tmep_unit_convert)
    {
        
        case 0:
        {
        }
        break;
        
        case 1:
        {
        }
        break;
        
        default:
            break;
    }

    UartReportEnumTypeData(CMD_TEMP_UNIT_CONVERT, tmep_unit_convert);
}







/**
 * @brief   故障类型状态
 * 
 * @param[in]   value      故障类型状态数据
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-29
 */
static void ControlFaultTypeState(unsigned char *value)
{
    tmep_unit_convert = ProGetdpidValueEnum(value);
    switch (tmep_unit_convert)
    {
        
        case 0:
        {
        }
        break;
        
        case 1:
        {
        }
        break;
        
        default:
            break;
    }

    UartReportEnumTypeData(CMD_FAULT_TYPE, faultType_status);
}



/**
 * @brief   故障告警状态
 * 
 * @param[in]   value      故障告警状态数据
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-29
 */
static void ControlFaultCodeState(unsigned char *value)
{
    faultCode_status = ProGetdpidValueEnum(value);
    switch (faultCode_status)
    {
        
        default:
            break;
    }

    UartReportEnumTypeData(CMD_FAULT_CODE, faultCode_status);
}



/**
 * @brief   滤芯复位状态
 * 
 * @param[in]   value      滤芯复位状态数据
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-29
 */
static void ControlFilterResetState(unsigned char *value)
{
    filter_reset_status = ProGetdpidValueBool(value);
    if(filter_reset_status)
    {
        // 滤芯复位开状态
    }
    else
    {
        // 滤芯复位关闭状态
    }

    UartReportBoolTypeData(CMD_FILTER_RESET, filter_reset_status);
}







/**
 * @brief   空气质量状态
 * 
 * @param[in]   value      空气质量状态数据
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-29
 */
static void ControlAirQualityState(unsigned char *value)
{
    airquality_status = ProGetdpidValueEnum(value);
    switch (airquality_status)
    {
        
        case 0:
        {
        }
        break;
        
        case 1:
        {
        }
        break;
        
        case 2:
        {
        }
        break;
        
        case 3:
        {
        }
        break;
        
        case 4:
        {
        }
        break;
        
        default:
            break;
    }

    UartReportEnumTypeData(CMD_AIR_QUALITY, airquality_status);
}













/**
 * @brief   风速状态
 * 
 * @param[in]   value      风速状态数据
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-29
 */
static void ControlSpeedState(unsigned char *value)
{
    speed_status = ProGetdpidValueEnum(value);
    switch (speed_status)
    {
        case 0:
        {
        }
        break;
        
        case 1:
        {
        }
        break;
        
        case 2:
        {
        }
        break;
        
        case 3:
        {
        }
        break;
        
        case 4:
        {
        }
        break;
        
        case 5:
        {
        }
        break;
        
        default:
            break;
    }

    UartReportEnumTypeData(CMD_SPEED, speed_status);
}



/**
 * @brief   模式状态
 * 
 * @param[in]   value      模式状态数据
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-29
 */
static void ControlModeState(unsigned char *value)
{
    mode_status = ProGetdpidValueEnum(value);
    switch (mode_status)
    {
        case 0:
        {
        }
        break;
        
        case 1:
        {
        }
        break;
        
        case 2:
        {
        }
        break;
        
        case 3:
        {
        }
        break;
        
        default:
            break;
    }
    UartReportEnumTypeData(CMD_MODE, mode_status);
}



/**
 * @brief   开关状态
 * 
 * @param[in]   value      开关状态数据
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-29
 */
static void ControlSwitchState(unsigned char *value)
{
    switch_status = ProGetdpidValueBool(value);
    if(switch_status)
    {
        // 开关开状态
    }
    else
    {
        // 开关关闭状态
    }

    UartReportBoolTypeData(CMD_SWITCH, switch_status);
}



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
void ProcessDataCmd(unsigned char cmd_id, unsigned char *FrameData)
{
    switch (cmd_id)
    {
        
        case CMD_LIGHT:
        {
            // 提示灯处理函数
            ControlLightState(FrameData);
        }
        break;        
        
        case CMD_LOCK:
        {
            // 童锁处理函数
            ControlLockState(FrameData);
        }
        break;        
        
        case CMD_UV:
        {
            // UV杀菌处理函数
            ControlUvState(FrameData);
        }
        break;        
        
        case CMD_ANION:
        {
            // 负离子处理函数
            ControlAnionState(FrameData);
        }
        break;        
             
        case CMD_TEMP_UNIT_CONVERT:
        {
            // 温标切换处理函数
            ControlTempUnitConvertState(FrameData);
        }
        break;                  
        
        case CMD_FAULT_TYPE:
        {
            // 故障类型处理函数
            ControlFaultTypeState(FrameData);
        }
        break;        
        
        case CMD_FAULT_CODE:
        {
            // 故障告警处理函数
            ControlFaultCodeState(FrameData);
        }
        break;        
        
        case CMD_FILTER_RESET:
        {
            // 滤芯复位处理函数
            ControlFilterResetState(FrameData);
        }
        break;               
        
        case CMD_AIR_QUALITY:
        {
            // 空气质量处理函数
            ControlAirQualityState(FrameData);
        }
        break;                              
        
        case CMD_SPEED:
        {
            // 风速处理函数
            ControlSpeedState(FrameData);
        }
        break;        
        
        case CMD_MODE:
        {
            // 模式处理函数
            ControlModeState(FrameData);
        }
        break;        
        
        case CMD_SWITCH:
        {
            // 开关处理函数
            ControlSwitchState(FrameData);
        }
        break;        

        
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
//	#error "请将MCU串口发送函数填入该函数,并删除该行"	
//	USART2SendDatas(dataBuff, dataLen);
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
