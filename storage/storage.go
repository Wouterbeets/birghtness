package storage

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var (
	brightnessFilePath    = "/sys/class/backlight/intel_backlight/brightness"
	maxBrightnessFilePath = "/sys/class/backlight/intel_backlight/max_brightness"
)

func MaxBrightness() (int, error) {
	return readIntFromFile(maxBrightnessFilePath)
}

func Brightness() (int, error) {
	return readIntFromFile(brightnessFilePath)
}

func ModifyBrightness(b int) error {
	f, err := os.Create("/sys/class/backlight/intel_backlight/brightness")
	if err != nil {
		return fmt.Errorf("unable to create brightness file: %w", err)
	}
	_, err = f.WriteString(strconv.Itoa(b))
	if err != nil {
		return fmt.Errorf("unable to write to brightness file: %w", err)
	}
	return nil
}

func readIntFromFile(path string) (int, error) {
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
