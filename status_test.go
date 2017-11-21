package pena

import "testing"

var (
	circuitStatus CircuitStatus
)

func TestCircuitStatusDial(t *testing.T) {
	db := "your_db"
	circuitStatus.Dial("localhost:6379", db)

	if circuitStatus.source != db {
		t.Error("failed set source of redis")
	}
}

func TestCircuitStatusFailed(t *testing.T) {
	circuitStatus.SetFail("your_dest", "route_dest")

	if circuitStatus.Fail != true {
		t.Error("failed set failed")
	}
}
