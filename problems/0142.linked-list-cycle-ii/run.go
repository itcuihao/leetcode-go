package ii

import "github.com/itcuihao/leetcode-go/structure"

func Run() {

}

func detectCycle(head *structure.ListNode) *structure.ListNode {
	f, s := head, head
	isC := false
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
		if s == f {
			isC = true
			break
		}
	}
	if !isC {
		return nil
	}
	s = head
	for s != f {
		f = f.Next
		s = s.Next
	}
	return f
}
