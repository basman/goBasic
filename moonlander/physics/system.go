package physics

import (
	"fmt"
	"log"
	"sync"
	"time"

	"moonlander/engine"
)

type Rocket struct {
	tickSerial   int64     // serial number of motion information
	startTime    time.Time // start of simulation
	altitude     float64   // altitude in m
	deadWeight   float64   // dead weight in kg
	fuel         float64   // remaining fuel in kg
	velocity     float64   // vertical velocity in m/s
	acceleration float64   // vertical acceleration in m/s^2
	engine       *engine.Engine
	mu           *sync.RWMutex
}

func NewRocket(weight, fuel, altitude float64) *Rocket {
	return &Rocket{
		mu:           &sync.RWMutex{},
		tickSerial:   0,
		altitude:     altitude,
		deadWeight:   weight,
		fuel:         fuel,
		velocity:     0,
		acceleration: 0,
		engine:       engine.NewEngine(),
	}
}

func (r *Rocket) UpdateThrottle(throttle float64) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.engine.SetThrustValve(throttle)
}

/*
	Tick updates the physical and kinetic state of the rocket.

It uses the current engine thrust setting to apply the effective thrust to the rocket.
Motion vectors and fuel state are then being calculated and updated.

If the engine is not producing any thrust, only gravity and weight will be taken into account.
*/
func (r *Rocket) Tick() error {
	now := time.Now()

	r.mu.Lock()
	defer r.mu.Unlock()

	// observe lock contention closely
	delay := time.Now().Sub(now).Microseconds()
	if float64(delay) > tickDelayLimitMs*1e3 {
		log.Printf("WARNING: possible lock contention. Time tick delay %v microseconds\n", delay)
	}

	if r.tickSerial == 0 {
		r.startTime = time.Now()
	}
	r.tickSerial++

	// We would burn as much during one time tick at current throttle. But do we have as much fuel?
	deltaFuel := MaxFuelBurnRate * TimeTickMs / 1000 * r.engine.GetThrustValve() / 100

	if r.fuel <= 0 {
		r.engine.RunOutOfFuel()
		deltaFuel = 0
	}

	burnTime := TimeTickMs / 1000

	// burn fuel (=reduce fuel)
	if deltaFuel > r.fuel {
		// we run out of fuel this very time tick - reduce burn time
		burnTime = r.fuel / (MaxFuelBurnRate * r.engine.GetThrustValve() / 100)
		deltaFuel = r.fuel
	}

	// engine acceleration m/s^2
	aEn := r.engine.GetThrustValve() / 100 * MaxEngineThrustRatio * r.deadWeight / (r.deadWeight + r.fuel)

	// gravity acceleration
	aGrav := MoonGravity

	// effective acceleration per second
	// velocity change during one time tick
	deltaV := aEn*burnTime - aGrav*TimeTickMs/1000
	r.velocity += deltaV
	r.acceleration = deltaV / (TimeTickMs / 1000)
	r.altitude += r.velocity * (TimeTickMs / 1000)

	// update fuel level and total weight
	r.fuel -= deltaFuel

	return nil
}

func (r *Rocket) HasLanded() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if r.altitude <= 0 && r.velocity <= 0 && r.velocity >= -MaxTouchDownVelocity {
		return true
	}
	return false
}

func (r *Rocket) HasCrashed() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if r.altitude <= 0 && r.velocity <= 0 && r.velocity < -MaxTouchDownVelocity {
		return true
	}
	return false
}

// IsAlive determines if the physical state can still change (true) or the simulation reached a dead end (false)
func (r *Rocket) IsAlive() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if r.velocity != 0 || r.engine.IsRunning() {
		return true
	}
	return false
}

func (r *Rocket) GetMotion() Motion {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return Motion{
		TickSerial:   r.tickSerial,
		Timestamp:    r.elapsed(),
		Velocity:     r.velocity,
		Acceleration: r.acceleration,
		Fuel:         r.fuel,
		Altitude:     r.altitude,
		TotalWeight:  r.deadWeight + r.fuel,
	}
}

func (r *Rocket) Run(motion chan Motion, valve chan float64) chan struct{} {
	done := make(chan struct{})
	go func() {
		r.startTime = time.Now()
		r.sendMotion(motion, time.Microsecond*10)
		initialValveSetting := <-valve
		r.engine.Ignite(initialValveSetting)
		for !r.HasLanded() && !r.HasCrashed() && r.IsAlive() {
			// compute adaptive wait time, accounting for time slew caused by previous ticks
			nextTickTime := r.startTime.Add(time.Duration(float64(r.tickSerial+1)*TimeTickMs) * time.Millisecond)

			if time.Now().After(nextTickTime) {
				panic(fmt.Sprintf("invalid wait time in Rocket.Run(): negative waitTime %v", time.Now().Sub(nextTickTime)))
			}
			waitTime := nextTickTime.Sub(time.Now())
			if float64(waitTime.Milliseconds()) > TimeTickMs*2 {
				panic(fmt.Sprintf("invalid wait time in Rocket.Run(): exceeds TimeTickMs*2 (%v): %v", TimeTickMs*2, waitTime))
			}

			time.Sleep(waitTime)

			/* Maximum wait time for signals is a tenth of a time tick.
			Note the engine control algorithm has far more time than that to make up its mind about the valve setting,
			as the sleep in the previous line is almost as long as a full timeTickMs.
			*/
			maxWait := time.Duration(float64(time.Millisecond) * TimeTickMs / 10)
			r.peekValveSetting(valve, maxWait)
			r.Tick() // update physical system state, i.e. move the rocket
			r.sendMotion(motion, maxWait)
		}

		done <- struct{}{}
		r.UpdateThrottle(0)
		log.Printf("INFO: Simulation ended after %v ticks, %.6f seconds\n", r.tickSerial, r.elapsed())
	}()

	return done
}

// peekValveSetting performs an unblocking read, making sure to return before waitTime has elapsed if there is no value available
func (r *Rocket) peekValveSetting(valve chan float64, waitTimeSec time.Duration) error {
	select {
	case v, ok := <-valve:
		if !ok {
			err := fmt.Errorf("(tick #%v) peekValveSetting() tried reading from a closed channel. Moving on.", r.tickSerial)
			log.Printf("WARNING: %v\n", err)
			return err
		}
		err := r.UpdateThrottle(v)
		return err
	case <-time.After(waitTimeSec):
		err := fmt.Errorf("(tick #%v) engine control failed to provide throttle setting in time. Moving on.", r.tickSerial)
		log.Printf("WARNING: %v\n", err)
		return err // don't wait for when value becomes available
	}
}

// sendMotion performs an unblocking write, making sure to return before waitTime has elapsed if noone is reading on the other end
func (r *Rocket) sendMotion(motion chan Motion, waitTime time.Duration) error {
	m := r.GetMotion()
	fmt.Printf("INFO: %v\n", m)
	select {
	case motion <- m:
		return nil
	case <-time.After(waitTime):
		err := fmt.Errorf("(tick #%v) engine control failed to consume motion tick in time. Moving on.\n", m.TickSerial)
		log.Printf("WARNING: %v\n", err)
		return err
	}
}

// elapsed returns the time since the start of the simulation in seconds with microseconds resolution
func (r *Rocket) elapsed() float64 {
	return float64(time.Now().Sub(r.startTime).Microseconds()) / 1e6
}
