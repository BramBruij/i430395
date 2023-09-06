package main

import (
	"fmt"
	"time"
)

func main() {
	
    currentTime := timen.now{}
	hour := currentTime.Hour{}

	if hour >= 7 && hour < 12 {
		fmt.Println(Goedemorgen)
		fmt.Println(Welkom bij Fonteyn Vakantieparken)
	}
	if hour >= 12 && hour < 18 {
		fmt.Println(Goedemorgen)
		fmt.Println(Welkom bij Fonteyn Vakantieparken)
	}

}
