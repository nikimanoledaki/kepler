/*
Copyright 2021.

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

package power

import (
	"fmt"
)

type powerInterface interface {
	GetEnergyFromDram() (uint64, error)
	GetEnergyFromCore() (uint64, error)
	GetEnergyFromUncore() (uint64, error)
	GetEnergyFromPackage() (uint64, error)
	StopPower()
	IsSupported() bool
}

var (
	dummyImpl                = &powerDummy{}
	sysfsImpl                = &raplSysfs{}
	msrImpl                  = &raplMSR{}
	powerImpl powerInterface = sysfsImpl
)

func init() {
	if sysfsImpl.IsSupported() /*&& false */ {
		fmt.Println("use sysfs to obtain power")
		powerImpl = sysfsImpl
	} else {
		if msrImpl.IsSupported() {
			fmt.Println("use MSR to obtain power")
			powerImpl = msrImpl
		} else {
			fmt.Println("power not supported")
			powerImpl = dummyImpl
		}
	}
}

func GetEnergyFromDram() (uint64, error) {
	return powerImpl.GetEnergyFromDram()
}

func GetEnergyFromCore() (uint64, error) {
	return powerImpl.GetEnergyFromCore()
}

func GetEnergyFromUncore() (uint64, error) {
	return powerImpl.GetEnergyFromUncore()
}

func GetEnergyFromPackage() (uint64, error) {
	return powerImpl.GetEnergyFromPackage()
}

func StopPower() {
	powerImpl.StopPower()
}