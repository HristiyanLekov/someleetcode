package main

import (
	"fmt"
)

func int_value(letter string) int{
	if (letter == "I") {return 1}
	if (letter == "V") {return 5}
	if (letter == "X") {return 10}
	if (letter == "L") {return 50}
	if (letter == "C") {return 100}
	if (letter == "D") {return 500}
	if (letter == "M") {return 1000}
	return 0
}

func romanToInt(s string) int{
	var res int
	var tmp_res int
	var nex_val int
	var current_val int
	for i:=0; i<len(s); i++	{
		current_val = int_value(string(s[i]))
		if (i<len(s)-1){nex_val = int_value(string(s[i+1]))}
		if ((i<len(s)-1) && current_val<nex_val){ 
			tmp_res+=current_val
			res-=tmp_res
			tmp_res=0
		} else if (i<len(s)-1 && current_val>nex_val){
			tmp_res+=current_val
			res+=tmp_res
			tmp_res=0
		} else if (i<len(s)-1 && current_val==nex_val){
			tmp_res+=current_val
		}else{
			res+=tmp_res
			res+=current_val
		}
	}
	return res
}


func main(){
	var numb string = "DCXXI" 
	fmt.Println(romanToInt(numb))
}
