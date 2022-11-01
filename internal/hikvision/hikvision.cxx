#include <stdio.h>
#include <iostream>
#include "./include/HCNetSDK_C.h"
#include "hikvision.hxx"

void CALLBACK g_ExceptionCallBack(DWORD dwType, LONG lUserID, LONG lHandle, void *pUser)
{
    char tempbuf[256] = {0};
    switch(dwType)
    {
    case EXCEPTION_RECONNECT:    //预览时重连
        printf("----------reconnect--------%d\n", time(NULL));
    break;
    default:
    break;
    }
}

void CALLBACK MessageCallback(LONG lCommand, NET_DVR_ALARMER *pAlarmer, char *pAlarmInfo, DWORD dwBufLen, void* pUser)
{
      int i;
      NET_DVR_ALARMINFO struAlarmInfo;
      memcpy(&struAlarmInfo, pAlarmInfo, sizeof(NET_DVR_ALARMINFO));
//      goMessageCallback(lCommand, pAlarmInfo, dwBufLen, pUser);
      switch(lCommand)
      {
      case COMM_ALARM:
          {
              switch (struAlarmInfo.dwAlarmType)
              {
              case 3: //移动侦测报警
                   for (i=0; i<16; i++)   //#define MAX_CHANNUM   16  //最大通道数
                   {
                       if (struAlarmInfo.dwChannel[i] == 1)
                       {
                           printf("发生移动侦测报警的通道号 %d\n", i+1);
                       }
                   }
              break;
              default:
              break;
              }
           }
      break;
      default:
      break;
      }
}

typedef NET_DVR_ALARMER export_NET_DVR_ALARMER;

void test() {
    printf("hello world\r\n");
}

unsigned long getVersion() {
    return NET_DVR_GetSDKVersion();
}

unsigned long getBuildVersion() {
    return NET_DVR_GetSDKBuildVersion();
}

unsigned long login(
    char* ip,
    unsigned long port,
    char* username,
    char* password
) {
    // 初始化
    NET_DVR_Init();
    //设置连接时间与重连时间
    NET_DVR_SetConnectTime(10000, 1);
    NET_DVR_SetReconnect(10000, 1);
    NET_DVR_SetExceptionCallBack_V30(0, NULL, g_ExceptionCallBack, NULL);
    LONG lUserID;
    //登录参数，包括设备地址、登录用户、密码等
    NET_DVR_USER_LOGIN_INFO struLoginInfo = {0};
    struLoginInfo.bUseAsynLogin = 0; //同步登录方式
    strcpy(struLoginInfo.sDeviceAddress, ip); //设备IP地址
    struLoginInfo.wPort = port; //设备服务端口
    strcpy(struLoginInfo.sUserName, username); //设备登录用户名
    strcpy(struLoginInfo.sPassword, password); //设备登录密码

    //设备信息, 输出参数
    NET_DVR_DEVICEINFO_V40 struDeviceInfoV40 = {0};

    lUserID = NET_DVR_Login_V40(&struLoginInfo, &struDeviceInfoV40);
    printf("last user id: %ld\n", lUserID);
    if (lUserID < 0)
    {
        printf("Login failed, error code: %d\n", NET_DVR_GetLastError());
        NET_DVR_Cleanup();
        return -2;
    }
    printf("device serial number: %s\n", struDeviceInfoV40.struDeviceV30.sSerialNumber);

    NET_DVR_SetDVRMessageCallBack_V30(MessageCallback, NULL);
    return lUserID;
}
