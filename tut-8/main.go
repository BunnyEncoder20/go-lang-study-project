package main

import "fmt"

func main() {
	intSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("Sum of int elements:", sumSlice(intSlice))

	float32Slice := []float32{10, 20, 30}
	fmt.Println("Sum of float32 elements:", sumSlice(float32Slice))

	float64Slice := []float64{1.5, 2.5, 3.5}
	fmt.Println("Sum of float64 elements:", sumSlice(float64Slice))

	stringSlice := []string{}
	fmt.Println("Is stringSlice empty?", isEmpty(stringSlice))

	// using generics with structs
	gasCar := car[gasEngine]{
		company: "Toyota",
		model:   "Corolla",
		engine: gasEngine{
			liters: 1.8,
			kmph:   15.0,
		},
	}

	electricCar := car[electricEngine]{
		company: "Tesla",
		model:   "Model 3",
		engine: electricEngine{
			kWh:    75,
			kmpkWh: 5.0,
		},
	}

	fmt.Println("Gas Car:\n", gasCar)
	fmt.Println("Electric Car:\n", electricCar)
}

type gasEngine struct {
	liters float32
	kmph   float32
}

type electricEngine struct {
	kWh    float32
	kmpkWh float32
}

type car[T gasEngine | electricEngine] struct {
	company string
	model   string
	engine  T
}

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for v := range slice {
		sum += slice[v]
	}
	return sum
}

func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}
