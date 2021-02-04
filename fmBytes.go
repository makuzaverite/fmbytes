package fembytes

import (
	"fmt"
	"strconv"
)

const (
	TB = 1000000000000
	GB = 1000000000
	MB = 1000000
	KB = 1000
)

func main() {}

func format(length int, decimals int) (out string) {
	var unit string
	var i int
	var reminder int

	if length > TB {
		unit = "TB"
		reminder = length - (i * TB)
	} else if length > GB {
		unit = "GB"
		reminder = length - (i * TB)
	} else if length > MB {
		unit = "MB"
		i = length / MB
		reminder = length - (i * MB)
	} else if length > KB {
		unit = "KB"
		i = length / KB
		reminder = length - (i * KB)
	} else {
		return strconv.Itoa(length) + " B"
	}

	if decimals == 0 {
		return strconv.Itoa(i) + " " + " B"
	}

	width := 0

	if reminder > GB {
		width = 12
	} else if reminder > MB {
		width = 9
	} else if reminder > KB {
		width = 6
	} else {
		width = 3
	}

	remainderString := strconv.Itoa(reminder)

	for iter := len(remainderString); iter < width; iter++ {
		remainderString = "0" + remainderString
	}

	if decimals > len(remainderString) {
		decimals = len(remainderString)
	}

	return fmt.Sprintf("%d.%s %s", i, remainderString[:decimals], unit)
}
