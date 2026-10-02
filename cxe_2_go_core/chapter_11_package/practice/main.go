package main

import (
	"fmt"
	"math/rand/v2"
	"practice/booking"
	"practice/room"
)

const rooms = 134
var availableRooms = rand.IntN(10)

func occupancyRate(availableRooms int) int {
	return (rooms - availableRooms) * 100 / rooms
}

func occupancyLevel(occupancyRate int) string {
	switch {
	case occupancyRate < 30:	return "Low"
	case occupancyRate < 60:	return "Medium"
	default:		return "High"
	}
}

func printRoomDetails(availableRooms int) {
	if availableRooms == 0 {
		fmt.Printf("No rooms available for tonight\n")
	} else {
		fmt.Printf("Rooms:\n")
		for i := 110; i < 110 + availableRooms; i++ {
			people := rand.IntN(10) + 1
			nights := rand.IntN(10) + 1
			room.PrintDetails(i, people, nights)
		}
	}
}
func main() {
	occupancyRate := occupancyRate(availableRooms)
	occupancyLevel := occupancyLevel(occupancyRate)
	fmt.Printf("Hotel: California\n")
	fmt.Printf("Number of rooms: %d\n", rooms)
	fmt.Printf("Available rooms: %d\n", availableRooms)
	fmt.Printf("Occupancy level: %s\n", occupancyLevel)
	fmt.Printf("Occupancy rate: %d%%\n", occupancyRate)
	printRoomDetails(availableRooms)
	booking.PrintVatRate()
}