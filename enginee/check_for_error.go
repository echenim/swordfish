package enginee

import "log"

//Check for error function
func Check(e error, message string) {
	if e != nil {
		log.Printf("%s %s\n", message, e)
	}
}
