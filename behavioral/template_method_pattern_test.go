package behavioral

import (
	"fmt"
	"testing"
)

func TestFuncs_Result(t *testing.T) {
	b := new(B)
	executor := Executor{Temp: b}
	fmt.Print(executor.Result())
}
