/**
 * @file    mcu_api.h
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
#ifndef __MCU_API_H
#define __MCU_API_H
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include "system.h"
// #include <stdint.h>
// #include <stdbool.h>




/**
 * @brief   获取设置配网状态
 * 
 * @return  unsigned char  配网设置是否成功
 *          SET_CONFIG_INVAILD：设置失败
 *          SET_CONFIG_SUCCESS：设置成功
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-23
 */
unsigned char GetWifiConfigState(void);


/**
 * @brief   MCU向模组查询工作模式
 * 
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-29
 */
void UartReqModuleWorkMode(void);


/**
 * @brief   MCU向模组上报选择进入指定配网模式
 * 
 * @param[in]   state   选择的配网模式
 *              0：BLE配网
 *              1：AP配网
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-28
 */
void UartReportWifiConfig(unsigned char state);


/**
 * @brief   获取工wifi作状态
 * 
 * @return  unsigned char  返回工作状态
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-23
 */
unsigned char GetWifiWorkState(void);



/**
 * @brief   MCU向模组发送重置网络状态
 * 
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-28
 */
void UartReportResetState(void);


/**
 * @brief   MCU向模组查询工作模式
 * 
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-28
 */
void UartCheckWorkMode(void);



/**
 * @brief   MCU向模组上报bool类型的状态数据
 * 
 * @param[in]   cmd_id  数据id
 * @param[in]   value   状态数据
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-28
 */
void UartReportBoolTypeData(unsigned char cmd_id, BOOL value);


/**
 * @brief   MCU向模组上报整数类型的状态数据
 * 
 * @param[in]   cmd_id  数据id
 * @param[in]   value   状态数据
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-28
 */
void UartReportValueTypeData(unsigned char cmd_id, unsigned long value);



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
void UartReportStringTypeData(unsigned char cmd_id, unsigned char *value, unsigned char value_len);

/**
 * @brief   MCU向模组上报Enum类型的状态数据
 * 
 * @param[in]   cmd_id  数据id
 * @param[in]   value   状态数据
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-28
 */
void UartReportEnumTypeData(unsigned char cmd_id, unsigned char value);


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
BOOL ProGetdpidValueBool(unsigned char *value);


/**
 * @brief   获取枚举数据
 * 
 * @param[in]   value       数据内容
 * @return  unsigned char   获取的数据内容
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-08-31
 */
unsigned char ProGetdpidValueEnum(unsigned char *value);


/**
 * @brief   获取整形数据
 * 
 * @param[in]   value       数据内容
 * @return  unsigned long   返回数据值
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-08-31
 */
unsigned long ProGetdpidValueInt(unsigned char *value);


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
signed char AilinkCheckVersion(char *target_version, char *source_version);


#ifdef MCU_CMD_UPLOAD_SYN

/**
 * @brief   MCU向模组同步上报bool类型的状态数据
 * 
 * @param[in]   cmd_id  数据id
 * @param[in]   value   状态数据
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-28
 */
void UartSynReportBoolTypeData(unsigned char cmd_id, BOOL value);

/**
 * @brief   MCU向模组同步上报整数类型的状态数据
 * 
 * @param[in]   cmd_id  数据id
 * @param[in]   value   状态数据
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-28
 */
void UartSynReportValueTypeData(unsigned char cmd_id, unsigned long value);


/**
 * @brief   MCU向模组同步上报字符串类型的状态数据
 * 
 * @param[in]   cmd_id  数据id
 * @param[in]   value   状态数据
 * @param[in]   value_len   状态数据长度
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-28
 */
void UartSynReportStringTypeData(unsigned char cmd_id, char *value, unsigned char value_len);

/**
 * @brief   MCU向模组同步上报Enum类型的状态数据
 * 
 * @param[in]   cmd_id  数据id
 * @param[in]   value   状态数据
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-28
 */
void UartSynReportEnumTypeData(unsigned char cmd_id, unsigned char value);


#endif




/**
 * @brief   打包所有类型的数据，将所有类型的数据缓存于buff中，当将所需数据打包好后，可通过McuSendAllTypeData该函数发送给模组
 * 
 * @param[in]   cmd_id          数据id
 * @param[in]   cmd_type        数据类型
 * @param[in]   FrameData            数据内容
 * @param[in]   data_len        数据长度
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-30
 */
void PacketAllTypeData(unsigned char cmd_id, unsigned char cmd_type, unsigned char *FrameData, unsigned char data_len);


/**
 * @brief   通过串口将缓存于buff中的所有类型数据发送出去
 * 
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-30
 */
void UartReportAllTypeData(void);



/**
 * @brief   MCU向模组请求本地时间
 * 
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-28
 */
void UartReqLocalTime(void);


/**
 * @brief   MCU请求模组连接指定路由
 * 
 * @param[in]   ssid        路由的ssid
 * @param[in]   passwd      路由的passwd
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-28
 */
void UartReqConnectWifi(char *ssid, char *passwd);


/**
 * @brief   MCU向模组请求获取模组的MAC地址
 * 
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-28
 */
void UartReqModuleMac(void);

/**
 * @brief   该函数主要是逐一字节接收串口数据，若是串口采用缓存方式处理串口数据，可调用UartRevStream该函数处理。
 * 
 * @param[in]   FrameData        串口数据
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-22
 */
void UartRevOneByte(unsigned char FrameData);

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
char UartRevStream(unsigned char *FrameData, unsigned char data_len);

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
char UartProtocolInit(void);



/**
 * @brief   从bool类型的数据中获取数据值
 * 
 * @param[in]   FrameData        数据内容
 * @param[in]   data_len    数据长度
 * @return  TRUE            数据值
 * @return  FALSE           数据值
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-29
 */
BOOL GetValueFromBoolTypeData(unsigned char *FrameData, unsigned char data_len);


/**
 * @brief   从枚举类型的数据中获取数据值
 * 
 * @param[in]   FrameData            数据内容
 * @param[in]   data_len        数据长度
 * @return  unsigned char             数据值
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-29
 */
unsigned char GetValueFromEnumTypeData(unsigned char *FrameData, unsigned char data_len);


/**
 * @brief   从整数类型的数据中获取数据值
 * 
 * @param[in]   FrameData            数据内容
 * @param[in]   data_len        数据长度
 * @return  unsigned long            数据值
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-29
 */
unsigned long GetValueFromValueTypeData(unsigned char *FrameData, unsigned char data_len);


/**
 * @brief   串口数据处理
 * 
 * @note    该函数需在while循环中循环调用
 * 
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-22
 */
void UartProcessPro(void);


/**
 * @brief   获取路由的rssi信号值
 * 
 * @note    通过调用UartReqConnectWifi该函数，模组已连接WiFi时，模组将会返回rssi信号值。若是模组未连接，则会返回未连接状态；且该函数将返回0值。
 * 
 * @return  int          路由的rssi信号值
 *          0：未获取到rssi信号值
 *          小于0：已获取到rssi信号值
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2023-02-23
 */
int GetWifiRssiValue(void);




#endif
