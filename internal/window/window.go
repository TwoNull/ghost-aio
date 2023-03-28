package window

import "C"
import (
	"fmt"
	"log"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/andybrewer/mack"
	"github.com/go-vgo/robotgo"
)

var screen [2]int
var bounds [4]int

func CheckLaunch(processName string) []int32 {
	var err error
	processArr := make([]int32, 0)
	for len(processArr) == 0 {
		processArr, err = robotgo.FindIds(processName)
		if err != nil {
			log.Println(err)
		}
		time.Sleep(2000 * time.Millisecond)
	}
	return processArr
}

func SetWindowBounds(processID int32) {
	var x, y, w, h, screenW, screenH int
	err := robotgo.ActivePID(processID)
	if err != nil {
		log.Fatal(err)
	}
	if runtime.GOOS == "darwin" {
		res, err := mack.Tell(
			"System Events",
			"set _P to a reference to (processes whose unix id is "+fmt.Sprint(processID)+")",
			"set _W to a reference to windows of _P",
			"[_P's name, _W's size, _W's position]",
		)
		if err != nil {
			log.Fatal(err)
		}
		parsedRes := strings.Split(res, ", ")
		w, err = strconv.Atoi(parsedRes[1])
		h, err = strconv.Atoi(parsedRes[2])
		x, err = strconv.Atoi(parsedRes[3])
		y, err = strconv.Atoi(parsedRes[4])
		y = y + h - 720
		if err != nil {
			log.Fatal(err)
		}
	} else {
		x, y, w, h = robotgo.GetBounds(processID)
	}
	screenW, screenH = robotgo.GetScreenSize()
	log.Println([4]int{x, y, w, h})
	log.Println([2]int{screenW, screenH})
	if x+w > screenW || y+h > screenH {
		log.Fatal("Window Extends Outside Screen Bounds")
	}
	bounds = [4]int{x, y, w, h}
	screen = [2]int{screenW, screenH}
}

func GetPlayButton(processID int32) string {
	buttonColor := robotgo.GetPixelColor(bounds[0]+125, bounds[1]+bounds[3]-625)
	return buttonColor
}
