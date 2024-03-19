/**
 * @file    mcu_api.c
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
#include "config.h"
#include <stdarg.h>
#include <stdio.h>


/* 串口接收数据缓存 */
static unsigned char revBuff[MCU_UART_BUFF_MAX_LEN] = {0};
static unsigned char processBuff[MCU_UART_BUFF_MAX_LEN] = {0};
static unsigned char  WifiWorkState = WIFI_SATE_UNKNOW;
static unsigned char  WifiResetFlag  = RESET_WIFI_INVAILD;
static unsigned char  *revBuffPt_in = NULL;
static unsigned char  *revBuffPt_out = NULL;




/**
* @brief   MCU向模组查询工作模式
* 
* 
* @author  Ai-Thinker (zhuolm@tech-now.com)
* @date    2022-06-29
*/
void UartReqModuleWorkMode(void)
{
   UartProFrameSend(WORK_MODE_CMD, NULL, 0);
}



/**
 * @brief   获取wifi工作状态
 * 
 * @return  unsigned char  返回工作状态
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-23
 */
unsigned char GetWifiWorkState(void)
{
    return WifiWorkState;
}


/**
 * @brief   设置WiFi的工作状态
 * 
 * @param[in]   state  设置WiFi的工作状态
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-23
 */
void SetWifiWorkState(unsigned char state)
{
    WifiWorkState = state;
    switch (WifiWorkState)
    {
        case BLE_CONFIG_STATE:
        {
            InfoLog("BLE_CONFIG_STATE \r\n");
        }
        break;

        case AP_CONFIG_STATE:
        {
            InfoLog("AP_CONFIG_STATE \r\n");
        }
        break;

        case WIFI_NOT_CONNECTED:
        {
            InfoLog("WIFI_NOT_CONNECTED \r\n");
        }
        break;

        case WIFI_CONNECTED:
        {
            InfoLog("WIFI_CONNECTED \r\n");
        }
        break;

        case WIFI_CONN_CLOUD:
        {
            InfoLog("WIFI_CONN_CLOUD \r\n");
        }
        break;

        case WIFI_LOW_POWER:
        {
            InfoLog("WIFI_LOW_POWER \r\n");
        }
        break;

        case BLE_AND_AP_STATE:
        {
            InfoLog("BLE_AND_AP_STATE \r\n");
        }
        break;
        
        default:
            break;
    }
}

/**
 * @brief   获取重置WiFi的情况
 * 
 * @return  unsigned char 返回重置WiFi的情况
 *          RESET_WIFI_INVAILD：重置无效
 *          RESET_WIFI_SUCCESS：已重置成功
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-23
 */
unsigned char GetResetState(void)
{
    return WifiResetFlag;
}


/**
 * @brief   设置WiFi重置状态
 * 
 * @param[in]   state 
 *              RESET_WIFI_INVAILD：设置无效状态
 *              RESET_WIFI_SUCCESS：设置已重置成功
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-23
 */
void SetResetState(unsigned char state)
{
    WifiResetFlag = state;
}


/**
 * @brief   MCU请求模组连接指定路由
 * 
 * @param[in]   ssid        路由的ssid
 * @param[in]   passwd      路由的passwd
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-28
 */
void UartReqConnectWifi(char *ssid, char *passwd)
{
    unsigned char buff[100] = {0};

    sprintf((char *)buff, "{\"%s\":\"%s\"}", "ssid", ssid);
    UartProFrameSend(WIFI_CONNECT_TEST_CMD, (unsigned char *)buff, strlen((char *)buff));
    memset(buff,0, sizeof(buff));
    sprintf((char *)buff, "{\"%s\":\"%s\"}","password", passwd);
    UartProFrameSend(WIFI_CONNECT_TEST_CMD, (unsigned char *)buff, strlen((char *)buff));
    memset(buff,0, sizeof(buff));
}

/**
 * @brief   MCU向模组发送重置网络状态
 * 
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-28
 */
void UartReportResetState(void)
{
    WifiResetFlag = RESET_WIFI_INVAILD;

    UartProFrameSend(WIFI_RESET_CMD, NULL, 0);
}






/**
 * @brief   MCU向模组上报bool类型的状态数据
 * 
 * @param[in]   cmd_id  数据id
 * @param[in]   value   状态数据
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-28
 */
void UartReportBoolTypeData(unsigned char cmd_id, BOOL value)
{
    unsigned char FrameData[5] = {0};
    unsigned char len = 1;

    FrameData[0] = cmd_id;
    FrameData[1] = DATA_TYPE_BOOL;
    FrameData[2] = ((len & 0xff00) >> 8);
    FrameData[3] = (len & 0xff);
    FrameData[4] = value;

    UartProFrameSend(STATE_UPLOAD_CMD, FrameData, len + 4);
}



/**
 * @brief   MCU向模组上报整数类型的状态数据
 * 
 * @param[in]   cmd_id  数据id
 * @param[in]   value   状态数据
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-28
 */
void UartReportValueTypeData(unsigned char cmd_id, unsigned long value)
{
    unsigned char data[8] = {0};
    unsigned short len = sizeof(value);

    data[0] = cmd_id;
    data[1] = DATA_TYPE_VALUE;
    data[2] = ((len & 0xff00) >> 8);
    data[3] = (len & 0xff);
    data[4] = (value >> 24);
    data[5] = (value >> 16);
    data[6] = (value >> 8);
    data[7] = (value & 0xff);

    UartProFrameSend(STATE_UPLOAD_CMD, data, len + 4);
}



/**
 * @brief   MCU向模组上报枚举类型的状态数据
 * 
 * @param[in]   cmd_id  数据id
 * @param[in]   value   状态数据
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-28
 */
void UartReportEnumTypeData(unsigned char cmd_id, unsigned char value)
{
    unsigned char FrameData[5] = {0};
    unsigned char len = 1;

    FrameData[0] = cmd_id;
    FrameData[1] = DATA_TYPE_ENUM;
    FrameData[2] = ((len & 0xff00) >> 8);
    FrameData[3] = (len & 0xff);
    FrameData[4] = value;

    UartProFrameSend(STATE_UPLOAD_CMD, FrameData, len + 4);
}


/**
 * @brief   MCU向模组上报字符串类型的状态数据
 * 
 * @param[in]   cmd_id  数据id
 * @param[in]   value   状态数据
 * @param[in]   value_len   状态数据长度
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-28
 */
void UartReportStringTypeData(unsigned char cmd_id, unsigned char *value, unsigned char value_len)
{
    unsigned char *data = NULL;
    unsigned short len = value_len;

    data = (unsigned char *)malloc(value_len + 1);
    if(data == NULL)
    {
        ErrorLog("malloc fail \r\n");
        return ;
    }

    memset(data, 0, value_len+1);
    data[0] = cmd_id;
    data[1] = DATA_TYPE_STRING;
    data[2] = ((len & 0xff00) >> 8);
    data[3] = (len & 0xff);
    memcpy((char *)&data[4], value, value_len);

    UartProFrameSend(STATE_UPLOAD_CMD, data, len + 4);
    free(data);
}



/**
 * @brief   比较OTA版本号
 * 
 * @param[in]   target_version      OTA固件版本号
 * @param[in]   source_version      模组固件版本号
 * @return  int8_t 
 *          1: OTA固件版本号大于模组固件版本号
 *          0：OTA固件版本号等于模组固件版本号
 *          -1：OTA固件版本号小于模组固件版本号
 *          -2：比较失败
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-10-24
 */
signed char AilinkCheckVersion(char *target_version, char *source_version)
{
    int order = 0;
    double res = 0;
    const char *p_v1 = target_version;
    const char *p_v2 = source_version;

    if(p_v1 == NULL || p_v2 == NULL)
    {
        ErrorLog("param err \r\n");
        return -2;
    }

    InfoLog("target_version = %s \r\n", p_v1);
    InfoLog("source_version = %s \r\n", p_v2);

    if(strlen(target_version) == 0)
    {
        InfoLog("target_version < source_version \r\n");
        return -1;
    }
    else if(strlen(source_version) == 0)
    {
        InfoLog("target_version > source_version \r\n");
        return 1;
    }

    while (*p_v1 && *p_v2) 
    {
        char buf_v1[32] = {0};
        char buf_v2[32] = {0};

        char *i_v1 = strchr(p_v1, '.');
        char *i_v2 = strchr(p_v2, '.');

        if (!i_v1 || !i_v2) 
        {
            break;
        }

        if (i_v1 != p_v1) 
        {
            strncpy(buf_v1, p_v1, i_v1 - p_v1);
            p_v1 = i_v1;
        }
        else
        {
            p_v1++;
        }

        if (i_v2 != p_v2) 
        {
            strncpy(buf_v2, p_v2, i_v2 - p_v2);
            p_v2 = i_v2;
        }
        else
        {
            p_v2++;
        }

        order = atoi(buf_v1) - atoi(buf_v2);
        if (order != 0)
        {
            InfoLog("order = %d \r\n", order);
            res = order < 0 ? 2 : 1;
            if (res == 2)
            {
                InfoLog("target_version < source_version \r\n");
                return -1;
            }
            else
            {
                InfoLog("target_version > source_version \r\n");
                return 1;
            }
        }
    }

    res = atof(p_v1) - atof(p_v2);

    if (res < 0)
    {
        InfoLog("target_version < source_version \r\n");
        return -1;
    }
    if (res > 0)
    {
        InfoLog("target_version > source_version \r\n");
        return 1;
    }

    InfoLog("target_version == source_version \r\n");
    return 0;
}


/**
 * @brief   从bool值种获取数据
 * 
 * @param[in]   value       数据内容
 * @return  true            返回true的数据内容
 * @return  false           返回flash的书内容
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-08-31
 */
BOOL ProGetdpidValueBool(unsigned char *value)
{
    if(*value)
    {
        return TRUE;
    }
    else
    {
        return FALSE;
    }
}

/**
 * @brief   获取枚举数据
 * 
 * @param[in]   value       数据内容
 * @return  unsigned char   获取的数据内容
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-08-31
 */
unsigned char ProGetdpidValueEnum(unsigned char *value)
{
    unsigned long TranslateValue = 0;

    TranslateValue = (unsigned long)value[0];
    TranslateValue <<= 8;
    TranslateValue |= (unsigned long)value[1];
    TranslateValue <<= 8;
    TranslateValue |= (unsigned long)value[2];
    TranslateValue <<= 8;
    TranslateValue |= (unsigned long)value[3];
    
    return ((unsigned char)TranslateValue);
}

/**
 * @brief   获取整形数据
 * 
 * @param[in]   value       数据内容
 * @return  unsigned long   返回数据值
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-08-31
 */
unsigned long ProGetdpidValueInt(unsigned char *value)
{
    unsigned long TranslateValue = 0;

    TranslateValue = (unsigned long)value[0];
    TranslateValue <<= 8;
    TranslateValue |= (unsigned long)value[1];
    TranslateValue <<= 8;
    TranslateValue |= (unsigned long)value[2];
    TranslateValue <<= 8;
    TranslateValue |= (unsigned long)value[3];
    
    return TranslateValue;
}


/**
 * @brief   该函数主要是逐一字节接收串口数据，若是串口采用缓存方式处理串口数据，可调用UartRevStream该函数处理。
 * 
 * @param[in]   FrameData        串口数据
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-22
 */
void UartRevOneByte(unsigned char FrameData)
{
    if((revBuffPt_in > revBuffPt_out) && ((revBuffPt_in - revBuffPt_out) >= sizeof(revBuff)))
    {
        ErrorLog("buff is full \r\n");
    }
    else
    {
        if(revBuffPt_in >= (unsigned char *)(revBuff + sizeof(revBuff)))
        {
            revBuffPt_in = (unsigned char *)(revBuff);
        }

        *revBuffPt_in++ = FrameData;
    }
}


/**
 * @brief   接收串口缓存数据
 * 
 * @param[in]   FrameData            待处理的串口数据
 * @param[in]   data_len        串口数据大小
 * @return  char             返回数据处理情况
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-22
 */
char UartRevStream(unsigned char *FrameData, unsigned char data_len)
{
    unsigned char n = 0;

    if(FrameData == NULL)
    {
        ErrorLog("param err \r\n");
        return 1;
    }
    for(n = 0; n < data_len; n++)
    {
        UartRevOneByte(FrameData[n]);
    }

    return 0;
}

/**
 * @brief   串口协议初始化
 * 
 * @note    该函数需在mcu初始化时调用该函数初始化串口处理。
 *  
 * @return  char      返回初始化状态 
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-22
 */
char UartProtocolInit(void)
{
    revBuffPt_in = (unsigned char *)revBuff;
    revBuffPt_out = (unsigned char *)revBuff;

    WifiWorkState = WIFI_SATE_UNKNOW;

    return 0;
}

/**
 * @brief   判断接收缓冲区是否为空
 * 
 * @return  TRUE  接收缓冲区为空
 * @return  FALSE 接收缓冲区不为空
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-28
 */
static BOOL revBuffIsEmpty(void)
{
    if(revBuffPt_in != revBuffPt_out)
    {
        return FALSE;
    }
    else
    {
        return TRUE;
    }
}


/**
 * @brief   从接收缓冲区获取数据
 * 
 * @return  unsigned char  已获取到的数据
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-28
 */
static unsigned char TakeDataFromBuff(void)
{
    unsigned char value = 0;

    if(revBuffPt_in != revBuffPt_out)
    {
        //有数据
        if(revBuffPt_out >= (unsigned char *)(revBuff + sizeof(revBuff))) {
            //数据已经到末尾
            revBuffPt_out = (unsigned char *)(revBuff);
        }
        
        value = *revBuffPt_out ++;   
    }

    return value;
}


/**
 * @brief   串口数据处理
 * 
 * @note    该函数需在while循环中循环调用
 * 
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-22
 */
void UartProcessPro(void)
{
    unsigned int dataLen = 0;
    static unsigned int revLen = 0;
    unsigned int offset = 0;

    while ((revLen < sizeof(revBuff)) && (!revBuffIsEmpty()))
    {
        processBuff[revLen++] = TakeDataFromBuff();
    }
    
    if(revLen < FRAME_MINI_LEN)
    {
        // InfoLog("frame len[%d] err \r\n", revLen);
        return;
    }

    while ((revLen - offset) >= FRAME_MINI_LEN)
    {
        if(processBuff[offset + HEAD_H] != FRAME_H)
        {
            offset++;
            InfoLog("frame head err \r\n");
            continue;
        }

        if(processBuff[offset + HEAD_L] != FRAME_L)
        {
            offset++;
            InfoLog("frame head err \r\n");
            continue;
        }


        if(processBuff[offset + PRO_VER] != MCU_REV_VER)
        {
            offset += 2;
            InfoLog("frame ver err \r\n");
            continue;
        }

        dataLen = ((processBuff[offset + LEN_H] << 8) | processBuff[offset + LEN_L]);
        dataLen += FRAME_MINI_LEN;
//                InfoLog("dataLen = %d \r\n", dataLen);
//                InfoLog("revLen = %d \r\n", revLen);
        if(dataLen > (sizeof(processBuff) + FRAME_MINI_LEN))
        {
            offset += 3;
            continue;
        }

        if((revLen - offset) < dataLen)
        {
            InfoLog("frame len err \r\n");
            break;
        }

        if(processBuff[offset + dataLen - 1] != UartCheckValue(&processBuff[offset], dataLen-1))
        {
            offset += 3;
            InfoLog("frame check err \r\n");
            continue;
        }
        ProcessUartCmd(&processBuff[offset]);
        offset += dataLen;
//                InfoLog("offset = %d \r\n", offset);
    }
    
    revLen -= offset;
    if(revLen > 0)
    {
//                InfoLog("revLen = %d \r\n", revLen);
        memcpy(processBuff, &processBuff[offset], revLen);
    }
}


