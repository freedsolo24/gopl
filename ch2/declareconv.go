package main

import (
	"fmt"
)

type Celsius float32
type Fahrenheit float32

type Meter float32
type Mile float32

type Pound float32
type Kilogram float32

// const (
// 	AbsoluteZeroC Celsius = -273.15
// 	FreezingC     Celsius = 0
// 	BoilingC      Celsius = 100
// )

func (c Celsius) String() string    { return fmt.Sprintf("%.3f°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%.3f°F", f) }

func (m Meter) String() string { return fmt.Sprintf("%.3fm", m) }
func (mi Mile) String() string { return fmt.Sprintf("%.3fmi", mi) }

func (p Pound) String() string    { return fmt.Sprintf("%.3flb", p) }
func (k Kilogram) String() string { return fmt.Sprintf("%.3fkg", k) }
