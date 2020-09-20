package creational

import (
	"fmt"
	"testing"
)

func Test_SingletonPattern(t *testing.T) {
	s1 := NewSingleton()
	s1["this"] = "s1"
	s2 := NewSingleton()
	fmt.Println("This is ", s2["this"])
}
