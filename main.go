package lidarlite       

import (                                                                                                                          
        "time"                                                                                                                    
        "github.com/kidoman/embd"                                                                                                 
        _ "github.com/kidoman/embd/host/rpi" // This loads the RPi driver                                                         
)                                                                                                                                 

const (                                                                                                                           
        address = 0x62     

        kLidarLiteCommandControlRegister          = 0x00                                                                          
        kLidarLiteVelocityMeasurementOutput               = 0x09                                                                  
        kLidarLiteCalculateDistanceMSB                    = 0x8f                                                                  
        kLidarLiteCalculateDistanceLSB                    = 0x10                                                                  
        kLidarLitePreviousMeasuredDistanceMSB             = 0x94                                                                  
        kLidarLitePreviousMeasuredDistanceLSB             = 0x15                                                                  
        kLidarLiteHardwareVersion                          = 0x41                                                                 
        kLidarLiteSoftwareVersion                          = 0x4f                                                                 
        kLidarLiteMeasure                                  = 0x04                                                                 
		
        pollDelay = 250                                                                                                           
)      

// BMP085 represents a LIDAR-Lite v2 sensor.
type LIDARLITE struct {
	Bus  embd.I2CBus
	Poll int
}
  

// New returns a handle to a LidarLite sensor.
// https://static.garmin.com/pumac/LIDAR_Lite_v3_Operation_Manual_and_Technical_Specifications.pdf
func New(bus embd.I2CBus) *LIDARLITE {
	return &LIDARLITE{Bus: bus, Poll: pollDelay}
}


func (d *LIDARLITE) ReadDistance() (byte, error) {
	//Write 0x04 to register 0x00
	//Read register 0x01. Repeat until bit 0 (LSB) goes low
	//Read two bytes from 0x8f (High byte 0x0f then low byte 0x10) to obtain the 16-bit measured distance in cm
	err := d.Bus.WriteByteToReg(address, kLidarLiteCommandControlRegister, kLidarLiteMeasure)
	if err != nil {
		return 0, err
	}
	/*
	msb, err := d.Bus.ReadByteFromReg(address, kLidarLiteCalculateDistanceMSB)
	if err != nil {
		return 0, err
	}
	*/
	lsb, err := d.Bus.ReadByteFromReg(address, kLidarLiteCalculateDistanceLSB)
	if err != nil {
		return 0, err
	}
	distance := (msb << 8) + lsb
	
	
	return distance, nil
}

func (d *LIDARLITE) SoftwareVersion() (byte, error) {
	version, err := d.Bus.ReadByteFromReg(address, kLidarLiteHardwareVersion)
	if err != nil {
		return 0, err
	}
	return version, nil
}	



