package creational

type Object struct {
}

func (*Object) Do() {

}

func NewPool(total int) chan *Object {
	p := make(chan *Object, total)
	for i := 0; i < total; i++ {
		p <- new(Object)
	}
	return p
}
