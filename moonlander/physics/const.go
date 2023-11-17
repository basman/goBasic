package physics

/*
Physical and simulation constants
*/

// TimeTickMs defines the interval of simulation updates
const TimeTickMs = float64(50)

// MoonGravity defines the effective simulated gravity in m/s^2
const MoonGravity = float64(9.81) / 16

// MaxEngineThrustRatio is the ratio of the vertical engine acceleration in N at full throttle to the deadweight
const MaxEngineThrustRatio = float64(4.0)

// MaxFuelBurnRate in kg/s
const MaxFuelBurnRate = float64(2000) / 90 // 2 tons consumed in 90 seconds at full throttle

// MaxTouchDownVelocity is the maximum landing velocity to be regarded as a successful landing in m/s
const MaxTouchDownVelocity = float64(0.5)
