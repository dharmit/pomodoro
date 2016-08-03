package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/0xAX/notificator"
	"github.com/ddo/go-spin"
)

var notify *notificator.Notificator

func printCountDown(message, notifier string, minutes int) {
	notify = notificator.New(notificator.Options{
		DefaultIcon: "icon/default.png",
		AppName:     "Pomodoro",
	})

	workSeconds := 60 * minutes
	spinner := spin.New("")

	// ticker ensures time being printed is refreshed every 1s
	workTicker := time.NewTicker(time.Second * 1)

	go func() {
		for _ = range workTicker.C {
			fmt.Printf("\r %s %s %.2d:%.2d  ", message, spinner.Spin(), workSeconds/60, workSeconds%60)
			workSeconds--
		}
	}()

	<-time.After(time.Second * time.Duration(workSeconds))

	notify.Push("Pomodoro", "Time for "+notifier, "", notificator.UR_NORMAL)

	workTicker.Stop()
}

func main() {
	work := flag.Int("w", 25, "Work duration")
	brk := flag.Int("b", 5, "Break duration")
	pomodoros := flag.Int("p", 4, "Number of pomodoros in one sitting")

	flag.Parse()

	for i := 0; i < *pomodoros; i++ {
		printCountDown("Work for", "a break!", *work)
		printCountDown("Break for", "work", *brk)
	}
}
