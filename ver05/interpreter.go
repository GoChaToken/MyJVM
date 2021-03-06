package main

import "fmt"
import "jvmgo/ver05/classfile"
import "jvmgo/ver05/instructions"
import "jvmgo/ver05/instructions/base"
import "jvmgo/ver05/rtda"

func interpret(methodInfo *classfile.MemberInfo){
  codeAttr := methodInfo.CodeAttribute()
  maxLocals := codeAttr.MaxLocals()
  maxStack := codeAttr.MaxStack()
  bytecode := codeAttr.Code()

  thread := rtda.NewThread()
  frame := thread.NewFrame(maxLocals,maxStack)
  thread.PushFrame(frame)

  defer catchErr(frame)
  loop(thread, bytecode)
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
