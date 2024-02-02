package main

import "fmt"

const (
	a1 = '('
	a2 = ')'
	b1 = '{'
	b2 = '}'
	c1 = '['
	c2 = ']'
)

type item struct {
	val  uint8
	next *item
}

type stack struct {
	top *item
}

func newStack() *stack {
	return &stack{
		top: &item{},
	}
}

func (s *stack) pop() (uint8, bool) {
	if s.top.next == nil {
		return 0, false
	}
	val := s.top.next.val
	s.top.next = s.top.next.next
	return val, true
}

func (s *stack) push(val uint8) {
	i := &item{val: val}
	i.next = s.top.next
	s.top.next = i
	return
}

func isValid(s string) bool {
	stack := newStack()

	for i := 0; i < len(s); i++ {
		ch := s[i]
		switch ch {
		case a1, b1, c1:
			stack.push(ch)
		case a2:
			lCh, ok := stack.pop()
			if !ok {
				return false
			}
			if lCh != a1 {
				return false
			}
		case b2:
			lCh, ok := stack.pop()
			if !ok {
				return false
			}
			if lCh != b1 {
				return false
			}
		case c2:
			lCh, ok := stack.pop()
			if !ok {
				return false
			}
			if lCh != c1 {
				return false
			}
		}
	}
	if _, ok := stack.pop(); !ok {
		return true
	}
	return false
}

func main() {
	fmt.Println(isValid("()"))
}
