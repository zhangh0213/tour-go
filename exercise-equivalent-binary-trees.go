package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int){
	Recur(t, ch)
	close(ch)
}

func Recur(t *tree.Tree, ch chan int){
	if t == nil{
		return
	}
	Recur(t.Left, ch)
	ch <- t.Value
	Recur(t.Right, ch)
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool{
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	Walk(t1, ch1)
	Walk(t2, ch2)
	for value1 := range ch1{
		if value1 != <-ch2{
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int, 10)
	go Walk(tree.New(1), ch)
	for value := range ch{
		fmt.Printf("%d ", value)
	}
	fmt.Println()
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
