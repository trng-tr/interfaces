package constructors

import "github.com/trng-tr/interfaces/models"

func NewBMW(traveledDistance float64, fuelQuantityUsed float64, averageSpeed string) models.BMW {
	return models.BMW{
		TraveledDistance: traveledDistance,
		FuelQuantityUsed: fuelQuantityUsed,
		AverageSpeed:     averageSpeed,
	}
}

func NewAudi(traveledDistance float64, fuelQuantityUsed float64) models.Audi {
	return models.Audi{
		TraveledDistance: traveledDistance,
		FuelQuantityUsed: fuelQuantityUsed,
	}
}
