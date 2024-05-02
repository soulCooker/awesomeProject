package greedy

/**
初始化当前买入日期、最高利润

从头开始取股票第一天的价格作为买入日期，向后遍历可能的卖出天数
如果当天价格大于买入日期
	计算利润，如果大于当前的卖出利润，更新最大利润。

如果当天价格小于买入日期，说明出现新的低点，可以作为买入日期
	更新买入日期i=当天天数
**/

func maxProfit(prices []int) int {
	return 0
}
