package dynamic_programming

import (
	"fmt"
	"testing"
)

func TestClimbingStairsBacktrack(t *testing.T) {
	climbingStairsBacktrack()
}

func TestClimbingStairDfs(t *testing.T) {
	n := 3
	res := climbingStairsDp(n)
	fmt.Println(res)
}

func TestClimbingStairDfs2(t *testing.T) {
	n := 3
	elm := make([]int, n)
	res := climbingStairsDfs2(n, elm)
	fmt.Println(res)
	fmt.Println(elm)
}

func TestClimbingStairDp(t *testing.T) {

	n := 3
	res := climbingStairsDp(n)
	fmt.Println(res)
}

/*
测试斐波那契数列数
动态规划
*/
func TestFebnachi(t *testing.T) {
	n := 10
	febnachi := FebnachiDp(n)
	fmt.Println(febnachi)
}

/*
测试斐波那契数列数
递归
*/
func TestFebnachiDfs(t *testing.T) {
	n := 10
	dfs := FebnachiDfs(n)
	fmt.Println(dfs)
}
