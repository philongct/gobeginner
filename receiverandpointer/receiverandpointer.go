package receiverandpointer

type Object struct {
	Attr string
}

// Pass by value doesn't retain changes
func (o Object) SetAttr(value string) {
	o.Attr = value
}

// Pass by reference retains changes
func (o *Object) ReplaceAttr(newValue string) string {
	ret := o.Attr
	o.Attr = newValue

	return ret
}
