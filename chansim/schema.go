package chansim

import "time"

/*
Simulators for supporting qz/cmd/ChanExec

*/

type SimGenHelper interface {
	GenRec() interface{}
}

type ChanSimPayload struct {
	Timestamp time.Time
	Count     int
}
