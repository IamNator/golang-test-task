package main

import (
	"flag"
	"fmt"
	"twitch_chat_analysis/app/pusher"
	"twitch_chat_analysis/app/report"
	"twitch_chat_analysis/app/worker"
)

func Run() {
	app := flag.String("app", "", "the application to run")
	flag.Parse()

	switch *app {
	case "pusher":
		pusher.Run()
	case "worker":
		worker.Run()
	case "report":
		report.Run()
	default:
		fmt.Println("Invalid app name")
	}
}
