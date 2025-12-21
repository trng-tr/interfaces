package main

type MotorVehicle interface {
	Mileage() float64
	PrintInfos()
}

type BMW struct {
	TraveledDistance float64
	FuelQuantityUsed float64
	AverageSpeed     string
}

type Audi struct {
	TraveledDistance float64
	FuelQuantityUsed float64
}
