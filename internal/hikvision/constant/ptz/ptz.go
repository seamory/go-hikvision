package ptz

type PTZCommand = uint

const LIGHT_PWRON PTZCommand = 2          //	接通灯光电源
const WIPER_PWRON PTZCommand = 3          //	接通雨刷开关
const FAN_PWRON PTZCommand = 4            //	接通风扇开关
const HEATER_PWRON PTZCommand = 5         //	接通加热器开关
const AUX_PWRON1 PTZCommand = 6           //	接通辅助设备开关
const AUX_PWRON2 PTZCommand = 7           //	接通辅助设备开关
const ZOOM_IN PTZCommand = 11             //	焦距变大(倍率变大)
const ZOOM_OUT PTZCommand = 12            //	焦距变小(倍率变小)
const FOCUS_NEAR PTZCommand = 13          //	焦点前调
const FOCUS_FAR PTZCommand = 14           //	焦点后调
const IRIS_OPEN PTZCommand = 15           //	光圈扩大
const IRIS_CLOSE PTZCommand = 16          //	光圈缩小
const TILT_UP PTZCommand = 21             //	云台上仰
const TILT_DOWN PTZCommand = 22           //	云台下俯
const PAN_LEFT PTZCommand = 23            //	云台左转
const PAN_RIGHT PTZCommand = 24           //	云台右转
const UP_LEFT PTZCommand = 25             //	云台上仰和左转
const UP_RIGHT PTZCommand = 26            //	云台上仰和右转
const DOWN_LEFT PTZCommand = 27           //	云台下俯和左转
const DOWN_RIGHT PTZCommand = 28          //	云台下俯和右转
const PAN_AUTO PTZCommand = 29            //	云台左右自动扫描
const TILT_DOWN_ZOOM_IN PTZCommand = 58   //	云台下俯和焦距变大(倍率变大)
const TILT_DOWN_ZOOM_OUT PTZCommand = 59  //	云台下俯和焦距变小(倍率变小)
const PAN_LEFT_ZOOM_IN PTZCommand = 60    //	云台左转和焦距变大(倍率变大)
const PAN_LEFT_ZOOM_OUT PTZCommand = 61   //	云台左转和焦距变小(倍率变小)
const PAN_RIGHT_ZOOM_IN PTZCommand = 62   //	云台右转和焦距变大(倍率变大)
const PAN_RIGHT_ZOOM_OUT PTZCommand = 63  //	云台右转和焦距变小(倍率变小)
const UP_LEFT_ZOOM_IN PTZCommand = 64     //	云台上仰和左转和焦距变大(倍率变大)
const UP_LEFT_ZOOM_OUT PTZCommand = 65    //	云台上仰和左转和焦距变小(倍率变小)
const UP_RIGHT_ZOOM_IN PTZCommand = 66    //	云台上仰和右转和焦距变大(倍率变大)
const UP_RIGHT_ZOOM_OUT PTZCommand = 67   //	云台上仰和右转和焦距变小(倍率变小)
const DOWN_LEFT_ZOOM_IN PTZCommand = 68   //	云台下俯和左转和焦距变大(倍率变大)
const DOWN_LEFT_ZOOM_OUT PTZCommand = 69  //	云台下俯和左转和焦距变小(倍率变小)
const DOWN_RIGHT_ZOOM_IN PTZCommand = 70  //	云台下俯和右转和焦距变大(倍率变大)
const DOWN_RIGHT_ZOOM_OUT PTZCommand = 71 //	云台下俯和右转和焦距变小(倍率变小)
const TILT_UP_ZOOM_IN PTZCommand = 72     //	云台上仰和焦距变大(倍率变大)
const TILT_UP_ZOOM_OUT PTZCommand = 73    //	云台上仰和焦距变小(倍率变小)
