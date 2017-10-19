package btwan

import (
	"log"
	"os"
)

var syslog *log.Logger

func initLog() error {
	if "" == logfile || "stdout" == logfile {
		syslog = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	} else {
		fi, err := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		syslog = log.New(fi, "", log.Ldate|log.Ltime)
	}
	return nil
}

func info(msg ...interface{}) {
	syslog.Println(msg...)
}

func debug(msg ...interface{}) {
	syslog.Println(msg...)
}

func fatal(msg ...interface{}) {
	syslog.Println(msg...)
}
