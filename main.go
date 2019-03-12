package main

import (
	"./converter"
	"flag"
	"time"
)

func main() {
	json := flag.String("j", "sample/header.json", "string flag")
	log := flag.String("l", "sample/sqloutput.log", "string flag")
	output := flag.String("o", "sample/convert.log", "string flag")
	flag.Parse()

	logName := *output + "_" + time.Now().Format("2006-01-02")
	ltsvconvert.Ltsvout(*json, *log, logName)
}
