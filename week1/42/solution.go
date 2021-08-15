package _2

func trap(height []int) int {
	stack := newStack()
	var water int

	for index, value := range height {
		top := stack.top()

		for !stack.isEmpity() && height[top] < value {
			stack.pop()

			if stack.isEmpity() {
				break
			}

			cur := stack.top()
			if height[cur] < value {
				water += (index - cur - 1) * (height[cur] - height[top])
			} else {
				water += (index - cur - 1) * (value - height[top])
			}
			top = cur
		}

		stack.push(index)
	}

	return water
}

type stack struct {
	array []int
}

func newStack() *stack {
	return &stack{
		array: []int{},
	}
}

func (s *stack) top() int {
	if s.isEmpity() {
		return -1
	}

	return s.array[len(s.array)-1]
}

func (s *stack) pop() {
	s.array = s.array[:len(s.array)-1]
	return
}

func (s *stack) push(e int) {
	s.array = append(s.array, e)
	return
}

func (s *stack) isEmpity() bool {
	return len(s.array) == 0
}
