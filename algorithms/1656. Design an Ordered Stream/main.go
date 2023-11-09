package main

import "fmt"

type OrderedStream struct {
	values []string
	ptr    int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{values: make([]string, n)}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.values[idKey-1] = value

	lastPtr := this.ptr
	for ; this.ptr < len(this.values) && this.values[this.ptr] != ""; this.ptr++ {
	}
	return this.values[lastPtr:this.ptr]

}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */

func main() {
	obj := Constructor(5)
	param_1 := obj.Insert(3, "asdasd")
	fmt.Println(param_1)
}
