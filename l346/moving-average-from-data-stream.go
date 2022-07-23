package l346

type MovingAverage struct {
	size  int
	acc   float64
	nums  []float64
	t, h  int
	count int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{
		size: size,
		nums: make([]float64, size+1),
	}
}

func (this *MovingAverage) Next(val int) float64 {
	if this.count < this.size {
		this.count++
	} else {
		this.acc -= this.nums[this.t]
		this.t = (this.t + 1) % len(this.nums)
	}

	this.nums[this.h] = float64(val)
	this.h = (this.h + 1) % len(this.nums)
	this.acc += float64(val)
	return this.acc / float64(this.count)
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
