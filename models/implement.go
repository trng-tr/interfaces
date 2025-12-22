package models

import "fmt"

// la structure BMW implémente (implicitement) l'interface MotorVehicle👇
func (bmw BMW) Mileage() float64 {
	if bmw.FuelQuantityUsed == 0 {
		return -1
	}
	return bmw.TraveledDistance / bmw.FuelQuantityUsed
}
func (bmw BMW) PrintInfos() {
	fmt.Println("bmw motor:", bmw)
}

// la structure Audi implémente (implicitement) l'interface MotorVehicle
func (audi Audi) Mileage() float64 {
	if audi.FuelQuantityUsed == 0 {
		return -1
	}
	return audi.TraveledDistance / audi.FuelQuantityUsed
}
func (audi Audi) PrintInfos() {
	fmt.Println("audi motor:", audi)
}
