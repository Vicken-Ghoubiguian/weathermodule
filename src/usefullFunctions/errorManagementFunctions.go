package usefullFunctions

// Import all necessary packages
import (
	"fmt"
	"os"
)

// Function which display HTTP request error's code and message when the first occurs, with color and reset strings
func OwmErrorHandlerFunction(codeError string, color string, errorMessage string, reset string) {

        fmt.Println(color + "Occured error (" + codeError + "): " + errorMessage + reset)

        fmt.Println("\n")

        os.Exit(1)
}

// Function which display other errors when they occurs, with color and reset strings
func OtherErrorHandlerFunction(err error, color string, reset string) {

        if err != nil {

                fmt.Println(color + err.Error() + reset)

                os.Exit(1)
        }
}
