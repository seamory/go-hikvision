package hikvision

import "C"
import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"main/internal/hikvision/constant/ptz"
	"mime/multipart"
	"time"
	"unsafe"
)

//#cgo CFLAGS: -I. -I./include
////#cgo LDFLAGS: -L./include -lHCNetSDK -lstdc++
//#cgo LDFLAGS: -L./include -lHCNetSDK
//#include "./include/HCNetSDK_C.h"
//#include "stdlib.h"
//#include "stdio.h"
// typedef void (*fExceptionCallBack)(unsigned long dwType, long lUserID, long lHandle, void *pUser);
// void goExceptionCallBack(unsigned long dwType, long lUserID, long lHandle, void *pUser);
// void goMessageCallback(LONG lCommand, NET_DVR_ALARMER *pAlarmer, char *pAlarmInfo, DWORD dwBufLen, void* pUser);
import "C"

//export goExceptionCallBack
func goExceptionCallBack(
	dwType C.ulong,
	lUserID C.long,
	lHandle C.long,
	pUser *C.void,
) {
	fmt.Println(dwType, lUserID, lHandle, pUser)
	switch dwType {
	case C.EXCEPTION_RECONNECT:
		fmt.Println("----------reconnect--------", time.Now())
		break
	default:
		break
	}
}

func getSDKAbility() {
	sdkAbl := C.struct_tagNET_DVR_SDKABL{}
	if C.NET_DVR_GetSDKAbility(&sdkAbl) == C.int(0) {
		fmt.Println("get sdk ability error, error: ", C.NET_DVR_GetLastError())
	}
	fmt.Println(sdkAbl)
	return
}

func initSDK() {
	C.NET_DVR_Init()
	getSDKAbility()
	C.NET_DVR_SetConnectTime(10000, 1)
	C.NET_DVR_SetReconnect(10000, 1)
}

func login(ip string, port int, username, password string) C.long {
	//C.test()
	//fmt.Println(C.getVersion(), C.getBuildVersion())
	//fmt.Println(C.NET_DVR_GetSDKVersion())
	//fmt.Println(C.NET_DVR_GetSDKBuildVersion())
	C.NET_DVR_SetExceptionCallBack_V30(0, nil, C.fExceptionCallBack(C.goExceptionCallBack), nil)

	login := C.struct_tagNET_DVR_USER_LOGIN_INFO{}
	login.bUseAsynLogin = 0
	C.strcpy(&login.sDeviceAddress[0], C.CString(ip))
	login.wPort = C.ushort(port)
	C.strcpy(&login.sUserName[0], C.CString(username))
	C.strcpy(&login.sPassword[0], C.CString(password))
	device := C.struct_tagNET_DVR_DEVICEINFO_V40{}

	var lUserID C.long
	lUserID = C.NET_DVR_Login_V40(&login, &device)
	if lUserID < 0 {
		fmt.Println("Login failed, error code: ", C.NET_DVR_GetLastError())
	}
	sn := (*[48]byte)(unsafe.Pointer(&device.struDeviceV30.sSerialNumber[0]))
	fmt.Println(string(sn[:]))
	return lUserID
}

// 云台控制
func ptzControl(lUserID int, channel uint, cmd ptz.PTZCommand, speed uint, duration int) {
	if C.NET_DVR_PTZControlWithSpeed_Other(C.long(lUserID), C.long(channel), C.ulong(cmd), C.ulong(0), C.ulong(speed)) == C.int(0) {
		fmt.Println("PTZ operate error. error: ", C.NET_DVR_GetLastError())
		return
	}
	time.Sleep(time.Duration(duration) * time.Millisecond)
	C.NET_DVR_PTZControlWithSpeed_Other(C.long(lUserID), C.long(channel), C.ulong(cmd), C.ulong(1), C.ulong(speed))
}

//export goMessageCallback
func goMessageCallback(lCommand C.long, pAlarmer *C.struct_tagNET_DVR_ALARMER, pAlarmInfo *C.char, dwBufLen C.ulong, pUser *C.void) {
	fmt.Println(lCommand, pAlarmer, pAlarmInfo, dwBufLen, pUser)
	fmt.Println("============================================================================")
	switch int(lCommand) {
	case 0x1100: // 4352
		var alarmInfo C.struct_tagNET_DVR_ALARMINFO
		C.memcpy(unsafe.Pointer(&alarmInfo), unsafe.Pointer(pAlarmInfo), C.sizeof_struct_tagNET_DVR_ALARMINFO)
		fmt.Println(alarmInfo)
		break
	case 0x4000: // 16384
		var alarmInfo C.struct_tagNET_DVR_ALARMINFO_V30
		C.memcpy(unsafe.Pointer(&alarmInfo), unsafe.Pointer(pAlarmInfo), C.sizeof_struct_tagNET_DVR_ALARMINFO_V30)
		fmt.Println(alarmInfo)
		break
	case 0x4993: // 18835 COMM_VCA_ALARM, reference 人员排队检测
		buf := C.GoBytes(unsafe.Pointer(pAlarmInfo), C.int(dwBufLen))
		r := multipart.NewReader(bytes.NewReader(buf), "MIME_boundary")
		for {
			if part, err := r.NextPart(); err != nil {
				if err == io.EOF {
					break
				} else {
					continue
				}
			} else {
				ct := part.Header.Get("Content-Type")
				slurp, err := io.ReadAll(part)
				if err != nil {
					log.Println(err)
					continue
				}
				if ct == "application/json" {
					fmt.Println(string(slurp))
				} else {
					fmt.Println(ct)
				}
			}
		}
		break
	case 0x6009: // 24585 COMM_ISAPI_ALARM, reference 人员排队检测
		isapi := C.struct_tagNET_DVR_ALARM_ISAPI_INFO{}
		C.memcpy(unsafe.Pointer(&isapi), unsafe.Pointer(pAlarmInfo), C.sizeof_struct_tagNET_DVR_ALARM_ISAPI_INFO)
		if isapi.dwAlarmDataLen > 0 {
			// alarm message
			// alarmMessage := ""
		}
		//emit
		break
	default:
		//fmt.Println(C.GoString(pAlarmInfo), dwBufLen) // this will be output the overflow char.
		fmt.Println(string(C.GoBytes(unsafe.Pointer(pAlarmInfo), C.int(dwBufLen))))
	}
	fmt.Println("============================================================================")
}

// CameraEventListener 事件监听
func cameraEventListener(lUserID int) int {
	C.NET_DVR_SetDVRMessageCallBack_V30(C.MSGCallBack(C.goMessageCallback), nil)
	var lHandle C.long
	alarmParam := C.struct_tagNET_DVR_SETUPALARM_PARAM{}
	alarmParam.dwSize = C.sizeof_struct_tagNET_DVR_SETUPALARM_PARAM
	alarmParam.byAlarmInfoType = 0
	alarmParam.byRetAlarmTypeV40 = 0
	lHandle = C.NET_DVR_SetupAlarmChan_V41(C.long(lUserID), &alarmParam)
	if lHandle < 0 {
		fmt.Println("NET_DVR_SetupAlarmChan_V41 error, error:", C.NET_DVR_GetLastError())
	}
	return int(lHandle)
}

func cameraCancelEventListener(lHandle int) bool {
	if C.NET_DVR_CloseAlarmChan_V30(C.long(lHandle)) == C.int(0) {
		log.Println("NET_DVR_CloseAlarmChan_V30 error, error:", C.NET_DVR_GetLastError())
		return false
	}
	return true
}

//func GetDeviceAbility() {
//	C.NET_DVR_GetDeviceAbility()
//}

// CapturePicture 抓取图片
func capturePicture(lUserID int, lChannel int) ([]byte, error) {
	size := C.ulonglong(3 * 1024 * 1024 * C.sizeof_char) // malloc 3MB memory size.
	pBuffer := C.malloc(size)
	defer C.free(pBuffer)
	jpegParam := C.struct_tagNET_DVR_JPEGPARA{}
	jpegParam.wPicSize = 22
	jpegParam.wPicQuality = 0
	var rSize C.ulong
	if C.NET_DVR_CaptureJPEGPicture_NEW(C.long(lUserID), C.long(lChannel), &jpegParam, (*C.char)(pBuffer), C.ulong(size), &rSize) == C.int(0) {
		fmt.Println("capture picture error. error:", C.NET_DVR_GetLastError())
		return nil, errors.New("capture picture error")
	}
	pbuf := C.GoBytes(pBuffer, C.int(rSize))
	return pbuf, nil
	// save picture data, read picture data from c memory to go memory.
}

func logout(lUserID int) {
	C.NET_DVR_Logout(C.long(lUserID))
}

func cleanup() {
	C.NET_DVR_Cleanup()
}
