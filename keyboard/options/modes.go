package options

import "errors"

const DefaultModeId = 10

var modes = []*Mode{
	&Mode{
		ID:   uint16(1),
		Name: "Simple color change with dark",
		Icon: "Flame",
	},
	&Mode{
		ID:   uint16(2),
		Name: "Rainbow",
		Icon: "Glasses",
	},
	&Mode{
		ID:   uint16(3),
		Name: "Simple color change smooth",
		Icon: "Flower",
	},
	&Mode{
		ID:   uint16(4),
		Name: "Circle",
		Icon: "RefreshCircle",
	},
	&Mode{
		ID:   uint16(5),
		Name: "Wtf",
		Icon: "CogSharp",
	},
	&Mode{
		ID:   uint16(6),
		Name: "Click color",
		Icon: "HandLeft",
	},
	&Mode{
		ID:   uint16(7),
		Name: "Click line",
		Icon: "ReorderFourOutline",
	},
	&Mode{
		ID:   uint16(8),
		Name: "Click feel",
		Icon: "WaterOutline",
	},
	&Mode{
		ID:   uint16(9),
		Name: "Random color (?)",
		Icon: "TelescopeOutline",
	},
	&Mode{
		ID:   uint16(10),
		Name: "Manual",
		Icon: "Settings",
	},
	&Mode{
		ID:   uint16(11),
		Name: "Click plus",
		Icon: "Add",
	},
	&Mode{
		ID:   uint16(12),
		Name: "Equalizer clicked",
		Icon: "AlertSharp",
	},
	&Mode{
		ID:   uint16(13),
		Name: "Equalizer clicked 2 (??)",
		Icon: "AirplaneSharp",
	},
	&Mode{
		ID:   uint16(14),
		Name: "Auto diagonal feel",
		Icon: "BanOutline",
	},
	&Mode{
		ID:   uint16(15),
		Name: "Auto lines feel",
		Icon: "Bandage",
	},
	&Mode{
		ID:   uint16(16),
		Name: "Click circle",
		Icon: "At",
	},
}

type SetModeError error

type Mode struct {
	ID   uint16 `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

func GetModeList() []*Mode {
	return modes
}

func SetModeWithName(options *Options, name string) SetModeError {
	for _, m := range modes {
		if m.Name == name {
			options.Mode = m.ID
			return nil
		}
	}

	return SetModeError(errors.New("mode not found"))
}

func SetModeWithID(options *Options, id uint) SetModeError {
	for _, m := range modes {
		uID := uint16(id)
		if m.ID == uID {
			options.Mode = m.ID
			return nil
		}
	}

	return SetModeError(errors.New("mode not found"))
}

func SetModeDefault(options *Options) {
	options.Mode = uint16(DefaultModeId)
}
