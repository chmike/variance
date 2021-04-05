# Computes the mean or variance

This small package can be used to compute the mean or variance of 
values in one pass. Values are added with the `Add` method. 
The `Mean` method returns the mean of the values. If no values
have been adde `Mean` returns 0. The method `Variance` returns 
the variance of the added values. If there are less then two
values in the accumulator, `Variance` returns 0.

The `Del` method can be used to remove an added value. It panics
if no values have been added or all values have been removed.
This method can be used to compute the mean and variance over
a sliding window. Values are added on one side, and removed on 
the other.

```
import "github.com/chmike/variance"

...

var acc variance.Accumulator
acc.Add(0)
acc.Add(1)
acc.Add(2)
fmt.Println(acc.Mean())     // outputs 1.
fmt.Println(acc.Variance()) // outputs 1.
acc.Del(0)
fmt.Println(acc.Mean())     // outputs 1.5
fmt.Println(acc.Variance()) // outputs 0.5
```