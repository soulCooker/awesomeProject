package stack

func dailyTemperatures(temperatures []int) []int {
	//初始化结果数组
	//对于每个日期的温度
	//用栈来存储温度连续降低的日期的温度和所在日期，同时栈顶为最近一天的温度和所在日期
	//当栈为空，将温度和所在日期进栈
	//栈非空
	//	当前问题大于栈顶温度时，出栈操作，更新栈顶日期的温度所在日期下标的值为当前日期-所在日期的差值
	//	当前温度小于栈顶温度时，进栈
	return nil
}
