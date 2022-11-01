package main

import (
	"context"
	"fmt"
	"main/internal/hikvision"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ctx := context.Background()
	hc := hikvision.New()
	fmt.Println(hc.Login("192.168.0.10", 8000, "admin", "password"))
	fmt.Println(hc.Login("192.168.0.11", 8000, "admin", "password"))
	wg.Add(1)
	fmt.Println(hc.Users)
	go func(ctx context.Context) {
		hc.EventListener("192.168.0.10")
		hc.EventListener("192.168.0.11")
		for {
			select {
			case <-ctx.Done():
				wg.Done()
			}
		}
	}(ctx)
	wg.Wait()
	hc.Cleanup()
	//lUserID := hikvision.Login()
	// hikvision.CapturePicture(lUserID, 1)
	// hikvision.PTZControl(lUserID, 1, ptz.PAN_RIGHT, 1, 1000)
	//wg.Add(1)
	//go func(ctx context.Context) {
	//	hikvision.CameraEventListener(lUserID)
	//	for {
	//		select {
	//		case <-ctx.Done():
	//			wg.Done()
	//		}
	//	}
	//}(ctx)
	//wg.Wait()
	//hikvision.Cleanup()
}
