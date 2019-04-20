package string

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
	l1, l2, l3 := len(s1), len(s2), len(s3)
	if l1+l2 != l3 {
		return false
	}

	// 二维初始化
	dp := make([][]bool, l2+1)
	for k := range dp {
		dp[k] = make([]bool, l1+1)
	}
	// 第一种方式
	// dp[0][0] = true
	// ldp := len(dp)
	// for i := 1; i < ldp; i++ {
	// 	dp[i][0] = dp[i-1][0] && s2[i-1] == s3[i-1]
	// }
	// ldpi := len(dp[0])
	// for i := 1; i < ldpi; i++ {
	// 	dp[0][i] = dp[0][i-1] && s1[i-1] == s3[i-1]
	// }

	// for i := 1; i < len(dp); i++ {
	// 	for j := 1; j < ldpi; j++ {
	// 		dp[i][j] = (dp[i-1][j] && s2[i-1] == s3[i+j-1]) ||
	// 			(dp[i][j-1] && s1[j-1] == s3[i+j-1])
	// 	}
	// }

	// 第二种方式

	for i := 0; i < l2+1; i++ {
		for j := 0; j < l1+1; j++ {
			switch {
			case i == 0 && j == 0:
				dp[i][j] = true
			case i == 0:
				dp[0][j] = dp[0][j-1] && s1[j-1] == s3[i+j-1]
			case j == 0:
				fmt.Println("i:", i)
				dp[i][0] = dp[i-1][0] && s2[i-1] == s3[i+j-1]
			default:
				dp[i][j] = (dp[i-1][j] && s2[i-1] == s3[i+j-1]) || (dp[i][j-1] && s1[j-1] == s3[i+j-1])
			}
		}
	}
	fmt.Println(dp)
	return dp[l2][l1]
}
