package base

import "jvmgo/ver05/rtda"

func Branch(frame *rtda.Frame, offset int) {
  pc := frame.Thread().PC()
  nextPC := pc + offset
  frame.SetNextPC(nextPC)
}
