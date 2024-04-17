package linknode

import "sync"

// ListNode 定义链表的节点，使用泛型
type ListNode[T any] struct {
	Value T
	Next  *ListNode[T]
}

type LinkedList[T any] struct {
	mu     sync.Mutex
	Head   *ListNode[T]
	Length int
}

// Append 方法，现在更新长度
func (ll *LinkedList[T]) Append(value T) {
	ll.mu.Lock()
	defer ll.mu.Unlock()
	newNode := &ListNode[T]{Value: value}
	if ll.Head == nil {
		ll.Head = newNode
	} else {
		current := ll.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	ll.Length++
}

// Pop 方法，现在更新长度
func (ll *LinkedList[T]) Pop() *T {
	ll.mu.Lock()
	defer ll.mu.Unlock()

	if ll.Head == nil {
		return nil
	}
	value := ll.Head.Value
	ll.Head = ll.Head.Next
	ll.Length--
	return &value
}

// Length
func (ll *LinkedList[T]) GetLength() int {
	return ll.Length
}
