package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

var kentekenNummer string

func init() {
	flag.StringVar(&kentekenNummer, "kentekennummer", "", "kentekennummer specieficeren")
	flag.Parse()
	if kentekenNummer == "" {
		fmt.Println("Voer een kentekennummer in")
	}
}

func main() {

	currentTime := time.Now()
	dayPart := GetDayPart(currentTime)

	if dayPart == "Nacht" {
		fmt.Println("Wij zijn momenteel gesloten")
	}
	if dayPart == "Morgen" {
		fmt.Println("Welkom bij Fonteyn Vakantieparken")
	}
	if dayPart == "Middag" {
		fmt.Println("Welkom bij Fonteyn Vakantieparken")
	}
	if dayPart == "Avond" {
		fmt.Println("welkom bij Fonteyn Vakantieparken")
	}

	Boeking, err := CheckBoeking(kentekenNummer)
	if err != nil {
		fmt.Println("Sorry, wij zijn gesloten")
		log.Println(err)
		os.Exit(2)
	}
	fmt.Println("Goeden" + dayPart + "! Welcome" + Boeking.Name + " tot Fonteyn Vakantieparken!")
}

func CheckBoeking(kentekenNummer string) (Boeking, error) {
	var boeking Boeking
	boekingen := []Boeking{
		{Naam: "Bram", KentekenNummer: "1-AAA-11"},
		{Naam: "Peter", KentekenNummer: "1-BBB-11"},
	}

	for _, boeking := range boeking {
		if boeking.KentekenNummer == kentekenNummer {
			return boeking, nil
		}
	}

	return boeking, fmt.Errorf("no reservation found for %v", kentekenNummer)
}

func GetDayPart(time time.Time) string { Current
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

type Boeking struct {
	Name           string
	KentekenNummer string
}

type Vakantieparken struct{
	Name string
	Country string
	bookingen []Boeking
}