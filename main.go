package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	up := flag.Int("up", 0, "up birghtness")
	down := flag.Int("down", 0, "up birghtness")
	flag.Parse()
	var modif int
	if *up != 0 {
		modif = *up
	} else if *down != 0 {
		modif = -*down
	} else {
		return
	}

	fmt.Println("modif", modif)
	f, err := os.Open("/sys/class/backlight/intel_backlight/brightness")
	if err != nil {
		panic(err.Error())
	}
	br, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err.Error())
	}

	currentStr := strings.Trim(string(br), "\n")
	current, err := strconv.Atoi(currentStr)
	if err != nil {
		panic(err)
	}
	fmt.Println(current + modif)
	f, err = os.Create("/sys/class/backlight/intel_backlight/brightness")
	if err != nil {
		panic(err)
	}
	_, err = f.WriteString(strconv.Itoa(current + modif))
	if err != nil {
		panic(err)
	}
}
