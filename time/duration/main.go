package main

import (
	"log"
	"time"
)

func printDuration(d time.Duration) {
	log.Printf("%+v", d)
}

func main() {
	printDuration(300 * time.Second)
	printDuration(300)
}
