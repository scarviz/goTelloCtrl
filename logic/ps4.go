package logic

import (
	"fmt"
	"os"
	"sync/atomic"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
)

/*
PS4JoystickInfo : PS4Joystick情報
*/
type PS4JoystickInfo struct {
	drone *tello.Driver
}

type pair struct {
	x float64
	y float64
}

var leftX, leftY, rightX, rightY atomic.Value

const offset = 32767.0

func init() {
	leftX.Store(float64(0.0))
	leftY.Store(float64(0.0))
	rightX.Store(float64(0.0))
	rightY.Store(float64(0.0))
}

/*
GetPS4JoystickInfo : PS4Joystick情報を取得する
*/
func GetPS4JoystickInfo(drone *tello.Driver) *PS4JoystickInfo {
	return &PS4JoystickInfo{drone: drone}
}

/*
TakeOff : 離陸
*/
func (ji *PS4JoystickInfo) TakeOff(data interface{}) {
	fmt.Println("TakeOff")
	ji.drone.TakeOff()
}

/*
Land : 着陸
*/
func (ji *PS4JoystickInfo) Land(data interface{}) {
	fmt.Println("Land")
	ji.drone.Land()
}

/*
BackFlip : バックフリップ操作
*/
func (ji *PS4JoystickInfo) BackFlip(data interface{}) {
	fmt.Println("BackFlip")
	ji.drone.BackFlip()
}

/*
UpDown : 昇降操作
*/
func (ji *PS4JoystickInfo) UpDown(data interface{}) {
	val := float64(data.(int16))
	leftY.Store(val)
}

/*
Turn : 旋回操作
*/
func (ji *PS4JoystickInfo) Turn(data interface{}) {
	val := float64(data.(int16))
	leftX.Store(val)
}

/*
RightLeft : 左右移動操作
*/
func (ji *PS4JoystickInfo) RightLeft(data interface{}) {
	val := float64(data.(int16))
	rightX.Store(val)
}

/*
ForwardBack : 前後移動操作
*/
func (ji *PS4JoystickInfo) ForwardBack(data interface{}) {
	val := float64(data.(int16))
	rightY.Store(val)
}

/*
RightStick : 右スティック操作処理
*/
func (ji *PS4JoystickInfo) RightStick() {
	rightStick := getRightStick()

	switch {
	case rightStick.y < -10:
		ji.drone.Forward(tello.ValidatePitch(rightStick.y, offset))
	case rightStick.y > 10:
		ji.drone.Backward(tello.ValidatePitch(rightStick.y, offset))
	default:
		ji.drone.Forward(0)
	}

	switch {
	case rightStick.x > 10:
		ji.drone.Right(tello.ValidatePitch(rightStick.x, offset))
	case rightStick.x < -10:
		ji.drone.Left(tello.ValidatePitch(rightStick.x, offset))
	default:
		ji.drone.Right(0)
	}
}

/*
LeftStick : 左スティック操作処理
*/
func (ji *PS4JoystickInfo) LeftStick() {
	leftStick := getLeftStick()
	switch {
	case leftStick.y < -10:
		ji.drone.Up(tello.ValidatePitch(leftStick.y, offset))
	case leftStick.y > 10:
		ji.drone.Down(tello.ValidatePitch(leftStick.y, offset))
	default:
		ji.drone.Up(0)
	}

	switch {
	case leftStick.x > 20:
		ji.drone.Clockwise(tello.ValidatePitch(leftStick.x, offset))
	case leftStick.x < -20:
		ji.drone.CounterClockwise(tello.ValidatePitch(leftStick.x, offset))
	default:
		ji.drone.Clockwise(0)
	}
}

/*
AppEnd : アプリ終了
*/
func (ji *PS4JoystickInfo) AppEnd(data interface{}) {
	fmt.Println("AppEnd")
	ji.drone.Land()
	gobot.After(5*time.Second, func() {
		os.Exit(0)
	})
}

func getLeftStick() pair {
	s := pair{x: 0, y: 0}
	s.x = leftX.Load().(float64)
	s.y = leftY.Load().(float64)
	return s
}

func getRightStick() pair {
	s := pair{x: 0, y: 0}
	s.x = rightX.Load().(float64)
	s.y = rightY.Load().(float64)
	return s
}
