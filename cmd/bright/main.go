package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/wouterbeets/brightness/pkg/bright"
)

func main() {
	up := flag.Float64("up", 0, "up brightness by x percentage points")
	down := flag.Float64("down", 0, "down brightness by x percentage points")
	flag.Parse()

	var modif float64
	if *up != 0 {
		modif = *up
	} else if *down != 0 {
		modif = -*down
	} else {
		c, err := bright.Current()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(c)
		return
	}
	err := bright.Modify(modif / 100)
	if err != nil {
		log.Fatal(err.Error())
	}
}
