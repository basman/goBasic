package physics

const tickDelayLimitMs = TimeTickMs / 10 // milliseconds; alert if waiting for write lock longer than this
const controlWaitLimitMs = float64(0.02) // milliseconds; wait at most 20 microseconds for engine control signal
