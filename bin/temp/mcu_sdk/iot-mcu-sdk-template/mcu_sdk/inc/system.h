/**
 * @file    system.h
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

#ifndef __SYSTEM_H
#define __SYSTEM_H
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
// #include <stdint.h>
// #include <stdbool.h>


typedef enum
{
    FALSE = 0,
    TRUE = !0
}BOOL;


    
 #define         LOG_OUTPUT                          (0)
         
 #define ErrorLog(fmt, ...)\
                 do{\
                         if(LOG_OUTPUT)\
                         {\
                             printf("[Error : %d, %s] ", __LINE__,__FUNCTION__);\
                             printf("> "fmt"\r\n", ##__VA_ARGS__);\
                             printf("\n");\
                         }\
                     }while(0)


 #define InfoLog(fmt, ...)\
                 do{\
                         if(LOG_OUTPUT)\
                         {\
                             printf("[Info : %d, %s] ", __LINE__,__FUNCTION__);\
                             printf("> "fmt"\r\n", ##__VA_ARGS__);\
                             printf("\n");\
                         }\
                     }while(0)





/*  模块工作方式选择,只能三选一,推荐使用默认模式  */
#define CONFIG_MODE     1             //默认工作模式，即模组在进入配网时，同时开启AP和BLE功能
//#define CONFIG_MODE     2            //单AP模式，即模组进入配网时，只有AP配网功能
// #define CONFIG_MODE     3             //单BLE模式，即模式进入配网时，只有BLE配网功能


/* 串口缓存最大长度， 根据MCU可支配的内存进行分配给串口缓存 */
#define  MCU_UART_BUFF_MAX_LEN              (1024)



/* MCU 软件版本号，用户可根据MCU版本而自定义该软件的版本号 */
#define  MCU_SOFTWARE_VER                   "1.0.0"


//=============================================================================
//物模型数据类型
//=============================================================================
#define         DATA_TYPE_BOOL                    0x01        //bool 类型
#define         DATA_TYPE_VALUE                   0x02        //value 类型
#define         DATA_TYPE_STRING                  0x03        //string 类型
#define         DATA_TYPE_ENUM                    0x04        //enum 类型

//=============================================================================



//=============================================================================
//数据帧字段位
//=============================================================================
#define         HEAD_H                          0
#define         HEAD_L                          1        
#define         PRO_VER                         2
#define         CMD_TYPE                        3
#define         LEN_H                           4
#define         LEN_L                           5
#define         DATA_LOCAL                      6

//=============================================================================
//数据帧类型
//=============================================================================
#define         HEART_BEAT_CMD                  0                               //心跳包
#define         PRODUCT_INFO_CMD                1                               //产品信息
#define         WORK_MODE_CMD                   2                               //查询MCU 设定的模块工作模式    
#define         WIFI_STATE_CMD                  3                               //wifi工作状态    
#define         WIFI_RESET_CMD                  4                               //重置wifi
#define         WIFI_MODE_CMD                   5                               //选择ble/AP模式    
#define         DATA_QUERT_CMD                  6                               //命令下发
#define         STATE_UPLOAD_CMD                7                               //状态上报     
#define         STATE_QUERY_CMD                 8                               //状态查询   
#define         UPDATE_START_CMD                0x0a                            //升级开始
#define         UPDATE_TRANS_CMD                0x0b                            //升级传输
#define         GET_LOCAL_TIME_CMD              0x1c                            //获取本地时间
#define         STATE_UPLOAD_SYN_CMD            0x22                            //状态上报（同步）
#define         GET_WIFI_STATUS_CMD             0x2b                            //获取当前wifi联网状态
#define         WIFI_CONNECT_TEST_CMD           0x2c                            //wifi功能测试(连接指定路由)
#define         GET_MAC_CMD                     0x2d                            //获取模块mac
#define         GET_WIFI_RSSI_CMD               0x24                            //获取路由的rssi信号值



//=============================================================================
#define         MCU_REV_VER              0x10                                            //模块发送帧协议版本号
#define         MCU_SED_VER              0x20                                            //MCU 发送帧协议版本号(默认)
#define         FRAME_MINI_LEN           0x07                                            //固定协议头长度
#define         FRAME_H                  0x5a                                            //帧头高位
#define         FRAME_L                  0xa5                                            //帧头低位
//============================================================================= 


//============================================================================
#define         MCU_MODE                CONFIG_MODE
//============================================================================

// 模组工作状态
//============================================================================
#define         BLE_CONFIG_STATE                 0x00
#define         AP_CONFIG_STATE                  0x01
#define         WIFI_NOT_CONNECTED              0x02
#define         WIFI_CONNECTED                  0x03
#define         WIFI_CONN_CLOUD                 0x04
#define         WIFI_LOW_POWER                  0x05
#define         BLE_AND_AP_STATE                0x06
#define         WIFI_SATE_UNKNOW                0xff
//============================================================================



//wifi配网的方式
//=============================================================================
#define         BLE_CONFIG                      0x0  
#define         AP_CONFIG                       0x1   
#define         AP_BLE_CONFIG                   0x2
//=============================================================================

//设置配网状态
//=============================================================================
#define         SET_CONFIG_INVAILD                  0
#define         SET_CONFIG_SUCCESS                  1

//=============================================================================

//wifi复位状态
//=============================================================================
#define         RESET_WIFI_INVAILD                  0
#define         RESET_WIFI_SUCCESS                  1

//=============================================================================
//ota 状态码
#define                 OTA_STATUS_INVAIL                                           0
#define                 OTA_STATUS_START                                            1
#define                 OTA_STATUS_MD5_ERR                                          2
#define                 OTA_STATUS_INSTALL                                          3
#define                 OTA_STATUS_OK                                               4
//=============================================================================
// 命令结构
//=============================================================================
typedef struct 
{
  unsigned char cmd_id;                              //dp序号
  unsigned char cmd_type;                            //dp类型

}McuCmdInfo;



/**
 * @brief   计算校验值，从帧头至协议内容尾字节累加求和后再对256取余
 * 
 * @param[in]   dataBuff        数据内容
 * @param[in]   dataLen         数据长度
 * @return  unsigned char             返回校验和
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-22
 */
unsigned char UartCheckValue(unsigned char *dataBuff, unsigned short dataLen);

/**
 * @brief   数据帧处理，执行帧命令
 * 
 * @param[in]   DataCmd     数据帧
 * @param[in]   data_len    数据帧长度
 * @return  unsigned char         返回执行结果
 *          0：执行成功
 *          -1：执行失败
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-22
 */
char ProcessUartCmd(unsigned char *DataCmd);


/**
 * @brief   发送协议帧数据的串口函数
 * 
 * @note    将数据域内容、帧命令通过该函数，可将数据包装成协议帧数据，然后通过串口发送出去。
 * 
 * @param[in]   cmdType     帧命令
 * @param[in]   FrameData        数据域内容
 * @param[in]   dataLen     数据长度
 * @return  char          返回数据发送情况
 *          0：发送成功
 *          -1：发送失败
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-23
 */
char UartProFrameSend(unsigned char cmdType, unsigned char *FrameData, unsigned char dataLen);



/**
 * @brief   设置WiFi的工作状态
 * 
 * @param[in]   state  设置WiFi的工作状态
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-23
 */
void SetWifiWorkState(unsigned char state);




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
unsigned char GetResetState(void);


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
void SetResetState(unsigned char state);


/**
 * @brief   发送OTA状态
 * 
 * @param[in]   cmdType     帧命令
 * @param[in]   cmdData     状态
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2023-04-04
 */
void UartSendOtaStatus(unsigned char cmdType, unsigned char cmdData);



#endif




