package receiverandpointer

type Object struct {
	Attr string
}

func (o Object) SetAttr(value string) {
	o.Attr = value
}

func (o *Object) ReplaceAttr(newValue string) string {
	ret := o.Attr
	o.Attr = newValue

	return ret
}
