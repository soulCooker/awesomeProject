package stack

type MinStack struct {
	Vals    []int
	MinVals []int
}

func Constructor() MinStack {
	return MinStack{
		Vals:    make([]int, 0),
		MinVals: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	minVal := val
	if len(this.MinVals) > 0 {
		minVal = min(minVal, this.MinVals[len(this.MinVals)-1])
	}
	this.MinVals = append(this.MinVals, minVal)
	this.Vals = append(this.Vals, val)
}

func (this *MinStack) Pop() {
	this.MinVals = this.MinVals[:len(this.MinVals)-1]
	this.Vals = this.Vals[:len(this.Vals)-1]
}

func (this *MinStack) Top() int {
	return this.Vals[len(this.Vals)-1]
}

func (this *MinStack) GetMin() int {
	return this.MinVals[len(this.MinVals)-1]
}

func min(v1, v2 int) int {
	if v1 > v2 {
		return v2
	} else {
		return v1
	}
}
