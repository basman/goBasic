package main

import (
	"fmt"

	"moonlander/physics"
)

// engineControlTick calls the engine controller to set the thrust level for the next tick
func engineControlTick(motion physics.Motion) float64 {
	// TODO implement engine control
	return 0.001
}

// engineControl is the feeding loop for the engine controller
func engineControl(motion chan physics.Motion, throttlePercentage chan float64) {
	for rocketState := range motion {
		throttlePercentage <- engineControlTick(rocketState)
	}
}

func main() {
	r := physics.NewRocket(2500, 4000, 1000)
	valve := make(chan float64)
	motion := make(chan physics.Motion)

	// run engineControl in separate goroutine to keep it from interfering with physical simulation timing
	go engineControl(motion, valve)
	done := r.Run(motion, valve)

	// wait for simulation to end
	<-done
	close(motion)

	m := r.GetMotion()
	if r.HasLanded() {
		fmt.Printf("Landing completed with velocity %.3f m/s and %.3f kg of remaining fuel\n", m.Velocity, m.Fuel)
	} else if r.HasCrashed() {
		fmt.Printf("Rocket has crashed with velocity %.3f m/s and %.3f kg of remaining fuel\n", m.Velocity, m.Fuel)
	} else {
		fmt.Printf("Unknown outcome. Simulation ended in dead state (velocity %.3f m/s, altitude %.3f m, remaining fuel %.3f kg)\n", m.Velocity, m.Altitude, m.Fuel)
	}
}
