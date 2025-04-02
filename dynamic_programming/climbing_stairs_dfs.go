package dynamic_programming

import "fmt"

/*
使用dp(动态规划) 问题的思路来解决

假设爬 n 层楼梯有 dp[n]个解答

爬到 n 层可能是从 n-1 层，或者是 n -2 层上来的

所以 dp[n]=dp[n-1]+dp[n-1]
*/

func climbingStairsDfs(n int) int {
	// 爬一层又一个解，爬2层有两个解，这里刚好爬的层数和解答数字相等
	// 所以直接返回，正所谓最小子问题的解是已知的
	if n == 1 || n == 2 {
		return n
	}
	// 当 n > 2 的时候

	return climbingStairsDfs(n-1) + climbingStairsDfs(n-2)

}

/*
但是这种算法的时间复杂度太高，通过观察递归树不难发现，这种算法的时间
复杂度为2^n,所以结下来改进 ⬇️

注意这是从顶至下的解决问题
*/

// elm[n]=dp[n]
func climbingStairsDfs2(n int, elm []int) int {
	// 爬一层又一个解，爬2层有两个解，这里刚好爬的层数和解答数字相等
	// 所以直接返回，正所谓最小子问题的解是已知的
	if n == 1 || n == 2 {
		elm[n-1] = n
		return n
	}
	// 当 n > 2 的时候,先从存储中找结果
	if elm[n-1] > 0 {
		return elm[n-1]
	}
	// 找不到就还按照原来的方法

	count := climbingStairsDfs2(n-1, elm) + climbingStairsDfs2(n-2, elm)
	elm[n-1] = count
	return count

}

/*
使用动态规划从 底至上的解决问题，
使用迭代的方式
*/

func climbingStairsDp(n int) int {

	if n == 1 || n == 2 {
		return n
	}
	// 动态规划数组,存储当楼梯数为 n 的时候，解的个数
	dp := make([]int, n+1) //[0,n],第一个数组不用
	dp[1], dp[2] = 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	fmt.Println(dp)
	return dp[n]
}

/*
由此我想到了斐波那契数列是经典的动态规划问题
a1=1
a2=2
an=a(n-1)+a(n-2)
使用动态规划求解第n个斐波那契数列数
*/

func FebnachiDp(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	// 创造一个数组，存储第 n 个斐波那契数列数
	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	fmt.Println(dp)
	return dp[n]
}

/*
使用递归的方式求解第n个斐波那契数列数
*/

func FebnachiDfs(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	return FebnachiDfs(n-1) + FebnachiDfs(n-2)
}
