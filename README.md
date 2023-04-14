# async

This is an implement of async/await by go. It uses generic (>= go1.18) to get returned values without any type asserts.

Here 10 patterns are supported:

- `Exec`: asynchronously executes a function 
- `ExecReturnOne`(alias `ExecOne`): asynchronously executes a function which returns one value
- `ExecReturnTwo`(alias `ExecTwo`): asynchronously executes a function which returns two value
- `ExecReturnThree`(alias `ExecThree`): asynchronously executes a function which returns three value
- `ExecReturnFour`(alias `ExecFour`): asynchronously executes a function which returns four value


- `ExecAll`: asynchronously executes multiple functions
- `ExecAllReturnOne`(alias `ExecAllOne`): asynchronously executes multiple functions which return one value
- `ExecAllReturnTwo`(alias `ExecAllTwo`): asynchronously executes multiple functions which return two value
- `ExecAllReturnThree`(alias `ExecAllThree`): asynchronously executes multiple functions which return three value
- `ExecAllReturnFour`(alias `ExecAllFour`): asynchronously executes multiple functions which return four value

And these functions can continue with `Await(ctx)` to wait goroutines done and returning values. 

This is especially useful when you want to run multiple goroutines and wait to get there returned values at the same time, just like:

```go
package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Decmoe47/async"
)

func main() {
	ctx := context.Background()
	
	var fns []func() (int, string)
	for i := 0; i < 5; i++ {
		i := i
		fns = append(fns, func() (int, string) {
			fmt.Println("executes function ", i)
			return i, strconv.Itoa(i + 5)
		})
	}
	res1, res2 := async.ExecAllTwo(ctx, fns).Await(ctx)
	fmt.Println("all functions done! result: ", res1, res2)
}
```
---

inspired from [Joker666/AsyncGoDemo](https://github.com/Joker666/AsyncGoDemo)