package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sumDigits(v1 int, v2 int) (int, int) {
	val := v1 + v2
	return val % 10, val / 10
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resultNode := &ListNode{Val: 0}
	loopJump := resultNode
	for !(l1 == nil && l2 == nil) {
		digit1 := 0
		if l1 != nil {
			digit1 = l1.Val
		}
		digit2 := 0
		if l2 != nil {
			digit2 = l2.Val
		}
		digit, sum1 := sumDigits(digit1+loopJump.Val, digit2)
		loopJump.Val = digit //+ loopJump.Val
		fmt.Println(loopJump)

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		if l1 != nil || l2 != nil || sum1 > 0 {
			loopJump.Next = &ListNode{sum1, nil}
			loopJump = loopJump.Next
		}
	}

	return resultNode
}

func main() {
	caseL1 := &ListNode{
		9, &ListNode{
			9, &ListNode{
				9, nil,
			},
		},
	}
	caseL2 := &ListNode{
		9, &ListNode{
			9, &ListNode{
				9, &ListNode{
					Val: 9, Next: &ListNode{
						Val: 9, Next: nil,
					},
				},
			},
		},
	}

	lr := addTwoNumbers(caseL1, caseL2)
	fmt.Println("LR:", lr)
	for lr != nil {
		fmt.Println("LR:", lr.Val)
		lr = lr.Next
	}
}
