package design_pattern

import "fmt"

/**
基于组合的方式实现运行时动态选择某些子步骤的具体实现
*/

type OrderPayContext struct {
	orderCreateStrategy IStrategy
}

func (receiver *OrderPayContext) createOrder() {
	fmt.Println("其他处理...")
	receiver.orderCreateStrategy.doCreateOrder()
	fmt.Println("其他处理...")
}

func (receiver *OrderPayContext) setStrategy(strategy IStrategy) {
	receiver.orderCreateStrategy = strategy
}

type IStrategy interface {
	doCreateOrder()
}

type LumpSumOrderStrategy struct {
}

func (l LumpSumOrderStrategy) doCreateOrder() {
	println("this is LumpSumOrderStrategy")
}

func test() {
	orderPayContext := &OrderPayContext{}
	orderPayContext.setStrategy(LumpSumOrderStrategy{})
	orderPayContext.createOrder()
}
