package main

import "fmt"

func main() {
	//un BMW  est un MotorVehicle
	var bmw MotorVehicle = NewBMW(405.50, 28.50, "60 km/h")
	//un Audi  est un MotorVehicle
	var audi MotorVehicle = NewAudi(506.50, 46.50)
	bmw.PrintInfos()
	audi.PrintInfos()
	var vehicles []MotorVehicle = make([]MotorVehicle, 0)
	vehicles = append(vehicles, bmw, audi)
	var tm float64 = totalMileage(vehicles)
	fmt.Printf("total mileage %.2f km/l\n", tm)
}
