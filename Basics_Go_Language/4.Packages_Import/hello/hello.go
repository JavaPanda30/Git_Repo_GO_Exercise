package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {

	greetings.Abc()
	_, err := greetings.Hello("") //will show error 
	fmt.Print(err)

}

// sample program to do rec in GoLang


// var fib func(n int) int
// fib = func(n int) int {
// 	if n < 2 {
// 		return n
// 	}
// 	return fib(n-1) + fib(n-2)
// }

// fmt.Println(fib(7))
// var rec func(arr []int,curr []int,idx int) string
// rec = func(arr []int,curr []int,idx int) string{
// 	if(idx==len(arr)){
// 		return ""
// 	}
// 	fmt.Println(curr)
// 	rec(arr,append(curr, arr[idx]),idx+1)
// 	rec(arr,curr,idx+1)
// 	return ""
// }

// var arr = []int{1, 2, 3, 4, 5}
// var curr =[]int{}
// rec(arr,curr,0)
