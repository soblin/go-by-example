package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(2 * time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	// you can request to stop a timer, if you could make it by the time it fires.
	time.Sleep(3 * time.Second)
	stop2 := timer2.Stop()

	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
