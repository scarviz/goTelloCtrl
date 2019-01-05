package logic

import (
	"fmt"
	"io"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
)

/*
TelloInfo : Tello情報
*/
type TelloInfo struct {
	Drone     *tello.Driver
	Data      *tello.FlightData
	mplayerIn io.WriteCloser
}

/*
GetTelloInfo : Tello情報を取得する
*/
func GetTelloInfo(drone *tello.Driver, mplayerIn io.WriteCloser) *TelloInfo {
	return &TelloInfo{Drone: drone, mplayerIn: mplayerIn}
}

/*
Connected : 接続時処理
*/
func (ti *TelloInfo) Connected(data interface{}) {
	fmt.Println("Connected")
	ti.Drone.StartVideo()
	ti.Drone.SetVideoEncoderRate(tello.VideoBitRate4M)
	gobot.Every(100*time.Millisecond, func() {
		ti.Drone.StartVideo()
	})
}

/*
VideoFrame : ビデオフレーム処理
*/
func (ti *TelloInfo) VideoFrame(data interface{}) {
	pkt := data.([]byte)
	if _, err := ti.mplayerIn.Write(pkt); err != nil {
		fmt.Println(err)
	}
}

/*
FlightData : フライト情報処理
*/
func (ti *TelloInfo) FlightData(data interface{}) {
	fd, ok := data.(*tello.FlightData)
	if !ok {
		fmt.Println("no data")
		return
	}
	ti.Data = fd
}

/*
BatteryCheck : バッテリーチェック
*/
func (ti *TelloInfo) BatteryCheck() {
	fmt.Println(fmt.Sprintf("%d%%", ti.Data.BatteryPercentage))
}
