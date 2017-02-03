package rtda

import "jvmgo/ver06/rtda/heap"

type Slot struct {
  num int32
  ref *heap.Object
}
