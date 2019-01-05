package page

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/scarviz/goTelloCtrl/logic"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
)

const (
	forward          = "Forward"
	backward         = "backward"
	left             = "left"
	right            = "right"
	clockwise        = "clockwise"
	counterClockwise = "counterClockwise"
)

/*
TelloData : Telloデータ
*/
type TelloData struct {
	telloInfo *logic.TelloInfo
}

/*
GetTelloData : Telloデータ取得
*/
func GetTelloData(telloInfo *logic.TelloInfo) *TelloData {
	return &TelloData{telloInfo: telloInfo}
}

/*
Battery : バッテリー
*/
func (td *TelloData) Battery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%d%%", td.telloInfo.Data.BatteryPercentage)))
}

/*
Height : 高さ
*/
func (td *TelloData) Height(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	height := (float64)(td.telloInfo.Data.Height) / 10.0
	w.Write([]byte(fmt.Sprintf("%.1fm", height)))
}

/*
TakeOff : 離陸
*/
func (td *TelloData) TakeOff(w http.ResponseWriter, r *http.Request) {
	td.telloInfo.Drone.TakeOff()
	w.Header().Set("Content-type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

/*
Land : 着地
*/
func (td *TelloData) Land(w http.ResponseWriter, r *http.Request) {
	td.telloInfo.Drone.Land()
	w.Header().Set("Content-type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

/*
PalmLand : 手のひら着地
*/
func (td *TelloData) PalmLand(w http.ResponseWriter, r *http.Request) {
	td.telloInfo.Drone.PalmLand()
	w.Header().Set("Content-type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

/*
Up : 上昇
*/
func (td *TelloData) Up(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	add := keys.Get("add")

	up := (int)(td.telloInfo.Data.Height)
	if add == "" {
		up++
	} else {
		val, err := strconv.Atoi(add)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		up += val
	}
	td.telloInfo.Drone.Up(up)

	gobot.After(2*time.Second, func() {
		td.telloInfo.Drone.Up(0)
	})
	w.Header().Set("Content-type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

/*
Down : 降下
*/
func (td *TelloData) Down(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	add := keys.Get("add")

	down := (int)(td.telloInfo.Data.Height)
	if add == "" {
		down--
	} else {
		val, err := strconv.Atoi(add)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		down -= val
	}
	td.telloInfo.Drone.Down(down)

	gobot.After(2*time.Second, func() {
		td.telloInfo.Drone.Down(0)
	})
	w.Header().Set("Content-type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

/*
Forward : 前方移動
*/
func (td *TelloData) Forward(w http.ResponseWriter, r *http.Request) {
	ctrlTello(w, r, td.telloInfo.Drone, forward)
}

/*
Backward : 後方移動
*/
func (td *TelloData) Backward(w http.ResponseWriter, r *http.Request) {
	ctrlTello(w, r, td.telloInfo.Drone, backward)
}

/*
Left : 左移動
*/
func (td *TelloData) Left(w http.ResponseWriter, r *http.Request) {
	ctrlTello(w, r, td.telloInfo.Drone, left)
}

/*
Right : 右移動
*/
func (td *TelloData) Right(w http.ResponseWriter, r *http.Request) {
	ctrlTello(w, r, td.telloInfo.Drone, right)
}

/*
Clockwise : 右旋回
*/
func (td *TelloData) Clockwise(w http.ResponseWriter, r *http.Request) {
	ctrlTello(w, r, td.telloInfo.Drone, clockwise)
}

/*
CounterClockwise : 左旋回
*/
func (td *TelloData) CounterClockwise(w http.ResponseWriter, r *http.Request) {
	ctrlTello(w, r, td.telloInfo.Drone, counterClockwise)
}

/*
ctrlTello : Tello操縦処理
*/
func ctrlTello(w http.ResponseWriter, r *http.Request, drone *tello.Driver, ctrl string) {
	val, err := getValue(r)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch ctrl {
	case forward:
		drone.Forward(val)
		break
	case backward:
		drone.Backward(val)
		break
	case left:
		drone.Left(val)
		break
	case right:
		drone.Right(val)
		break
	case clockwise:
		drone.Clockwise(val)
		break
	case counterClockwise:
		drone.CounterClockwise(val)
		break
	}

	gobot.After(1*time.Second, func() {
		drone.Forward(0)
		drone.Left(0)
		drone.Right(0)
		drone.Clockwise(0)
	})

	w.Header().Set("Content-type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

/*
getValue : 値を取得する
*/
func getValue(r *http.Request) (int, error) {
	keys := r.URL.Query()
	valStr := keys.Get("val")

	val, err := strconv.Atoi(valStr)
	return val, err
}
