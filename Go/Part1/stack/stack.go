package stack

type CustomStack struct {
	list      []int
	index     int
}

func ConstructCustomStack() CustomStack {
	return CustomStack{index: -1}
}

func (stack *CustomStack) Pop() int {
	if stack.index == -1 {
		panic("underflow")
	}

	val := stack.list[stack.index]
	stack.list = stack.list[:stack.index]
	stack.index--
	return val
}

func (stack *CustomStack) Push(value int) {
	stack.list = append(stack.list, value)
	stack.index++
}
