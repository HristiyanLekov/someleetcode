package main
import "fmt"
func twoSum(nums []int, target int) []int {
	for i:=0; (i < len(nums)) ;i++{
		for j:=i+1; j < len(nums) ;j++{
			if nums[i]+nums[j]==target{
				return []int{i,j}
			}
		}
	}
	return []int{}
}


func main(){
	v := []int{3,2,4}
	var result []int
	result = twoSum(v,6)
	fmt.Println(len(v))
	fmt.Println(result)
}
