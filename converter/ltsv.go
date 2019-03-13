package ltsvconvert

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

type Keyjson struct {
	Sql    string
	Header []string
}

func Ltsvout(header string, log string, output string) {

	jsonFile, err := os.Open(header)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	logFile, err := os.Open(log)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()
	logScanner := bufio.NewScanner(logFile)

	outputLog, err := os.OpenFile(output, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer outputLog.Close()
	ltsvWriter := bufio.NewWriter(outputLog)

	jsonBytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	key := new(Keyjson)
	json.Unmarshal(jsonBytes, key)
	var prn string

	for i := 1; logScanner.Scan(); i++ {
		vAry := strings.Split(logScanner.Text(), "|")
		for i, ary := range key.Header {
			prn = prn + ary + ":" + strings.TrimSpace(vAry[i]) + "\t"
		}
		ltsvWriter.WriteString(strings.TrimRight(prn, "\t") + "\n")
		prn = ""
	}
	ltsvWriter.Flush()
}
