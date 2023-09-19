package main

import (
	"fmt"
	"time"
)

func main() {
	
    currentTime := time.Now()
	dayPart := GetDayPart(currentTime)

	if dayPart == "Nacht" {
		fmt.Println("Wij zijn momenteel gesloten")
	}
	if dayPart == "Morgen"{
		fmt.Println("Welkom bij Fonteyn Vakantieparken")
	}
	if dayPart == "Middag"{
		fmt.Println("Welkom bij Fonteyn Vakantieparken")
	}
	if dayPart == "Avond"{
		fmt.Println("welkom bij Fonteyn Vakantieparken")
	}
}

func GetDayPart(time time.Time) string{
	currentHour := time.Hour()
	if currentHour >= 7 && currentHour < 12 {
		return "Morgen"
	}
	if currentHour >= 12 && currentHour < 18 {
		return "Middag"
	}
	if currentHour >= 18 && currentHour < 23 {
		return "Avond"
	}
	return "Nacht"
}