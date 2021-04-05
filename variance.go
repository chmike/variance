/*
Compute a mean or variance in one pass over the values. Values may
also be removed to compute the mean or variance of a sliding window.

	import "github.com/chmike/variance"

	...

	var acc variance.Accumulator
	acc.Add(0)
	acc.Add(1)
	acc.Add(2)
	fmt.Println(acc.Mean())     // outputs 1.
	fmt.Println(acc.Variance()) // outputs 1.
	acc.Del(0)
	fmt.Println(acc.Mean())     // outputs 1.25
	fmt.Println(acc.Variance()) // outputs 1.625
*/
package variance

type Accumulator struct {
	k, n, ex, ex2 float64
}

// Reset is equivalent to removing all values added to the variance
// accumulator.
func (v *Accumulator) Reset() {
	v.k = 0
	v.n = 0
	v.ex = 0
	v.ex2 = 0
}

// Add add the value x to the accumulator.
func (v *Accumulator) Add(x float64) {
	if v.n == 0 {
		v.k = x
	}
	v.n++
	v.ex += x - v.k
	v.ex2 += (x - v.k) * (x - v.k)
}

// Del remove the value x from the accumulator. Panics
// if called on an empty accumulator. Del is intented to be
// used for sliding windows.
func (v *Accumulator) Del(x float64) {
	if v.n == 0 {
		panic("deleting from empty variance accumumator")
	}
	v.n--
	if v.n == 0 {
		v.Reset()
		return
	}
	v.ex -= x - v.k
	v.ex2 -= (x - v.k) * (x - v.k)
}

// Mean returns the mean of the added values.
func (v Accumulator) Mean() float64 {
	if v.n == 0 {
		return 0
	}
	return v.k + v.ex/v.n
}

// Variance returns the variance of the added values.
func (v Accumulator) Variance() float64 {
	if v.n < 2 {
		return 0
	}
	return (v.ex2 - (v.ex*v.ex)/v.n) / (v.n - 1)
}
