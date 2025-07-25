package goleetcode

type Stack struct {
	items []int
}

func NewStack(items []int) *Stack {
	return &Stack{items: items}
}

func (s *Stack) Len() int {
	return len(s.items)
}

func (s *Stack) Push(v int) {
	s.items = append(s.items, v)
}

func (s *Stack) Pop() {
	l := s.Len()
	s.items = s.items[:l-1]
}

func (s *Stack) Contains(v int) bool {
	for _, item := range s.items {
		if item == v {
			return true
		}
	}
	return false
}

func permute(nums []int) [][]int {
	var result [][]int

	if len(nums) == 1 {
		result = append(result, nums)
		return result
	}

	var backtrack func(stack *Stack)
	backtrack = func(stack *Stack) {
		if stack.Len() == len(nums) {
			permutation := make([]int, len(nums))
			copy(permutation, stack.items)
			result = append(result, permutation)
			return
		}

		for _, num := range nums {
			if !stack.Contains(num) {
				stack.Push(num)
				backtrack(stack)
				stack.Pop()
			}
		}
	}

	stack := NewStack([]int{})
	backtrack(stack)

	return result
}
