package main

import (
	"fmt"
	"os"
	OmniVision "player/OmniVision_pkg"
	"strconv"
	"sync"
)

// WaitGroup is used to wait for the program to finish goroutines.
var WG sync.WaitGroup
var goroutines_nmr = 3

func main() {
	fmt.Println("START Player")
	// Add a count of three, one for each goroutine.
	WG.Add(goroutines_nmr)
	
	// Other goroutines
	// (...)
	
	my_id := os.Args[3]
	my_grpc_ip := os.Args[6]
	shirt, _ := strconv.Atoi(os.Args[7]) //0 red_shirt 1 blue shirt
	ID, _ := strconv.Atoi(my_id)
	
	// # Start OmniVision goroutine
	go OmniVision.Run(shirt == 1, my_grpc_ip, ID)


	// Wait for the goroutines to finish.
	WG.Wait()
}
