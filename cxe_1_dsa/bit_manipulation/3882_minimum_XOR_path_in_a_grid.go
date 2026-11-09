package main

import "math/bits"

type Bitset [16]uint64

var masks = [6]uint64{
	0x5555555555555555, 0x3333333333333333, 0x0f0f0f0f0f0f0f0f,
	0x00ff00ff00ff00ff, 0x0000ffff0000ffff, 0x00000000ffffffff,
}
var dp [1000]Bitset

func xorPerm(w, lo uint64) uint64 {
	for bit := range 6 {
		if (lo>>bit)&1 == 0 {
			continue
		}
		stride := 1 << bit
		mask := masks[bit]
		w = (w&mask)<<stride | (w&^mask)>>stride
	}
	return w
}

func minCost(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	clear(dp[0][:])
	v0 := grid[0][0]
	dp[0][v0>>6] = 1 << (v0 & 63)
	for j := 1; j < n; j++ {
		v := uint64(grid[0][j])
		hi, lo := v>>6, v&63
		for w := range uint64(16) {
			dp[j][w] = xorPerm(dp[j-1][w^hi], lo)
		}
	}
	for i := 1; i < m; i++ {
		above := dp[0]
		v := uint64(grid[i][0])
		hi, lo := v>>6, v&63
		for w := range uint64(16) {
			dp[0][w] = xorPerm(above[w^hi], lo)
		}
		for j := 1; j < n; j++ {
			above = dp[j]
			v := uint64(grid[i][j])
			hi, lo := v>>6, v&63
			for w := range uint64(16) {
				dp[j][w] = xorPerm(above[w^hi]|dp[j-1][w^hi], lo)
			}
		}
	}
	bs := dp[n-1]
	for k := range 16 {
		if bs[k] != 0 {
			return k*64 + bits.TrailingZeros64(bs[k])
		}
	}
	return -1
}