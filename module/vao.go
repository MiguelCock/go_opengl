package src

type vao struct {
	val uint32
}

func (v *vao) sum(a uint32) *vao {
	v.val += a
	return v
}
