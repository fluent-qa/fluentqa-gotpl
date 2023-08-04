package pool

const (
	// Created is a pool state after pool has been created and before it starts.
	Created Status = iota
	// Running is a pool state when the pool started by Start() function.
	Running
	// Closed is a pool state when the pool stopped by Close() function.
	Closed
)

var (
	status2string = map[Status]string{
		Created: "Created",
		Running: "Running",
		Closed:  "Closed",
	}
)

type (
	// Status represents pool current state.
	Status int
)

// String returns string value of pool state.
func (p Status) String() string {
	return status2string[p]
}
