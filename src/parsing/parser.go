package parsing

import "fmt"

// Activate is a placeholder for browser input parsing.
func Activate(url chan string) {
	for val := range url {
		fmt.Println("URL is: ", val)
	}
}
