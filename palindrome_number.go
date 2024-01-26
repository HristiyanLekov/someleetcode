package main

import (
	"fmt"
	"strconv"
)
//https://leetcode.com/problems/palindrome-number/submissions/1157811153/

func isPalindrome(x int) bool{
	array := strconv.Itoa(x)
	var lenght int = len(array)
	for i:=0; i<(lenght/2); i++{
		if (array[i]!=array[lenght-1-i]){
			return false
		}
	}
	return true
}

func main(){
	fmt.Println(isPalindrome(121))	
}
