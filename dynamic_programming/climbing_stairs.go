package dynamic_programming

import "fmt"

/*
爬楼梯问题
爬到3 层楼梯，每次只能爬 1个或者 2 个台阶
试问共有几中可行的方法
1，2，3
2，3
1，3

求解方法的数量
*/

/*
choices {1,2} 每次可以上几个台阶,这个选择是固定死的，不管上到哪个台阶都只有这两个选择
state 当前在第几个台阶
n 目标要上到第几个台阶
res 保存解的数量
*/

func backTrack(choices []int, state int, n int, res *int) {
	if state == n {
		*res += 1 //
		return
	}
	for _, choice := range choices {
		// 剪枝头
		if (state + choice) > n {
			continue
		}
		// 继续尝试选择，更新状态
		backTrack(choices, state+choice, n, res)
		// 从这个递归出来就是回退
	}
}

func climbingStairsBacktrack() {
	choices := []int{1, 2}
	n := 4
	state := 0
	var res int
	backTrack(choices, state, n, &res)
	fmt.Println(res)
}
