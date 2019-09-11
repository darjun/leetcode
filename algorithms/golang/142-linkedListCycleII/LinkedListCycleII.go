package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	mapNode := make(map[*ListNode]struct{})

	for head != nil {
		if _, exist := mapNode[head]; exist {
			return head
		}

		mapNode[head] = struct{}{}
		head = head.Next
	}

	return nil
}

type solutionFunc func(*ListNode) *ListNode

type testCase struct {
	head     *ListNode
	expected *ListNode
}

func main() {
	tests := []testCase{}

	n11 := &ListNode{Val: 3}
	n12 := &ListNode{Val: 2}
	n13 := &ListNode{Val: 0}
	n14 := &ListNode{Val: -4}
	n11.Next = n12
	n12.Next = n13
	n13.Next = n14
	n14.Next = n12
	tests = append(tests, testCase{n11, n12})

	n21 := &ListNode{Val: 1}
	n22 := &ListNode{Val: 2}
	n21.Next = n22
	n22.Next = n21
	tests = append(tests, testCase{n21, n21})

	n31 := &ListNode{Val: 1}
	tests = append(tests, testCase{n31, nil})

	solutions := []solutionFunc{detectCycle, detectCycleOpt1}

	for solutionIndex, f := range solutions {
		for testIndex, testCase := range tests {
			actual := f(testCase.head)
			if actual != testCase.expected {
				fmt.Printf("%dth solution, %dth test case failed, head:%v expected:%v actual:%v\n", solutionIndex+1, testIndex+1, testCase.head, testCase.expected, actual)
			}
		}
	}
}
