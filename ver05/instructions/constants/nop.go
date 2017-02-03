package constants

import "jvmgo/ver05/instructions/base"
import "jvmgo/ver05/rtda"

type NOP struct{base.NoOperandsInstruction}

func (self *NOP) Execute(frame *rtda.Frame) {

}
