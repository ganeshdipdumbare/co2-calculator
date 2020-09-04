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
	"bytes"
	"ganeshdipdumbare/co2-calculator/model"
	"io/ioutil"
	"strings"
	"testing"
)

func Test_ExecuteCommand(t *testing.T) {
	testCmd := rootCmd

	// without any flag
	b := bytes.NewBufferString("")
	testCmd.SetOut(b)
	testCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	errFlagRequired := `Error: required flag(s)`
	if !strings.Contains(string(out), errFlagRequired) {
		t.Fatalf("expected %v ,got %v", errFlagRequired, string(out))
	}

	// without flagTransportationMethod flag
	testCmd.SetArgs([]string{"--" + flagDistance, "10"})
	testCmd.SetOut(b)
	testCmd.Execute()
	out, err = ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(out), errFlagRequired) {
		t.Fatalf("expected %v ,got %v", errFlagRequired, string(out))

	}

	// with valid data
	testCmd.SetArgs([]string{"--" + flagTransportationMethod, string(model.Train)})
	testCmd.SetOut(b)
	testCmd.Execute()
	out, err = ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	outputMsg := `0.1kg of CO2-equivalent`
	if !strings.Contains(string(out), outputMsg) {
		t.Fatalf("expected %v ,got %v", outputMsg, string(out))
	}

	// with valid data
	testCmd.SetArgs([]string{"--" + flagOutput, string(model.WeightUnitGram)})
	testCmd.SetOut(b)
	testCmd.Execute()
	out, err = ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	outputMsg = `60.0g of CO2-equivalent`
	if !strings.Contains(string(out), outputMsg) {
		t.Fatalf("expected %v ,got %v", outputMsg, string(out))

	}
}
