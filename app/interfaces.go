package app

type MotorVehicle interface {
	Mileage() float64
	PrintInfos()
}

/*
avec cette implementation plus haut de MotorVehicle par les structures BMW et Audi,
un BMW est un MotorVehicle et les AUdi est un MotorVehicle
*/
func TotalMileage(m []MotorVehicle) float64 {
	var tm float64
	for _, v := range m { //
		tm += v.Mileage()
	}
	return tm
}
