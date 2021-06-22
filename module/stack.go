package module

import "errors"

type Stack []string

// Len 统计栈中元素的数量
func (stack Stack) Len() int {
	return len(stack)
}

// Cap 统计栈的容量
func (stack Stack) Cap() int {
	return cap(stack)
}

// Push 将元素押入栈
func (stack *Stack) Push(value string) {
	*stack = append(*stack, value)
}

// Top 获取栈顶元素
func (stack Stack) Top() (string, error) {
	if len(stack) == 0 {
		return "", errors.New("out of index")
	}
	return stack[len(stack)-1], nil
}
//Pop 获取栈顶的元素 并将其出栈
func (stack *Stack) Pop() (string, error) {
	theStack := *stack
	if len(theStack) == 0 {
		return "", errors.New("out of index")
	}
	value := theStack[len(theStack)-1]
	*stack = theStack[:len(theStack)-1]
	return value, nil
}

// IsEmpty 判断栈是否为空
func (stack Stack) IsEmpty() bool {
	return len(stack) == 0
}