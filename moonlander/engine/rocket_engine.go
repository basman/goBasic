package engine

import "errors"

/*
Simulates a vertical rocket engine.
*/

type Engine struct {
	thrustValve float64
	isShutdown  bool
}

func NewEngine() *Engine {
	return &Engine{
		isShutdown:  true,
		thrustValve: 0,
	}
}

// Ignite starts the engine
func (e *Engine) Ignite(thrustValvePercentage float64) error {
	if e.IsRunning() {
		return errors.New("engine already running")
	}

	e.isShutdown = false
	e.thrustValve = thrustValvePercentage

	return nil
}

/*
SetThrustValve updates the setting of the thrust percentage (0-100)
*/
func (e *Engine) SetThrustValve(thrustValvePercentage float64) error {
	if thrustValvePercentage > 100 {
		return errors.New("valve setting exceeds maximum of 100%")
	} else if thrustValvePercentage < 0 {
		return errors.New("valve setting must not be negative")
	}

	if e.isShutdown {
		return errors.New("engine is shut down (flameout)")
	} else if thrustValvePercentage == 0 {
		e.isShutdown = true
	}

	e.thrustValve = thrustValvePercentage
	return nil
}

// GetThrustValve returns the current thrust opening in percentage
func (e *Engine) GetThrustValve() float64 {
	return e.thrustValve
}

/*
IsRunning indicates whether the engine is active and running (true) or
if it suffered from a flame-out (false)
and does not generate any thrust, regardless of the valve opening
*/
func (e *Engine) IsRunning() bool {
	return !e.isShutdown
}

func (e *Engine) RunOutOfFuel() {
	e.isShutdown = true
	e.thrustValve = 0
}
