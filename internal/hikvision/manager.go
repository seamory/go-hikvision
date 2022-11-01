package hikvision

import "C"
import (
	"main/internal/hikvision/constant/code"
	"main/internal/hikvision/constant/ptz"
)

type UserContext struct {
	Id       int
	Ip       string
	Port     int
	Username string
	Password string
}

type HikvisionContext struct {
	Users map[string]UserContext
}

func New() HikvisionContext {
	initSDK()
	return HikvisionContext{
		Users: map[string]UserContext{},
	}
}

func (x *HikvisionContext) Login(ip string, port int, username, password string) code.MsgCode {
	if _, ok := x.Users[ip]; ok {
		return code.Logined
	}
	id := login(ip, port, username, password)
	if id < 0 {
		return code.Failed
	}
	x.Users[ip] = UserContext{
		Id:       int(id),
		Ip:       ip,
		Port:     port,
		Username: username,
		Password: password,
	}
	return code.Success
}

func (x HikvisionContext) PTZControl(ip string, channel uint, cmd ptz.PTZCommand, speed uint, duration int) code.MsgCode {
	u, ok := x.Users[ip]
	if !ok {
		return code.NoExist
	}
	ptzControl(u.Id, channel, cmd, speed, duration)
	return code.Success
}

func (x HikvisionContext) CapturePicture(ip string, channel int) ([]byte, code.MsgCode) {
	u, ok := x.Users[ip]
	if !ok {
		return nil, code.NoExist
	}
	picture, err := capturePicture(u.Id, channel)
	if err != nil {
		return nil, code.Failed
	}
	return picture, code.Success
}

func (x HikvisionContext) EventListener(ip string) code.MsgCode {
	u, ok := x.Users[ip]
	if !ok {
		return code.NoExist
	}
	lHandle := cameraEventListener(u.Id)
	if lHandle < 0 {
		return code.Failed
	}
	return code.Success
}

func (x *HikvisionContext) Cleanup() {
	for _, context := range x.Users {
		logout(context.Id)
	}
	cleanup()
}
