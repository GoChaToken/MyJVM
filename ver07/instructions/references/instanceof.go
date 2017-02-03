package references

import "jvmgo/ver07/instructions/base"
import "jvmgo/ver07/rtda"
import "jvmgo/ver07/rtda/heap"

type INSTANCE_OF struct { base.Index16Instruction }

func (self *INSTANCE_OF) Execute(frame *rtda.Frame) {
  stack := frame.OperandStack()
  ref := stack.PopRef()
  if ref == nil {
    stack.PushInt(0)
    return
  }

  cp := frame.Method().Class().ConstantPool()
  classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
  class := classRef.ResolvedClass()
  if ref.IsInstanceOf(class) {
    stack.PushInt(1)
  }else {
    stack.PushInt(0)
  }
}
