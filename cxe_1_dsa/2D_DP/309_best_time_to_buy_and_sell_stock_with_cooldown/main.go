package main

func maxProfit(prices []int) int {
	canBuy := 0
	cantBuy := -1000
	cooldown := -1000
	for i := 0; i < len(prices); i++ {
		cantBuy, canBuy, cooldown = max(cantBuy, canBuy-prices[i]), max(cooldown, canBuy), max(cantBuy+prices[i], cooldown)

	}
	return max(canBuy, cooldown)
}

func main() {
	prices := []int{1, 2, 3, 0, 2}
	println(maxProfit(prices))
}