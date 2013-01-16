package term

type Color string

const (
	Reset      = Color("\x1b[0m")
	Bright     = Color("\x1b[1m")
	Dim        = Color("\x1b[2m")
	Underscore = Color("\x1b[4m")
	Blink      = Color("\x1b[5m")
	Reverse    = Color("\x1b[7m")
	Hidden     = Color("\x1b[8m")
	Black      = Color("\x1b[30m")
	Red        = Color("\x1b[31m")
	Green      = Color("\x1b[32m")
	Yellow     = Color("\x1b[33m")
	Blue       = Color("\x1b[34m")
	Magenta    = Color("\x1b[35m")
	Cyan       = Color("\x1b[36m")
	White      = Color("\x1b[37m")
)
