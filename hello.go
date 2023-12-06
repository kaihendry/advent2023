package main

import (
	"fmt"
	"log"

	"github.com/kaihendry/one/advent"
)

func main() {
	var calibration int
	for {
		var line string
		_, err := fmt.Scanln(&line)
		if err != nil {
			log.Println(err)
			break
		}
		calibration += advent.Trebuchet(line)
	}
	fmt.Println(calibration)
}
