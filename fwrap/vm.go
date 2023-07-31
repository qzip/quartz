/*
Copyright (c) 2019, QzIP Blockchain Technology LLP

All Rights Reserved.

Author: Ashish Banerjee, <ashish@qzip.in>


*/

package fwrap

import "context"

// VirtualMachine vm
type VirtualMachine struct {
	GlobalVarSpace VarSpace
	LocalVarStack  Stack
	ProgStack      Stack
	Funcs          map[string]FuncExecutor
	BuiltIns       map[string]FuncExecutor
}

// VarSpace variable spaces
type VarSpace map[string]interface{}

// Stack stack of elements (variables or instructions)
type Stack interface {
	//Empty Tests if this stack is empty.
	Empty() bool
	//Peek Looks at the object at the top of this stack without removing it from the stack.
	Peek() interface{}
	// Pop Removes the object at the top of this stack and returns that object as the value of this function.
	Pop() interface{}
	// Push Pushes an item onto the top of this stack.
	Push(vs interface{})
}

// FuncExecutor Function executor
type FuncExecutor interface {
	Execute(ctx context.Context, vm VirtualMachine) error
}

// InstructionCode runtime unit of PM, instruction & args
type InstructionCode struct {
	Code OpCode
	Args []Argument
}

// OpCode a memeber of the Instruction set
type OpCode int8

// Argument is the runtime arg of an instruction
type Argument interface {
}

// Program is an array of instructions
type Program []InstructionCode

var instructionSet = make(map[OpCode]InstructionExec)

// RegisterInstructionExec Registers an instruction executor
func RegisterInstructionExec(op OpCode, exe InstructionExec) {
	instructionSet[op] = exe
}

// InstructionExec an single instruction
type InstructionExec func(ctx context.Context, vm VirtualMachine) error
