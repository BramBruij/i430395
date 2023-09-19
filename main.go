package main

import (
	"fmt"
	"time"
)

func main() {
	
    currentTime := timen.now{}
	dayPart := GetDayPart(currentTime)


	if dayPart == "Nacht" {
		fmt.Println(Wij zijn momenteel gesloten)
	}
	if dayPart == "Morgen"{
		fmt.Println(Welkom bij Fonteyn Vakantieparken)
	}
	if dayPart == "Middag"{
		fmt.Println(Welkom bij Fonteyn Vakantieparken)
	}
	if dayPart == "Avond"{
		fmt.Println(welkom bij Fonteyn Vakantieparken)
	}
}

func GetDayPart(time.Time) string{
	currentHour := time.Hour()
	if currentHour >= 7 && < 12 {
		return "Morgen"
	}
	if currentHour >= 12 && < 18 {
		return "Middag"
	}
	if currentHour >= 18 && < 23 {
		return "Avond"
	}
	return "Nacht"
}