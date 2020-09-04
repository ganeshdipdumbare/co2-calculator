package app

import (
	"ganeshdipdumbare/co2-calculator/model"
	"reflect"
	"testing"
)

func TestNewApp(t *testing.T) {
	tests := []struct {
		name string
		want CO2App
	}{
		{
			name: "should return valid co2 app",
			want: NewApp(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewApp(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewApp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_app_GetCO2AmoutForJourney(t *testing.T) {
	emissionMap := map[model.TransportationMethod]float64{
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
	}

	var res float64 = 0.00027
	type fields struct {
		co2emissionForTransport map[model.TransportationMethod]float64
	}
	type args struct {
		j *model.Journey
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *float64
		wantErr bool
	}{
		{
			name: "should return error for invalid args",
			args: args{
				j: &model.Journey{
					Distance:             10,
					OutputUnit:           "invalid",
					TransportationMethod: model.Bus,
					UnitOfDistance:       model.DistanceUnitKilometer,
				},
			},
			fields: fields{
				co2emissionForTransport: emissionMap,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should return success for valid args",
			args: args{
				j: &model.Journey{
					Distance:             10,
					OutputUnit:           model.WeightUnitKilogram,
					TransportationMethod: model.Bus,
					UnitOfDistance:       model.DistanceUnitMeter,
				},
			},
			fields: fields{
				co2emissionForTransport: emissionMap,
			},
			want:    &res,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &app{
				co2emissionForTransport: tt.fields.co2emissionForTransport,
			}
			got, err := a.GetCO2AmoutForJourney(tt.args.j)
			if (err != nil) != tt.wantErr {
				t.Errorf("app.GetCO2AmoutForJourney() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("app.GetCO2AmoutForJourney() = %v, want %v", *got, *tt.want)
			}
		})
	}
}

func Test_validateGetCO2AmoutInputs(t *testing.T) {
	type args struct {
		j *model.Journey
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "should return success for valid inputs",
			args: args{
				j: &model.Journey{
					Distance:             10,
					OutputUnit:           model.WeightUnitGram,
					TransportationMethod: model.Bus,
					UnitOfDistance:       model.DistanceUnitKilometer,
				},
			},
			wantErr: false,
		},
		{
			name: "should return error for invalid OutputUnit",
			args: args{
				j: &model.Journey{
					Distance:             10,
					OutputUnit:           "invalid",
					TransportationMethod: model.Bus,
					UnitOfDistance:       model.DistanceUnitKilometer,
				},
			},
			wantErr: true,
		},
		{
			name: "should return error for invalid TransportationMethod",
			args: args{
				j: &model.Journey{
					Distance:             10,
					OutputUnit:           model.WeightUnitGram,
					TransportationMethod: "invalid",
					UnitOfDistance:       model.DistanceUnitKilometer,
				},
			},
			wantErr: true,
		},
		{
			name: "should return error for invalid UnitOfDistance",
			args: args{
				j: &model.Journey{
					Distance:             10,
					OutputUnit:           model.WeightUnitGram,
					TransportationMethod: model.Bus,
					UnitOfDistance:       "invalid",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateGetCO2AmoutInputs(tt.args.j); (err != nil) != tt.wantErr {
				t.Errorf("validateGetCO2AmoutInputs() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
