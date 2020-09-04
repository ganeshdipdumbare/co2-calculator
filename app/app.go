package app

import (
	"fmt"
	"ganeshdipdumbare/co2-calculator/model"
)

// CO2App defines methods to handle co2 emission related operations
type CO2App interface {
	GetCO2AmoutForJourney(j *model.Journey) (*float64, error)
}

type app struct {
	co2emissionForTransport map[model.TransportationMethod]float64
}

// NewApp creates new CO2App instance
func NewApp() CO2App {
	return &app{
		co2emissionForTransport: map[model.TransportationMethod]float64{
			model.SmallDieselCar:        142,
			model.SmallPetrolCar:        154,
			model.SmallPluginHybridCar:  73,
			model.SmallElectricCar:      50,
			model.MediumDieselCar:       171,
			model.MediumPetrolCar:       50,
			model.MediumPluginHybridCar: 110,
			model.MediumElectricCar:     58,
			model.LargeDieselCar:        209,
			model.LargePetrolCar:        282,
			model.LargePluginHybridCar:  126,
			model.LargeElectricCar:      73,
			model.Bus:                   27,
			model.Train:                 6,
		},
	}
}

// GetCO2AmoutForJourney returns CO2 emission for given journey if inputs are valid,
// otherwise returns error
func (a *app) GetCO2AmoutForJourney(j *model.Journey) (*float64, error) {
	err := validateGetCO2AmoutInputs(j)
	if err != nil {
		return nil, err
	}

	distance := j.Distance
	if j.UnitOfDistance == model.DistanceUnitMeter {
		distance = distance / 1000
	}

	co2Amout := distance * a.co2emissionForTransport[j.TransportationMethod]
	if j.OutputUnit == model.WeightUnitKilogram {
		co2Amout = co2Amout / 1000
	}
	return &co2Amout, nil
}

func validateGetCO2AmoutInputs(j *model.Journey) error {
	invalidInputErr := "%v field is invalid"

	// validate output unit
	switch j.OutputUnit {
	case model.WeightUnitGram, model.WeightUnitKilogram:
	default:
		return fmt.Errorf(invalidInputErr, "OutputUnit")
	}

	// validate unit of distance
	switch j.UnitOfDistance {
	case model.DistanceUnitKilometer, model.DistanceUnitMeter:
	default:
		return fmt.Errorf(invalidInputErr, "UnitOfDistance")

	}

	// validate transportation method
	switch j.TransportationMethod {
	case model.Train, model.Bus, model.LargeDieselCar, model.LargeElectricCar, model.LargePetrolCar,
		model.LargePluginHybridCar, model.MediumDieselCar, model.MediumElectricCar, model.MediumPetrolCar,
		model.MediumPluginHybridCar, model.SmallDieselCar, model.SmallElectricCar, model.SmallPetrolCar,
		model.SmallPluginHybridCar:
	default:
		return fmt.Errorf(invalidInputErr, "TransportationMethod")
	}

	return nil
}
