package stack

import "jvmgo/ver06/instructions/base"
import "jvmgo/ver06/rtda"

type POP struct { base.NoOperandsInstruction }
type POP2 struct { base.NoOperandsInstruction }

func (self *POP) Execute(frame *rtda.Frame) {
  stack := frame.OperandStack()
  stack.PopSlot()
}

func (self *POP2) Execute(frame *rtda.Frame) {
  stack := frame.OperandStack()
  stack.PopSlot()
  stack.PopSlot()
}
