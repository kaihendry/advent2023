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
		result := advent.Trebuchet(line)
		log.Println(line, result)
		calibration += result
	}
	fmt.Println(calibration)
}
