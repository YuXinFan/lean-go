type eface struct {
	Typ, Val unsafe.Pointer
}

// x is interface type
// return the underlying type in the interface 
func PointerOf(x any) unsafe.Pointer {
	return (*eface)(unsafe.Pointer(&x)).Val
}
