package model

// TransportationMethod defines type for transportation method
type TransportationMethod string

// DistanceUnit defines type for distance unit
type DistanceUnit string

// WeightUnit defines type for weight unit
type WeightUnit string

const (
	SmallDieselCar        TransportationMethod = "small-diesel-car"
	SmallPetrolCar        TransportationMethod = "small-petrol-car"
	SmallPluginHybridCar  TransportationMethod = "small-plugin-hybrid-car"
	SmallElectricCar      TransportationMethod = "small-electric-car"
	MediumDieselCar       TransportationMethod = "medium-diesel-car"
	MediumPetrolCar       TransportationMethod = "medium-petrol-car"
	MediumPluginHybridCar TransportationMethod = "medium-plugin-hybrid-car"
	MediumElectricCar     TransportationMethod = "medium-electric-car"
	LargeDieselCar        TransportationMethod = "large-diesel-car"
	LargePetrolCar        TransportationMethod = "large-petrol-car"
	LargePluginHybridCar  TransportationMethod = "large-plugin-hybrid-car"
	LargeElectricCar      TransportationMethod = "large-electric-car"
	Bus                   TransportationMethod = "bus"
	Train                 TransportationMethod = "train"

	DistanceUnitMeter     DistanceUnit = "m"
	DistanceUnitKilometer DistanceUnit = "km"

	WeightUnitGram     WeightUnit = "g"
	WeightUnitKilogram WeightUnit = "kg"
)

// Journey a struct to hold journey related details
type Journey struct {
	Distance             float64
	TransportationMethod TransportationMethod
	UnitOfDistance       DistanceUnit
	OutputUnit           WeightUnit
}
