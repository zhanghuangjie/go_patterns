package behavioral

import (
	"testing"
)

func Test_StrategyPattern(t *testing.T) {
	add := Operation{Addition{}}
	add.Operate(3, 5) // 8

	mult := Operation{Multiplication{}}
	mult.Operate(3, 5) // 15
}
