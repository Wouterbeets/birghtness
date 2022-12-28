package bright

import (
	"fmt"

	"github.com/wouterbeets/brightness/storage"
)

type Brightness float64

func (b Brightness) String() string {
	return fmt.Sprintf("%.2f", b)
}

func validate(b Brightness) Brightness {
	if b > 1 {
		return 1
	}
	if b < 0.05 {
		return 0.05
	}
	return b
}

func Set(b Brightness) error {
	b = validate(b)
	max, err := storage.MaxBrightness()
	if err != nil {
		return err
	}

	return storage.ModifyBrightness(int(float64(max) * float64(b)))
}

func Modify(percentage float64) error {
	c, err := Current()
	if err != nil {
		return err
	}
	return Set(c + Brightness(percentage))
}

func Current() (Brightness, error) {
	max, err := storage.MaxBrightness()
	if err != nil {
		return 0, err
	}
	current, err := storage.Brightness()
	if err != nil {
		return 0, err
	}
	return Brightness(float64(current) / float64(max)), nil
}
