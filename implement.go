package main

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

/*avec cette implementation plus haut de MotorVehicle par les structures BMW et Audi,
un BMW est un MotorVehicle et les AUdi est un MotorVehicle*/
func totalMileage(m []MotorVehicle) float64 {
	var tm float64
	for _, v := range m { //
		tm += v.Mileage()
	}
	return tm
}
