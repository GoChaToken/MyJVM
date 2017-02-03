package math

import "jvmgo/ver05/instructions/base"
import "jvmgo/ver05/rtda"

type IOR struct { base.NoOperandsInstruction }
type LOR struct { base.NoOperandsInstruction }

func (self *IOR) Execute(frame *rtda.Frame) {
  stack := frame.OperandStack()
  v2 := stack.PopInt()
  v1 := stack.PopInt()
  result := v1 | v2
  stack.PushInt(result)
}

func (self *LOR) Execute(frame *rtda.Frame) {
  stack := frame.OperandStack()
  v2 := stack.PopLong()
  v1 := stack.PopLong()
  result := v1 | v2
  stack.PushLong(result)
}
