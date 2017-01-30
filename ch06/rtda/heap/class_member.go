package heap

import "jvmgo/ch06/classfile"

type ClassMember struct {
  accessFlags uint16
  name        string
  descriptor  string
  class       *Class
}
