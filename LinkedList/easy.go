package LinkedList

type ListNode struct {
	Val  int
	Next *ListNode
}

func EasyRun() {

}

//21. Merge Two Sorted Lists
//Example:
//Input: 1->2->4, 1->3->4
//Output: 1->1->2->3->4->4
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	//var head, temp *ListNode
	//if list1 == nil && list2 == nil {
	//	return nil
	//}
	//if list1 == nil {
	//	return list2
	//}
	//if list2 == nil {
	//	return list1
	//}
	//
	//找頭
	//if list1.Val <= list2.Val {
	//	head, temp = list1, list1
	//	list1 = list1.Next
	//} else {
	//	head, temp = list2, list2
	//	list2 = list2.Next
	//}
	//
	//拚後面的
	//for list1 != nil && list2 != nil {
	//	if list1.Val <= list2.Val {
	//		temp.Next = list1
	//		temp = temp.Next
	//		list1 = list1.Next
	//	} else {
	//		temp.Next = list2
	//		temp = temp.Next
	//		list2 = list2.Next
	//	}
	//}
	//看最後誰長
	//if list1 != nil {
	//	temp.Next = list1
	//}
	//if list2 != nil {
	//	temp.Next = list2
	//}
	//
	//return head
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}
