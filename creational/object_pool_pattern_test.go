package creational

import "testing"

func Test_ObjectPoolPattern(t *testing.T) {
	p := NewPool(2)
	select {
	case obj := <-p:
		obj.Do()
		p <- obj
	default:
		// No more objects left — retry later or fail
		return
	}
}
