package main

import (
	"flag"
	"github.com/MonaxGT/brofor"
)

func main() {
	dbPtr := flag.String("db", "", "Path to DB")
	broTypePtr := flag.String("type", "", "DB type. Supported: Firefox, Google Chrome")
	modePtr := flag.String("mode", "df", "Mode. Supported: Forensic mode and Threat Hunting mode with real-rime monitoring and sent nwe data to log server")
	flag.Parse()

	conf, err := brofor.New(*dbPtr, *broTypePtr)
	if err != nil {
		panic(err)
	}
	err = conf.Run(*modePtr)
	if err != nil {
		panic(err)
	}
}
