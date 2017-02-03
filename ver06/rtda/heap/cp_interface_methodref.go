package heap

import "jvmgo/ver06/classfile"

type InterfaceMethodRef struct {
  MemberRef
  method  *Method
}

func newInterfaceMethodRef(cp *ConstantPool,refInfo *classfile.ConstantInterfaceMethodrefInfo) *InterfaceMethodRef {
  ref := &InterfaceMethodRef{}
  ref.cp = cp
  ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
  return ref
}
