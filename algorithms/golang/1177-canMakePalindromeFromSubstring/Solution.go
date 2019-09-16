package main

func canMakePaliQueries(s string, queries [][]int) []bool {
	// i 表示前i字符中每个字符出现的次数
	sum := make([][26]int, len(s)+1, len(s)+1)
	for i := 0; i < len(s); i++ {
		sum[i+1] = sum[i]
		sum[i+1][s[i]-'a']++
	}

	can := make([]bool, len(queries), len(queries))
	for i, query := range queries {
		diff := 0
		for j := 0; j < 26; j++ {
			diff += (sum[query[1]+1][j] - sum[query[0]][j]) % 2
		}

		can[i] = diff/2 <= query[2]
	}
	return can
}
