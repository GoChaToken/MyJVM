package constants

import "jvmgo/ver06/instructions/base"
import "jvmgo/ver06/rtda"

type NOP struct{base.NoOperandsInstruction}

func (self *NOP) Execute(frame *rtda.Frame) {

}
