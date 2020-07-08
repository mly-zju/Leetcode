type ListNode struct {
	Val int
	Next *ListNode
}


func lastRemaining(n int, m int) int {
	// 1. 用数组来搞, 会超出时间限制
	// arr := make([]int, n)
	// for i := 0; i < n; i++ {
	// 	arr[i] = i
	// }
	
	// count := n
	// cursor := 0
	// for count > 1 {
	// 	cursor = (cursor + m - 1) % count
	// 	arr = append(arr[0:cursor], arr[cursor + 1:]...)
	// 	count--
	// }
	// return arr[0]

	// 2. 用循环链表来搞, 不过也会超时。。。
	// var head, cur *ListNode
	// for i := 0; i < n; i++ {
	// 	if head == nil {
	// 		head = &ListNode{
	// 			Val: i,
	// 		}
	// 		cur = head
	// 	} else {
	// 		cur.Next = &ListNode{
	// 			Val: i,
	// 		}
	// 		cur = cur.Next
	// 	}
	// }
	// cur.Next = head

	// count := n
	// cur = head
	// for count > 1 {
	// 	// 循环前进m-2次(m-1是走到要删除的点，m-2是要删除的前一个点)
	// 	for i := m - 2; i > 0; i-- {
	// 		cur = cur.Next
	// 	}
	// 	// 删除节点
	// 	cur.Next = cur.Next.Next
	// 	// 当前点前进
	// 	cur = cur.Next
	// 	count--
	// }
	// return cur.Val

	// 3. 数序公式推导法
	num := 0
	for i := 2; i <= n; i++ {
		num = (num + m) % i
	}
	return num
}