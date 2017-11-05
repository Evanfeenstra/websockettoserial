package main

import (
	"log"
	"net/http"
	"runtime"

	"github.com/olahol/melody"
	"github.com/tarm/serial"
	"io/ioutil"
	"strings"
)

var Melody = melody.New()

func init() {
	// Verbose logging with file name and line number
	log.SetFlags(log.Lshortfile)

	// Use all CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())

	Melody.Upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
}

func findArduino() string {
	contents, _ := ioutil.ReadDir("/dev")

	// Look for what is mostly likely the Arduino device
	for _, f := range contents {
		if strings.Contains(f.Name(), "cu.usbmodem") {
			return "/dev/" + f.Name()
		}
	}

	// Have not been able to find a USB device that 'looks'
	// like an Arduino.
	return ""
}

func main() {
	// Start the listener

	c := &serial.Config{Name: findArduino(), Baud: 9600}
	serial, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal("No Arduino Found")
	}

	Melody.HandleConnect(func(s *melody.Session) {
		log.Println("Melody Connected")
	})

	Melody.HandleMessage(func(s *melody.Session, msg []byte) {

		log.Println(string(msg))
		_, err := serial.Write([]byte(string(msg) + "\n"))
		if err != nil {
			log.Fatal(err)
		}

	})

	RunServer(LoadHTTP(), LoadHTTPS())
}
