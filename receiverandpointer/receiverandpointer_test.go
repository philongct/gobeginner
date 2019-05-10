package receiverandpointer

import (
	"fmt"
	"testing"
)

func TestReceiverAndPointer1(t *testing.T) {
	// Even o is a pointer, attribute remained unchanged if receiver is not pointer
	// NOTE: non-pointer receiver work for both pointer and non-pointer
	o := &Object{"attr1"}
	o.SetAttr("abc")
	// output: &{attr1}    {attr1}
	fmt.Println(o, "\t", *o)
}

func TestReceiverAndPointer2(t *testing.T) {
	o := Object{"attr1"}
	o.SetAttr("abc")
	// output: {attr1}
	fmt.Println(o)
}

func TestReceiverAndPointer3(t *testing.T) {
	// Because receiver is pointer, attribute change will retain for non-pointer
	// Pointer receiver accept both pointer and non-pointer
	o := Object{"attr"}
	// output: attr    {newAttr}
	fmt.Print(o.ReplaceAttr("newAttr"), "\t")
	fmt.Println(o)
}

func TestReceiverAndPointer4(t *testing.T) {
	// Attribute change will retained for pointer
	// Pointer receiver accept both pointer and non-pointer
	o := &Object{"attr"}
	o.ReplaceAttr("newAttr")
	// output: &{newAttr}       {newAttr}
	fmt.Println(o, "\t", *o)
}
