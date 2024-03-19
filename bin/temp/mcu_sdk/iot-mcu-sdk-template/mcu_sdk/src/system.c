/**
 * @file    system.c
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



#define             STR_PID                 "pid"
#define             STR_VER                 "ver"
#define             STR_FLAG                "flag"



/* 串口发送数据缓存 */
static unsigned char sedBuff[MCU_UART_BUFF_MAX_LEN] = {0};

extern const McuCmdInfo CmdInfoList[];

static int WifiRssiValue = 0;


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
int GetWifiRssiValue(void)
{
	return WifiRssiValue;
}


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
char UartProFrameSend(unsigned char cmdType, unsigned char *FrameData, unsigned char dataLen)
{
    char   ret = 0;

    sedBuff[HEAD_H] = FRAME_H;
    sedBuff[HEAD_L] = FRAME_L;
    sedBuff[PRO_VER] = MCU_SED_VER;
    sedBuff[CMD_TYPE] = cmdType;

    sedBuff[LEN_H] = ((dataLen & 0xff00) >> 8);
    sedBuff[LEN_L] =  (dataLen & 0xff);

    if(dataLen + FRAME_MINI_LEN > MCU_UART_BUFF_MAX_LEN)
    {
        ErrorLog("FrameData len is too big\r\n");
        return 0;
    }

    if(FrameData != NULL)
    {
        memcpy(&sedBuff[DATA_LOCAL], FrameData, dataLen);
    }
    sedBuff[dataLen + FRAME_MINI_LEN - 1] = UartCheckValue(sedBuff, (dataLen + FRAME_MINI_LEN) - 1);

    ret = UartWriteData(sedBuff, dataLen + FRAME_MINI_LEN );

    memset(sedBuff, 0, MCU_UART_BUFF_MAX_LEN);

    return ret;
}


/**
 * @brief   心跳包数据回应
 * 
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-23
 */
static void UartSendHeatBeat(void)
{
    static BOOL mcuResetState = FALSE;
    unsigned char  FrameData = 0;

    if(mcuResetState)
    {
        FrameData = 1;
        UartProFrameSend(HEART_BEAT_CMD, &FrameData, 1);
    }
    else
    {
        FrameData = 0;
        mcuResetState = TRUE;
        UartProFrameSend(HEART_BEAT_CMD, &FrameData, 1);
    }
}

/**
 * @brief   发送OTA状态
 * 
 * @param[in]   cmdType     帧命令
 * @param[in]   cmdData     状态
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2023-04-04
 */
void UartSendOtaStatus(unsigned char cmdType, unsigned char cmdData)
{
		unsigned char  Cmd = cmdData;
	
		UartProFrameSend(cmdType, &Cmd, 1);
}

/**
 * @brief   上报产品信息
 * 
 * 
 * @author  Ai-Thinker (zhuolm@tech-now.com)
 * @date    2022-06-23
 */
static void UartSendProductInfo(void)
{
    unsigned char buff[20] = {0};

    sprintf((char *)buff, "{\"%s\":\"%s\"}", STR_PID, PRODUCT_ID);
    UartProFrameSend(PRODUCT_INFO_CMD, (unsigned char *)buff, strlen((char *)buff));
	memset(buff,0, sizeof(buff));
    sprintf((char *)buff, "{\"%s\":\"%s\"}",STR_VER, MCU_SOFTWARE_VER);
    UartProFrameSend(PRODUCT_INFO_CMD, (unsigned char *)buff, strlen((char *)buff));
	memset(buff,0, sizeof(buff));
    sprintf((char *)buff, "{\"%s\":\"%s\"}",STR_FLAG, PRODUCT_FLAG);
    UartProFrameSend(PRODUCT_INFO_CMD, (unsigned char *)buff, strlen((char *)buff));
	memset(buff,0, sizeof(buff));

}



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
unsigned char UartCheckValue(unsigned char *dataBuff, unsigned short dataLen)
{
    unsigned char sum = 0;
    unsigned short n = 0;
    for(n = 0; n < dataLen; n++)
    {
        sum += dataBuff[n];
    }
    return sum;
}




char ProcessReportDataInfo(unsigned char *FrameData, unsigned char data_len)
{
    unsigned char cmd_id = 0;
    unsigned char index = 0;
    unsigned char cmd_type = 0;
    char ret = 0;
    unsigned char cmd_len = 0;
    unsigned char count = 0;
 

    if(FrameData == NULL)
    {
        InfoLog("param err \r\n");
        return 1;
    }

	
    for(count = 0; count < data_len; )
    {
        cmd_id = FrameData[count + 0];
        cmd_type = FrameData[count + 1];
        cmd_len = ((FrameData[count + 2] << 8) | FrameData[count + 3]);
        index = GetCmdIdIndex(cmd_id);

        if(cmd_type == CmdInfoList[index].cmd_type)
        {
            ProcessDataCmd(cmd_id, &FrameData[count + 4], cmd_len);
        }
        else
        {
			if(cmd_type == DATA_TYPE_VALUE && CmdInfoList[index].cmd_type == DATA_TYPE_ENUM)
			{
				ProcessDataCmd(cmd_id, &FrameData[count + 4], cmd_len);
			}
			else
			{
				ret = 1;
				InfoLog("cmd type err \r\n");
			}
        }

        count += (cmd_len + 4);
    }

    return ret;
}





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
char ProcessUartCmd(unsigned char *DataCmd)
{
    unsigned char cmdType = 0;
    unsigned int cmdLen = 0;

    if(DataCmd == NULL)
    {
        ErrorLog("param err \r\n");
        return 1;
    }

    cmdType = DataCmd[CMD_TYPE];
    switch (cmdType)
    {
        case HEART_BEAT_CMD:
        {
            InfoLog("HEART_BEAT_CMD \r\n");
            UartSendHeatBeat();  //心跳包
        }
        break;

       case WORK_MODE_CMD:
       {
           InfoLog("WORK_MODE_CMD \r\n");
           UartReqModuleWorkMode(); //MCU向模组查询工作模式
       }
       break;

        case WIFI_STATE_CMD:
        {
            InfoLog("WIFI_STATE_CMD \r\n");
            SetWifiWorkState(DataCmd[DATA_LOCAL]); //设置WiFi的工作状态
            UartProFrameSend(WIFI_STATE_CMD, NULL, 0);//发送协议帧数据的串口函数
        }
        break;

        case STATE_QUERY_CMD:
        {
            InfoLog("STATE_QUERY_CMD \r\n");
            ReportAllDeviceState();//  MCU向模组上报所有的状态数据
        }
        break;

        case WIFI_RESET_CMD:
        {
            InfoLog("WIFI_RESET_CMD \r\n");
            SetResetState(RESET_WIFI_SUCCESS);//设置WiFi重置状态
        }
        break;

        case WIFI_MODE_CMD:
        {
            InfoLog("WIFI_MODE_CMD \r\n");
        }
        break;

        case DATA_QUERT_CMD:
        {
            InfoLog("DATA_QUERT_CMD \r\n");
            cmdLen = ((DataCmd[LEN_H] << 8) | DataCmd[LEN_L]);
            ProcessReportDataInfo(&DataCmd[DATA_LOCAL], cmdLen);
        }
        break;

        case GET_WIFI_STATUS_CMD:
        {
            InfoLog("GET_WIFI_STATUS_CMD \r\n");
            SetWifiWorkState(DataCmd[DATA_LOCAL]);//设置WiFi的工作状态
        }
        break;

        case PRODUCT_INFO_CMD:
        {
            InfoLog("PRODUCT_INFO_CMD \r\n");
            UartSendProductInfo();//上报产品信息
        }
        break;
				
        case UPDATE_START_CMD:
        {
            InfoLog("UPDATE_START_CMD \r\n");
            cmdLen = ((DataCmd[LEN_H] << 8) | DataCmd[LEN_L]);
            InfoLog("len = %d \r\n", cmdLen);
            // memcpy(dataBuff, (char *)&DataCmd[DATA_LOCAL], cmdLen);
            // InfoLog("dataBuff = %s \r\n", dataBuff);
            // UartSendOtaStatus(UPDATE_START_CMD, OTA_STATUS_START);
            UpDateFirmStart(&DataCmd[DATA_LOCAL], cmdLen);
        }
        break;

        case UPDATE_TRANS_CMD:
        {
            InfoLog("UPDATE_TRANS_CMD \r\n");
            cmdLen = ((DataCmd[LEN_H] << 8) | DataCmd[LEN_L]);
            InfoLog("cmdLen = %d \r\n", cmdLen);
            UpDataFireDownload(&DataCmd[DATA_LOCAL], cmdLen);
        }
        break;

        case GET_WIFI_RSSI_CMD:
        {
            InfoLog("GET_WIFI_RSSI_CMD \r\n");
            // WifiRssiValue = ((DataCmd[DATA_LOCAL + 3] << 24) | (DataCmd[DATA_LOCAL + 2] << 16) | (DataCmd[DATA_LOCAL + 1] << 8) | DataCmd[DATA_LOCAL + 0]);
            WifiRssiValue = DataCmd[DATA_LOCAL];
            InfoLog("rssi = %d \r\n", WifiRssiValue);
        }
        break;


        default:
            break;
    }


    return 0;
}


