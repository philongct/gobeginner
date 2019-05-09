package receiverandpointer

import (
	"testing"
	"fmt"
)

func TestReceiverAndPointer1(t *testing.T) {
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
	o := Object{"attr"}
	// output: attr    {newAttr}
	fmt.Print(o.ReplaceAttr("newAttr"), "\t")
	fmt.Println(o)
}

func TestReceiverAndPointer4(t *testing.T) {
	o := &Object{"attr"}
	o.ReplaceAttr("newAttr")
	// output: &{newAttr}       {newAttr}
	fmt.Println(o, "\t", *o)
}
