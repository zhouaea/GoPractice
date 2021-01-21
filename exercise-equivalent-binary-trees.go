// Write a function to determine if two binary trees have the same values.
package main

import (
	"fmt"
	"golang.org/x/tour/tree"
	"sync"
	"container/heap"
	"strconv"
)

// Credit to Golang.org for priority queue implementation.
// An Item is something we manage in a priority queue.
type Item struct {
	value    int // The value of the item; priority of the item.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].value > pq[j].value
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface {}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface {} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	var err error
	item.value, err = strconv.Atoi(value)
	fmt.Println(err)
	heap.Fix(pq, item.index)
}

// Credit to Kuldeep for providing a way to close channel after recursively walking through the tree.
func Walker(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	ch <- t.Value
	
	// Iterate through left branches until there are no more.
	// Iterate through right branches afterwards.
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	
	go Walker(t1, ch1)
	go Walker(t2, ch2)
	
	// Add values from both trees into seperate max priority queues concurrently to sort their values.
	wg := sync.WaitGroup{}
	pq1 := make(PriorityQueue, 10)
	pq2 := make(PriorityQueue, 10)
	
	wg.Add(1)
	go func() {
		i1 := 0
		for v := range ch1 {
			pq1[i1] = &Item{
				value: v,
				index: i1,
			}
			i1++
		}
		heap.Init(&pq1)
		wg.Done()
	}()
	
	wg.Add(1)
	go func() {
		i2 := 0
		for v := range ch2 {
			pq1[i2] = &Item{
				value: v,
				index: i2,
			}
			i2++
		}
		heap.Init(&pq2)
		wg.Done()
	}()
	
	// Compare sorted elements within each priority queue once they are finished being loaded.
	wg.Wait()
	
	for i := 0; i < 10; i++ {
		v1 := pq1.Pop()
		v2 := pq2.Pop()
		if v1 != v2 {
			fmt.Printf("v1: %T v2: %\n", v1, v2)
			return false
		}
	}
	
	return true
}	

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
}
