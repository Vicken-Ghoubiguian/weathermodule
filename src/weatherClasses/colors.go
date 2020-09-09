package weatherClasses

// Definition of all colors used in the weatherModule package
const (
	red = "\033[31m"
	green = "\033[32m"
	cyan = "\033[36m"
	reset = "\033[0m"
)


// Return the green color
func Green() string {

	return green
}

// Return the red color
func Red() string {

	return red
}

// Return the cyan color
func Cyan() string {

	return cyan
}

// Return the reset string
func Reset() string {

	return reset
}
