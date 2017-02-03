package stack

import "jvmgo/ver05/instructions/base"
import "jvmgo/ver05/rtda"

type SWAP struct { base.NoOperandsInstruction }

func (self *SWAP) Execute(frame *rtda.Frame) {
  stack := frame.OperandStack()
  slot1 := stack.PopSlot()
  slot2 := stack.PopSlot()
  stack.PushSlot(slot1)
  stack.PushSlot(slot2)
}
