package physics

import "fmt"

type Motion struct {
	TickSerial   int64   // simulation tick's serial number
	Timestamp    float64 // seconds elapsed since start of simulation (microseconds precision)
	Velocity     float64
	Acceleration float64
	Fuel         float64
	Altitude     float64
	TotalWeight  float64
}

func (m Motion) String() string {
	return fmt.Sprintf("#%v %.6f [%.3fm %.3fm/s %.3fm/s2 %.3fkg]", m.TickSerial, m.Timestamp, m.Altitude, m.Velocity, m.Acceleration, m.Fuel)
}
