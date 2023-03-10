package options

import "errors"

const DefaultSpeed = 1

type SetSpeedError error

func SetSpeed(options *Options, speed int) SetSpeedError {
	if speed < 1 || speed > 6 {
		return SetSpeedError(errors.New("speed not valid (1-6)"))
	}

	return nil
}

func SetSpeedDefault(options *Options) {
	options.Speed = uint16(DefaultSpeed)
}
