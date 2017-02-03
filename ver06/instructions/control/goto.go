package control

import "jvmgo/ver06/instructions/base"
import "jvmgo/ver06/rtda"

type GOTO struct { base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
  base.Branch(frame, self.Offset)
}
