package testpkg

import (
	"fmt"
	"time"
)

var launchTime time.Time

func init() {
	launchTime = time.Now()
}

func PrintLaunchTime() {
	fmt.Println("Launch time is ", launchTime)
}
