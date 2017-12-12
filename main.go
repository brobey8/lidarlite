package lidarlite

import (
	"github.com/d2r2/go-i2c"
	"log"
	//"os"
)

const (
	address = 0x62

	kLidarLiteCommandControlRegister      = 0x00
	kLidarLiteVelocityMeasurementOutput   = 0x09
	kLidarLiteCalculateDistanceMSB        = 0x8f
	kLidarLiteCalculateDistanceLSB        = 0x10
	kLidarLitePreviousMeasuredDistanceMSB = 0x94
	kLidarLitePreviousMeasuredDistanceLSB = 0x15
	kLidarLiteHardwareVersion             = 0x41
	kLidarLiteSoftwareVersion             = 0x4f
	kLidarLiteMeasure                     = 0x04
)

type LIDARLITE struct {
	Bus *i2c.I2C
}

// represents a LIDAR-Lite v2 sensor.
func New(bus *i2c.I2C) *LIDARLITE {
	return &LIDARLITE{Bus: bus}
}

func (d *LIDARLITE) ReadDistance() int {
	//Read the distance. Needs to return an error and not log it. Should not have a log dependancy
	err := d.Bus.WriteRegU8(kLidarLiteCommandControlRegister, kLidarLiteMeasure)
	if err != nil {
		log.Fatal(err)
	}

	msb, err := d.Bus.ReadRegU8(kLidarLitePreviousMeasuredDistanceMSB)
	if err != nil {
		log.Fatal(err)
	}

	lsb, err := d.Bus.ReadRegU8(kLidarLitePreviousMeasuredDistanceLSB)
	if err != nil {
		log.Fatal(err)
	}

	distance := (msb << 8) + lsb

	return int(distance)
}
