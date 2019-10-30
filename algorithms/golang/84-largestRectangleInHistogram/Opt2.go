package main

type stack []int

func (s stack) top() int {
	return s[len(s)-1]
}

func (s stack) empty() bool {
	return len(s) == 0
}

func (s *stack) pop() {
	*s = (*s)[:len(*s)-1]
}

func (s *stack) push(val int) {
	*s = append(*s, val)
}

func largestRectangleAreaOpt2(heights []int) int {
	var maxArea int
	s := make(stack, 0, len(heights))
	heights = append(heights, 0)
	s = append(s, 0)
	for i := 1; i < len(heights); i++ {
		for !s.empty() && heights[i] < heights[s.top()] {
			h := heights[s.top()]
			s.pop()
			var prev int
			if s.empty() {
				prev = -1
			} else {
				prev = s.top()
			}
			area := (i - prev - 1) * h
			if maxArea < area {
				maxArea = area
			}
		}

		s.push(i)
	}

	return maxArea
}
