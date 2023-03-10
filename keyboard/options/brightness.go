package options

import "errors"

type SetBrightnessError error

const (
	BrightnessOff    = 1
	BrightnessLow    = 2
	BrightnessMedium = 3
	BrightnessHigh   = 4
)

func SetBrightness(options *Options, brightness int) SetBrightnessError {
	if brightness < 1 || brightness > 4 {
		return SetBrightnessError(errors.New("brightness not valid"))
	}

	options.Brightness = uint16(brightness)
	return nil
}
