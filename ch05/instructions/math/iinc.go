package math

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

type IINC struct {
  Index uint
  Const int32
}

func (self *IINC) FetchOperands(reader *base.BytecodeReader) {
  self.Index = uint(reader.ReadUint8())
  self.Const = int32(reader.ReadInt8())
}

func (self *IINC) Execute(frame *rtda.Frame) {
  LocalVars := frame.LocalVars()
  val := LocalVars.GetInt(self.Index)
  val += self.Const
  LocalVars.SetInt(self.Index, val)
}
