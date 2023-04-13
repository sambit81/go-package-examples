package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	//Using glog
	file, err := os.OpenFile("C:\\Users\\HP\\OneDrive\\Documents\\error.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.SetOutput(file)
	log.Println("Hello World")

	//Using logrus
	logrus.Fatal("Fatal error")
	logrus.Panic("Panic error")

	logrus.SetLevel(logrus.DebugLevel)
	logrus.Debug("HelloWorld")

	logrus.SetLevel(logrus.TraceLevel)
	logrus.Trace("HelloWorld")

}
