package math

import "jvmgo/ver07/instructions/base"
import "jvmgo/ver07/rtda"

type DMUL struct { base.NoOperandsInstruction }
type FMUL struct { base.NoOperandsInstruction }
type IMUL struct { base.NoOperandsInstruction }
type LMUL struct { base.NoOperandsInstruction }

func (self *DMUL) Execute(frame *rtda.Frame) {
  stack := frame.OperandStack()
  v2 := stack.PopDouble()
  v1 := stack.PopDouble()
  result := v1 * v2
  stack.PushDouble(result)
}

func (self *FMUL) Execute(frame *rtda.Frame) {
  stack := frame.OperandStack()
  v2 := stack.PopFloat()
  v1 := stack.PopFloat()
  result := v1 * v2
  stack.PushFloat(result)
}

func (self *IMUL) Execute(frame *rtda.Frame) {
  stack := frame.OperandStack()
  v2 := stack.PopInt()
  v1 := stack.PopInt()
  result := v1 * v2
  stack.PushInt(result)
}

func (self *LMUL) Execute(frame *rtda.Frame) {
  stack := frame.OperandStack()
  v2 := stack.PopLong()
  v1 := stack.PopLong()
  result := v1 * v2
  stack.PushLong(result)
}
