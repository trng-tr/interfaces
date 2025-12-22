package main

import (
	"fmt"

	"github.com/trng-tr/interfaces/app"
	"github.com/trng-tr/interfaces/constructors"
)

func main() {
	//un BMW  est un MotorVehicle
	var bmw app.MotorVehicle = constructors.NewBMW(405.50, 28.50, "60 km/h")
	//un Audi  est un MotorVehicle
	var audi app.MotorVehicle = constructors.NewAudi(506.50, 46.50)
	bmw.PrintInfos()
	audi.PrintInfos()
	var vehicles []app.MotorVehicle = make([]app.MotorVehicle, 0)
	vehicles = append(vehicles, bmw, audi)
	fmt.Println("total mileage", app.TotalMileage(vehicles))
}
