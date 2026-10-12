package main

func reverseBits(n uint32) uint32 {
	var ans uint32
	for i := 0; i < 32; i++ {
		ans = (ans << 1) | (n & 1)
		n = n >> 1
	}
	return ans
}

func main() {
	n := uint32(43261596)
	result := reverseBits(n)
	println(result) // Output: 964176192
}
