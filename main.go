package main

import (
	"image/color"
	"log"
	"machine"
	"time"

	"tinygo.org/x/drivers/gc9a01"
)

const (
	RESETPIN = machine.GPIO12
	CSPIN    = machine.GPIO9
	DCPIN    = machine.GPIO8
	BLPIN    = machine.GPIO25

	// Default Serial Clock Bus 1 for SPI communications
	SPI1_SCK_PIN = machine.GPIO10
	// Default Serial Out Bus 1 for SPI communications
	SPI1_SDO_PIN = machine.GPIO11 // Tx
	// Default Serial In Bus 1 for SPI communications
	SPI1_SDI_PIN = machine.GPIO11 //machine.GPIO12 // Rx
)

func main() {
	spi := machine.SPI1
	conf := machine.SPIConfig{
		Frequency: 40 * machine.MHz,
	}

	if err := spi.Configure(conf); err != nil {
		log.Fatal(err)
	}

	lcd := gc9a01.New(spi, RESETPIN, DCPIN, CSPIN, BLPIN)
	lcd.Configure(gc9a01.Config{})

	width, height := lcd.Size()

	white := color.RGBA{255, 255, 255, 255}
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	green := color.RGBA{0, 255, 0, 255}
	black := color.RGBA{0, 0, 0, 255}

	// lcd.FillScreen(white)

	lcd.FillRectangle(0, 0, width/2, height/2, white)
	lcd.FillRectangle(width/2, 0, width/2, height/2, red)
	lcd.FillRectangle(0, height/2, width/2, height/2, green)
	lcd.FillRectangle(width/2, height/2, width/2, height/2, blue)
	lcd.FillRectangle(width/4, height/4, width/2, height/2, black)

	for {
		time.Sleep(time.Hour)
	}
}
