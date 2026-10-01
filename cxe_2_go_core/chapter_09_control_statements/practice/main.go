package main

import (
	"fmt"
	"math/rand/v2"
)

const rooms = 134
var availableRooms = rand.IntN(135)
func main() {
	occupancyRate := (rooms - availableRooms) * 100 / rooms
	occupancyLevel := ""
	switch {
	case occupancyRate < 30:
		occupancyLevel = "Low"
	case occupancyRate < 60:
		occupancyLevel = "Medium"
	default:
		occupancyLevel = "High"
	}
	fmt.Printf("Hotel: California\n")
	fmt.Printf("Number of rooms: %d\n", rooms)
	fmt.Printf("Available rooms: %d\n", availableRooms)
	fmt.Printf("Occupancy level: %s\n", occupancyLevel)
	fmt.Printf("Occupancy rate: %d%%\n", occupancyRate)
	if availableRooms == 0 {
		fmt.Printf("No rooms available for tonight\n")
	} else {
		fmt.Printf("Rooms:\n")
		for i := 110; i < 110 + availableRooms; i++ {
			people := rand.IntN(10) + 1
			nights := rand.IntN(10) + 1
			fmt.Printf("- %d: %d people / %d nights\n", i, people, nights)
		}
	}
}