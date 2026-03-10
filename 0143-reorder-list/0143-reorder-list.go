/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
    if head == nil || head.Next == nil {
        return
    }

    // find middle
    slow := head
    fast := head.Next
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    cur := slow.Next
    slow.Next = nil
    var prev *ListNode

    // revserse linked list
    for cur != nil {
        cur.Next, prev, cur = prev, cur, cur.Next
    }

    first := head
    second := prev
    for second != nil {
        tmp1, tmp2 := first.Next, second.Next
        first.Next = second
        second.Next = tmp1
        first, second = tmp1, tmp2
    }
}