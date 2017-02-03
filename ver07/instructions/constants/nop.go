package constants

import "jvmgo/ver07/instructions/base"
import "jvmgo/ver07/rtda"

type NOP struct{base.NoOperandsInstruction}

func (self *NOP) Execute(frame *rtda.Frame) {

}
