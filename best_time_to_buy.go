package main

import "fmt"

func maxProfix(prices []int) int{
	var res int = 0
	var buying_price int 
	var lengt int = len(prices)
	for i:=0; i<lengt ;i++{
		if res ==0 || prices[buying_price]>prices[i]{ 
			for j:=i+1; j<lengt ;j++{
					if (res < prices[j]-prices[i]){
						res = prices[j]-prices[i]			
						buying_price = i 
					}	
			}
		}
	}
	return  res
}
func maxprofit2(prices []int) int {
	var res int 
	for i:=0; i<len(prices);i++{
		for j:=i; j<len(prices);j++{
			if i<j && res < prices[j]-prices[i]{
				res = prices[j]-prices[i]			
			}	
		}
	}
	return  res

}
//not my solution, sadly
func maxProfit(prices []int) int {
    var profit = 0
    var minPrice = prices[0]
    
    for i := 1; i < len(prices); i++ {
        // If we find any price which is lower than the current minPrice
        // update the minPrice
        if prices[i] < minPrice {
            minPrice = prices[i]
        } else if (prices[i] - minPrice) > profit {
            // If diff of current stock with minPrice is greater
            // update the profit
            profit = prices[i] - minPrice
        }
    }
    
    return profit
}
func main(){

	array := []int{5,224,465,296,1,415,968,40,494}
	array = []int{7,1,5,3,6,4}
	fmt.Println(maxProfix((array)))
}


