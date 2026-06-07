package main

import (
	"flag"
	"log"

	"github.com/safelyyou/fleet-monitor/internal/configuration"
	webserver "github.com/safelyyou/fleet-monitor/internal/web-server"
)

func main() {
	devicesPath := flag.String("devices", "devices.csv", "path to devices CSV file")
	addr := flag.String("addr", ":6733", "listen address")
	flag.Parse()

	validDevices, err := configuration.LoadDevices(*devicesPath)
	if err != nil {
		log.Fatalf("loading devices: %v", err)
	}
	log.Printf("loaded %d valid devices", len(validDevices))

	s := webserver.NewServer(validDevices)
	s.Run(*addr)
}
