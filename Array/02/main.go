package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println("maxProfit is: ", maxProfit(prices))
}

func maxProfit(prices []int) int {
	//set maxprofit=0
	//set minprice=price[index]
	maxProfit := 0
	minPrice := prices[0]

	//traverse the array check the price and profit

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}
	return maxProfit
}

// func maxProfit(prices []int) int {
//     var profit = 0
//     var minPrice = prices[0]

//     for i := 1; i < len(prices); i++ {
//         // If we find any price which is lower than the current minPrice
//         // update the minPrice
//         if prices[i] < minPrice {
//             minPrice = prices[i]
//         } else if (prices[i] - minPrice) > profit {
//             // If diff of current stock with minPrice is greater
//             // update the profit
//             profit = prices[i] - minPrice
//         }
//     }

//     return profit
// }
