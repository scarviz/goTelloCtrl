package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/scarviz/goTelloCtrl/logic"
	"github.com/scarviz/goTelloCtrl/page"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/joystick"
)

func main() {
	drone := tello.NewDriver("8890")

	mplayerIn, err := getMPlayerIn()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	ti := logic.GetTelloInfo(drone, mplayerIn)

	td := page.GetTelloData(ti)
	setHandler(td)
	go func(port int) {
		err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
		if err != nil {
			os.Exit(-1)
		}
	}(8880)

	joystickAdaptor := joystick.NewAdaptor()
	stick := joystick.NewDriver(joystickAdaptor, "joystick_ps4.json")
	ji := logic.GetPS4JoystickInfo(drone)

	work := func() {
		droneWork(drone, ti)
		joystickWork(stick, ji)
	}
	robot := gobot.NewRobot("tello",
		[]gobot.Connection{joystickAdaptor},
		[]gobot.Device{drone, stick},
		work,
	)

	robot.Start()
}

/*
setHandler : Handlerの設定
*/
func setHandler(td *page.TelloData) {
	http.HandleFunc("/battery", td.Battery)
	http.HandleFunc("/height", td.Height)
	http.HandleFunc("/takeoff", td.TakeOff)
	http.HandleFunc("/land", td.Land)
	http.HandleFunc("/palmland", td.PalmLand)
	http.HandleFunc("/up", td.Up)
	http.HandleFunc("/down", td.Down)
	http.HandleFunc("/forward", td.Forward)
	http.HandleFunc("/backward", td.Backward)
	http.HandleFunc("/left", td.Left)
	http.HandleFunc("/right", td.Right)
	http.HandleFunc("/turnright", td.Clockwise)
	http.HandleFunc("/turnleft", td.CounterClockwise)
}

/*
droneWork : ドローン動作
*/
func droneWork(drone *tello.Driver, ti *logic.TelloInfo) {
	drone.On(tello.ConnectedEvent, ti.Connected)
	drone.On(tello.VideoFrameEvent, ti.VideoFrame)
	drone.On(tello.FlightDataEvent, ti.FlightData)

	gobot.Every(5*time.Second, ti.BatteryCheck)
}

/*
joystickWork : Joystick動作
*/
func joystickWork(stick *joystick.Driver, ji *logic.PS4JoystickInfo) {
	stick.On(joystick.TriangleRelease, ji.TakeOff)
	stick.On(joystick.CirclePress, ji.Land)
	stick.On(joystick.SquarePress, ji.BackFlip)
	stick.On(joystick.LeftY, ji.UpDown)
	stick.On(joystick.LeftX, ji.Turn)
	stick.On(joystick.RightY, ji.ForwardBack)
	stick.On(joystick.RightX, ji.RightLeft)
	gobot.Every(10*time.Millisecond, ji.LeftStick)
	gobot.Every(10*time.Millisecond, ji.RightStick)

	stick.On(joystick.XPress, ji.AppEnd)
}

/*
getMPlayerIn : MPlayerのInputを取得する
*/
func getMPlayerIn() (in io.WriteCloser, err error) {
	mplayer := exec.Command("mplayer", "-fps", "30", "-")
	in, _ = mplayer.StdinPipe()
	err = mplayer.Start()
	return
}
