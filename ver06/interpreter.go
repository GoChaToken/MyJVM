package main

import "fmt"
import "jvmgo/ver06/instructions"
import "jvmgo/ver06/instructions/base"
import "jvmgo/ver06/rtda"
import "jvmgo/ver06/rtda/heap"

func interpret(method *heap.Method){
  thread := rtda.NewThread()
  frame := thread.NewFrame(method)
  thread.PushFrame(frame)

  defer catchErr(frame)
  loop(thread, method.Code())
}

func catchErr(frame *rtda.Frame) {
  if r := recover(); r != nil {
    fmt.Println("LocalVars: %v\n",frame.LocalVars())
    fmt.Println("OperandStack: %v\n", frame.OperandStack())
    panic(r)
  }
}

func loop(thread *rtda.Thread, bytecode []byte){
  frame := thread.PopFrame()
  reader := &base.BytecodeReader{}

  for {
    pc := frame.NextPC()
    thread.SetPC(pc)

    reader.Reset(bytecode, pc)
    opcode := reader.ReadUint8()
    inst := instructions.NewInstruction(opcode)
    inst.FetchOperands(reader)
    frame.SetNextPC(reader.PC())

    fmt.Printf("pc:%2d inst: %T %v\n",pc, inst, inst)
    inst.Execute(frame)
  }
}
