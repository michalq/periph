// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package mcp9808_test

import (
	"fmt"
	"log"

	"github.com/michalq/periph/conn/i2c/i2creg"
	"github.com/michalq/periph/conn/physic"
	"github.com/michalq/periph/experimental/devices/mcp9808"
	"github.com/michalq/periph/host"
)

func ExampleDev_SenseTemp() {
	// Make sure periph is initialized.
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	// Open default I²C bus.
	bus, err := i2creg.Open("")
	if err != nil {
		log.Fatalf("failed to open I²C: %v", err)
	}
	defer bus.Close()

	// Create a new temperature sensor.
	sensor, err := mcp9808.New(bus, &mcp9808.DefaultOpts)
	if err != nil {
		log.Fatalln(err)
	}

	// Read values from sensor.
	measurement, err := sensor.SenseTemp()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(measurement)
}

func ExampleDev_SenseWithAlerts() {
	// Make sure periph is initialized.
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	// Open default I²C bus.
	bus, err := i2creg.Open("")
	if err != nil {
		log.Fatalf("failed to open I²C: %v", err)
	}
	defer bus.Close()

	// Create a new temperature sensor.
	sensor, err := mcp9808.New(bus, &mcp9808.DefaultOpts)
	if err != nil {
		log.Fatalln(err)
	}

	lower := physic.ZeroCelsius
	upper := physic.ZeroCelsius + 25*physic.Celsius
	critical := physic.ZeroCelsius + 32*physic.Celsius

	// Read values from sensor.
	temperature, alerts, err := sensor.SenseWithAlerts(lower, upper, critical)

	if err != nil {
		log.Fatalln(err)
	}

	for _, alert := range alerts {
		fmt.Println(alert)
	}

	fmt.Println(temperature)
}
