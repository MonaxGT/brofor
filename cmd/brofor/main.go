package main

import (
	"flag"
	"log"

	"github.com/MonaxGT/brofor"
)

func main() {
	dbPtr := flag.String("db", "", "Path to DB")
	broTypePtr := flag.String("b", "", "DB type. Supported: Firefox, Google Chrome, Opera")
	outTypePtr := flag.String("o", "console", "Output destination. Supports: console (default), csv, excel, json")
	remoteSocketPtr := flag.String("c", "", "Remote collector for storing logs data")
	modePtr := flag.String("mode", "df", "Mode. Supported: Forensic mode and Threat Hunting mode with real-rime monitoring and sent nwe data to log server")
	livePtr := flag.Bool("live", false, "Additional extract hash downloaded files and find all existed databases in system")
	calcHash := flag.Bool("hash", false, "Enable calculate hash downloaded files")
	flag.Parse()

	conf, err := brofor.New(*broTypePtr, *outTypePtr, *remoteSocketPtr, *calcHash)
	if err != nil {
		log.Fatal(err)
	}
	err = conf.Run(*modePtr, *livePtr, *dbPtr, *broTypePtr)
	if err != nil {
		log.Fatal(err)
	}
}
