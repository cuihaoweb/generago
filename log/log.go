package log

import "fmt"

// Error xx
func Error(message string, err error) {
	fmt.Print(ERROR + "\n")
	fmt.Print("Title:\t" + message + "\n")
	panic("Detail:\t" + err.Error() + "\n")
}
