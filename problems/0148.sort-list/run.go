package list

import (
	"leetcode-go/structure"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortListMerge(head *structure.ListNode) *structure.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := findMiddleList(head)
	right := sortListMerge(mid.Next)
	mid.Next = nil
	left := sortListMerge(head)
	return merge(left, right)
}

func findMiddleList(head *structure.ListNode) *structure.ListNode {

	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func merge(head1, head2 *structure.ListNode) *structure.ListNode {
	tmp := new(structure.ListNode)
	tail := tmp
	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			tail.Next = head1
			head1 = head1.Next
		} else {
			tail.Next = head2
			head2 = head2.Next
		}
		tail = tail.Next
	}
	if head1 != nil {
		tail.Next = head1
	} else {
		tail.Next = head2
	}
	return tmp.Next
}

// public class Solution {
//     private ListNode findMiddle(ListNode head) {
//         ListNode slow = head, fast = head.next;
//         while (fast != null && fast.next != null) {
//             fast = fast.next.next;
//             slow = slow.next;
//         }
//         return slow;
//     }

//     private ListNode merge(ListNode head1, ListNode head2) {
//         ListNode dummy = new ListNode(0);
//         ListNode tail = dummy;
//         while (head1 != null && head2 != null) {
//             if (head1.val < head2.val) {
//                 tail.next = head1;
//                 head1 = head1.next;
//             } else {
//                 tail.next = head2;
//                 head2 = head2.next;
//             }
//             tail = tail.next;
//         }
//         if (head1 != null) {
//             tail.next = head1;
//         } else {
//             tail.next = head2;
//         }

//         return dummy.next;
//     }

//     public ListNode sortList(ListNode head) {
//         if (head == null || head.next == null) {
//             return head;
//         }

//         ListNode mid = findMiddle(head);

//         ListNode right = sortList(mid.next);
//         mid.next = null;
//         ListNode left = sortList(head);

//         return merge(left, right);
//     }
// }
