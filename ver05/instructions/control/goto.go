package control

import "jvmgo/ver05/instructions/base"
import "jvmgo/ver05/rtda"

type GOTO struct { base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
  base.Branch(frame, self.Offset)
}
