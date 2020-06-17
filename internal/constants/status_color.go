package constants

import "image/color"

type StatusColor struct {
	Color  color.RGBA
	Status string
}

var StatusCodeColorMap = map[string]StatusColor{
	"?": {
		Status: "NEW",
		Color:  color.RGBA{0, 128, 0, 1},
	},
	"M": {
		Status: "MODIFIED",
		Color:  color.RGBA{128, 128, 0, 1},
	},
	"A": {
		Status: "NEW",
		Color:  color.RGBA{0, 128, 0, 1},
	},
	"D": {
		Status: "DELETED",
		Color:  color.RGBA{128, 0, 0, 1},
	},
}
