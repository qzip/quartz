/*
Copyright (c) 2019-20, QzIP Blockchain Technology LLP

All Rights Reserved.

Author: Ashish Banerjee, <ashish@qzip.in>


*/

package fwrap

import "context"

const (
	// OCnop NOP
	OCnop OpCode = 0
	// OCcall calls a builtin or local function
	OCcall
	// OCcond if/then/else operation
	OCcond
	//OCassign assignment lvar = rvar
	OCassign
)

func init() {
	RegisterInstructionExec(OCnop, Nop)
}

// Nop OpNop
func Nop(ctx context.Context, vm VirtualMachine) error {
	return nil
}
