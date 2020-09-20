package behavioral

type Template interface {
	fun1() int
	fun2() int

	//算法骨架
	Result() int
}

// 抽象结构体
type Executor struct {
	Temp Template
}

// 抽象结构体部分实现接口
func (r *Executor) Result() int {
	return r.Temp.fun1() + r.Temp.fun2()
}

//A具体实现。继承自抽象结构体的方法+自身实现的方法=实现Template接口
type A struct {
	//会继承抽象结构体中的方法
	Executor
}

func (a *A) fun1() int {
	return 1
}

func (a *A) fun2() int {
	return 2
}

//B具体实现。继承自抽象结构体的方法+自身实现的方法=实现Template接口
type B struct {
	//会继承抽象结构体中的方法
	Executor
}

func (b *B) fun1() int {
	return 11
}

func (b *B) fun2() int {
	return 12
}
