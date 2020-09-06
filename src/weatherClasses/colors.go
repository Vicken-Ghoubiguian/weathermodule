package weatherClasses

// Definition of all colors used in the weatherModule package
/*var Red string = "\033[31m"
var Green string = "\033[32m"
var Cyan string = "\033[36m"
var Reset string = "\033[0m"*/

const (
	red = "\033[31m"
	green = "\033[32m"
	cyan = "\033[36m"
	reset = "\033[0m"
)

func Green() string {

	return green
}

func Red() string {

	return red
}

func Cyan() string {

	return cyan
}

func Reset() string {

	return reset
}
