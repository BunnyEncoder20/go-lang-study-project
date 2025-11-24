package main

import "fmt"

type gasEngine struct {
	kmph   uint8
	liters uint8
	company
}

// Functions assigned to specific struct (gasEngine struct here)
func (e gasEngine) kmsLeft() uint8 {
	return e.kmph * e.liters
}

type electricEngine struct {
	kmpkWh uint8
	kWh    uint8
	company
}

func (e electricEngine) kmsLeft() uint8 {
	return e.kmpkWh * e.kWh
}

type company struct {
	name string
}

// Interface
type engine interface {
	kmsLeft() uint8
}

// func to see if you can make it to your desitnation
func canMakeIt(kmsToGo uint8, e engine) bool {
	if kmsToGo <= e.kmsLeft() {
		fmt.Println("you can make it !")
	} else {
		fmt.Println("you canNOT make it !")
	}
	return false
}

func main() {
	myEngine := gasEngine{25, 13, company{"Toyota"}}
	fmt.Printf("Mileage: %v\n Tank: %v\n Company: %v\n", myEngine.kmph, myEngine.liters, myEngine.name)

	myElectric := electricEngine{25, 40, company{"Tesla"}}
	fmt.Printf("Mileage: %v\n Tank: %v\n Company: %v\n", myElectric.kmpkWh, myElectric.kWh, myElectric.name)

	println("Can the gasEngine make it?")
	canMakeIt(200, myEngine)
	println("Can the electricEngine make it?")
	canMakeIt(200, myElectric)

	// Anonymus Stucts (no name, defined and init at same time)
	myEngine2 := struct {
		kmph   uint8
		liters uint8
	}{25, 13}
	fmt.Printf("Mileage: %v\n Tank: %v\n", myEngine2.kmph, myEngine2.liters)
}
