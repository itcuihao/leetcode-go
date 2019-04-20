package string

func isInterleave(s1 string, s2 string, s3 string) bool {
	l1, l2, l3 := len(s1), len(s2), len(s3)
	if l1+l2 != l3 {
		return false
	}

	// 二维初始化
	dp := make([][]bool, l1+1)
	for k := range dp {
		dp[k] = make([]bool, l2+1)
	}

	ldp := len(dp)
	for i := 1; i < ldp; i++ {
		dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
	}
	ldpi := len(dp[0])
	for i := 1; i < ldpi; i++ {
		dp[0][i] = dp[0][i-1] && s2[i-1] == s3[i-1]
	}

	for i := 1; i < ldp; i++ {
		for j := 1; j < ldpi; j++ {
			dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) ||
				(dp[i][j-1] && s2[j-1] == s3[i+j-1])
		}
	}
	return dp[l1][l2]
}
