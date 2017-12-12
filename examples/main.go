package main

import (
	"github.com/d2r2/go-i2c"
	"lidarlite"
	"log"
	//"os"
)

func main() {
	//Init I2C on line 1 ( i2cdetect -y 1) address 0x62
	bus, err := i2c.NewI2C(0x62, 1)
	if err != nil {
		log.Fatal(err)
	}
	// Free I2C connection on exit
	defer bus.Close()

	sensor := lidarlite.New(bus)

	log.Println(sensor.ReadDistance())

}
