package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/wouterbeets/brightness/storage"

)

func readFile(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	br, err := ioutil.ReadAll(f)
	if err != nil {
		return 0, err
	}

	currentStr := strings.Trim(string(br), "\n")
	current, err := strconv.Atoi(currentStr)
	if err != nil {
		return 0, err
	}
	return current, nil
}

func main() {
	up := flag.Int("up", 0, "up brightness")
	down := flag.Int("down", 0, "down brightness")
	flag.Parse()

	var modif int
	if *up != 0 {
		modif = *up
	} else if *down != 0 {
		modif = -*down
	}

	max, err := storage.MaxBrightness()
	if err != nil {
		log.Fatal(err)
	}
	brightness, err := storage.MaxBrightness()
	if err != nil {
		log.Fatal(err)
	}

	var currentPercentage float64
	if modif == 0 {
		currentPercentage = float64(brightness) / float64(max)
		fmt.Printf("%.2f\n", currentPercentage)
		return
	}

	newBrightness := brightness + modif
	if newBrightness > max {
		newBrightness = max
	}

	if newBrightness < 1000 {
		newBrightness = 1000
	}

	f, err := os.Create("/sys/class/backlight/intel_backlight/brightness")
	if err != nil {
		panic(err)
	}
	_, err = f.WriteString(strconv.Itoa(newBrightness))
	if err != nil {
		panic(err)
	}

}
