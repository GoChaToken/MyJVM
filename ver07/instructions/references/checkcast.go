package references

import "jvmgo/ver07/instructions/base"
import "jvmgo/ver07/rtda"
import "jvmgo/ver07/rtda/heap"

type CHECK_CAST struct { base.Index16Instruction }

func (self *CHECK_CAST) Execute(frame *rtda.Frame) {
  stack := frame.OperandStack()
  ref := stack.PopRef()
  stack.PushRef(ref)
  if ref == nil {
    return
  }

  cp := frame.Method().Class().ConstantPool()
  classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
  class := classRef.ResolvedClass()
  if !ref.IsInstanceOf(class) {
    panic("java.lang.ClassCastException")
  }
}
