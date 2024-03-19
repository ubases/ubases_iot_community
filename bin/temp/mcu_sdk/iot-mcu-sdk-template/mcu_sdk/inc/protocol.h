/**
 * @file    protocol.h
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
#ifndef __PROTOCOL_H
#define __PROTOCOL_H
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
// #include <stdint.h>
// #include <stdbool.h>



/**
 * @brief   MCU向模组上报所有的状态数据
 * 
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-29
 */
void ReportAllDeviceState(void);


/**
 * @brief   根据数据Id获取序号
 * 
 * @param[in]   cmd_id   数据id
 * @return  unsigned char      返回序号
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-07-13
 */
unsigned char GetCmdIdIndex(unsigned char cmd_id);


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
void ProcessDataCmd(unsigned char cmd_id, unsigned char *FrameData, unsigned short DataLen);



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
char UartWriteData(unsigned char *dataBuff, unsigned char dataLen);


/**
 * @brief   通知MCU设备有新OTA固件可更新，可接收到可升级的版本号以及OTA固件的md5校验数据
 * 
 * @param[in]   data            版本好以及md5数据字符串
 * @param[in]   dataLen         数据长度
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2023-04-04
 */
void UpDateFirmStart(unsigned char *data, unsigned short dataLen);


/**
 * @brief    接收处理OTA固件数据，接收到前4个字节数据是OTA固件包的长度，其余字节均是OTA固件包数据
 * 
 * @param[in]   data       数据长度以及OTA固件数据
 * @param[in]   dataLen    数据长度
 *  
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2023-04-04
 */
void UpDataFireDownload(unsigned char *data, unsigned short dataLen);



#endif
