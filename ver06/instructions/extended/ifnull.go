package extended

import "jvmgo/ver06/instructions/base"
import "jvmgo/ver06/rtda"

type IFNULL struct { base.BranchInstruction }
type IFNONNULL struct { base.BranchInstruction }

func (self *IFNULL) Execute(frame *rtda.Frame) {
  ref := frame.OperandStack().PopRef()
  if ref == nil {
    base.Branch(frame, self.Offset)
  }
}

func (self *IFNONNULL) Execute(frame *rtda.Frame) {
  ref := frame.OperandStack().PopRef()
  if ref != nil {
    base.Branch(frame, self.Offset)
  }
}
