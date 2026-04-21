// maps.go

package main

import (
	"fmt"
)

func main() {
	var statusMessages = map[int]string{
		200: "OK: The request was successful.",
		404: "Not Found: The server cannot find the requested resource.",
		500: "Internal Server Error: The server encountered an unexpected condition.",
		503: "Service Unavailable: The server is currently unable to handle the request.",
	}

	// example
	code := 404

	if message, ok := statusMessages[code]; ok {
		fmt.Printf("Status %d: %s\n", code, message)
	} else {
		fmt.Println("Status code not found in our records.")
	}
}
