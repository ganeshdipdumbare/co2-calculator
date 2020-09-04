/*
Copyright © 2020 Ganeshdip <ganeshdip.dumbare@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"ganeshdipdumbare/co2-calculator/app"
	"ganeshdipdumbare/co2-calculator/model"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	flagDistance             = "distance"
	flagUnitOfDistance       = "unit-of-distance"
	flagTransportationMethod = "transportation-method"
	flagOutput               = "output"

	defaultUnitOfDistance = string(model.DistanceUnitKilometer)
	defaultOutputUnit     = string(model.WeightUnitKilogram)
	floatBase64           = 64
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "co2-calculator",
	Short: "Returns co2 emission amount for the iput journey parameters",
	Long:  `The command returns the CO2 emission amount for the given journey parameters`,

	Run: func(cmd *cobra.Command, args []string) {
		journey := &model.Journey{}
		cmd.Flags().VisitAll(func(f *pflag.Flag) {
			err := setJourneyParam(journey, f)
			if err != nil {
				cmd.Println("error occurred: ", err)
				return
			}
		})

		co2amount, err := co2app.GetCO2AmoutForJourney(journey)
		if err != nil {
			cmd.Println("error occurred: ", err)
			return
		}

		cmd.Printf("Your trip caused %.1f%v of CO2-equivalent.\n", *co2amount, journey.OutputUnit)
	},
}

var co2app app.CO2App

func setJourneyParam(j *model.Journey, f *pflag.Flag) error {
	value := f.Value.String()

	switch f.Name {
	case flagDistance:
		distance, err := strconv.ParseFloat(value, floatBase64)
		if err != nil {
			return err
		}
		j.Distance = distance
	case flagOutput:
		j.OutputUnit = model.WeightUnit(value)
	case flagTransportationMethod:
		j.TransportationMethod = model.TransportationMethod(value)
	case flagUnitOfDistance:
		j.UnitOfDistance = model.DistanceUnit(value)
	}

	return nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	co2app = app.NewApp()
	rootCmd.Flags().Float64(flagDistance, 0, "journey distance")
	rootCmd.Flags().String(flagUnitOfDistance, defaultUnitOfDistance, "journey distance unit")
	rootCmd.Flags().String(flagTransportationMethod, "", "transportation used for the journey")
	rootCmd.Flags().String(flagOutput, defaultOutputUnit, "co2 emission volume unit for output")

	errInvalidFlagName := "invalid flag name for marking required: %v\n"
	err := rootCmd.MarkFlagRequired(flagDistance)
	if err != nil {
		log.Fatalf(errInvalidFlagName, flagDistance)
	}

	err = rootCmd.MarkFlagRequired(flagTransportationMethod)
	if err != nil {
		log.Fatalf(errInvalidFlagName, flagTransportationMethod)
	}
}
