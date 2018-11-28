package accessory

import (
	"github.com/koenijn/hc/service"
)

type Thermostat struct {
	*Accessory

	Thermostat *service.Thermostat
}

// NewThermostat returns a Thermostat which implements model.Thermostat.
func NewThermostat(info Info, temp, min, max, steps float64) *Thermostat {
	acc := Thermostat{}
	acc.Accessory = New(info, TypeThermostat)
	acc.Thermostat = service.NewThermostat()
	acc.Thermostat.CurrentTemperature.SetValue(0)
	acc.Thermostat.CurrentTemperature.SetMinValue(0)
	acc.Thermostat.CurrentTemperature.SetMaxValue(100)
	acc.Thermostat.CurrentTemperature.SetStepValue(0.1)

	acc.Thermostat.TargetTemperature.SetValue(temp)
	acc.Thermostat.TargetTemperature.SetMinValue(min)
	acc.Thermostat.TargetTemperature.SetMaxValue(max)
	acc.Thermostat.TargetTemperature.SetStepValue(steps)

	acc.AddService(acc.Thermostat.Service)

	return &acc
}
