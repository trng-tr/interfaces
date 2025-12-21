package main

func NewBMW(traveledDistance float64, fuelQuantityUsed float64, averageSpeed string) BMW {
	return BMW{
		TraveledDistance: traveledDistance,
		FuelQuantityUsed: fuelQuantityUsed,
		AverageSpeed:     averageSpeed,
	}
}

func NewAudi(traveledDistance float64, fuelQuantityUsed float64) Audi {
	return Audi{
		TraveledDistance: traveledDistance,
		FuelQuantityUsed: fuelQuantityUsed,
	}
}
