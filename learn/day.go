package main

import (
	"fmt"
	"time"
)

func main()  {

	t := time.Now()
	dayTime := fmt.Sprintf("%d-%d-%d %d:%d:%d", t.Year(), t.Month(),t.Day(),t.Hour(),t.Minute(),t.Second())
	fmt.Printf("time now is %s,format time is %s", dayTime,t.Format("2006-01-02 15:04:05"))

}