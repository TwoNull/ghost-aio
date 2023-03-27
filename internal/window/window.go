package window

import (
	"log"
	"time"

	"github.com/go-vgo/robotgo"
)

func TestWindow() {
	pidArr, err := robotgo.FindIds("csgo_osx64")
	if err != nil {
		log.Fatal("L")
	}
	log.Println(pidArr)

	err = robotgo.ActivePID(pidArr[0])
	if err != nil {
		log.Fatal(err)
	}
	log.Println(robotgo.GetTitle())
}

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
