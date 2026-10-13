package main

func getSum(a int, b int) int {
	for b != 0 {
		carry := a & b
		a = a ^ b
		b = carry << 1
	}
	return a
}

func main() {
	a := 5
	b := 3
	result := getSum(a, b)
	println(result) // Output: 8
}
