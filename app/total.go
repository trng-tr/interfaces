package app

/*
lorsque les structures BMW et Audi implementent l'interface  MotorVehicle
c'est à dire BMW et Audi vont redéfinir les methodes de MotorVehicle ,
un BMW sera un MotorVehicle et un Audi est un MotorVehicle
*/
func TotalMileage(m []MotorVehicle) float64 {
	var tm float64
	for _, v := range m { //
		tm += v.Mileage()
	}
	return tm
}
