package async

import (
	"fmt"
	"testing"
)

func TestExecReturnOne(t *testing.T) {
	res := ExecOne(ctx, func() int {
		fmt.Println("executes function")
		return 1
	}).Await(ctx)
	fmt.Println("function done! result: ", res)
}

func TestExecAllOne(t *testing.T) {
	var fns []func() int
	for i := 0; i < 5; i++ {
		i := i
		fns = append(fns, func() int {
			fmt.Println("executes function ", i)
			return i
		})
	}
	res := ExecAllOne(ctx, fns).Await(ctx)
	fmt.Println("all functions done! result: ", res)
}
