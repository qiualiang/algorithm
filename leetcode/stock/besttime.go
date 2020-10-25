package stock

/*
122. 买卖股票的最佳时机 II
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii
关键：持有一股，买卖无数次，当天可以买也可以卖
*/

//贪心算法，如果后一天比前一天高就在前一天买入，后一天卖出 O(N)
//DP动态规划 保存每一天买和卖不同状态的收益 O(N)
func maxProfit(prices []int) int {
	profit := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] > prices[i] { //buy
			n := prices[i+1] - prices[i]
			profit += n
		}
	}
	return profit
}
