package async

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestExecReturnTwo(t *testing.T) {
	res1, res2 := ExecTwo(ctx, func() (int, string) {
		fmt.Println("executes function")
		return 1, "a"
	}).Await(ctx)
	fmt.Println("function done! result: ", res1, res2)
}

func TestExecAllTwo(t *testing.T) {
	var fns []func() (int, string)
	for i := 0; i < 5; i++ {
		i := i
		fns = append(fns, func() (int, string) {
			fmt.Println("executes function ", i)
			time.Sleep(time.Second * time.Duration(rand.Intn(10)))
			return i, strconv.Itoa(i + 5)
		})
	}
	res1, res2 := ExecAllTwo(ctx, fns).Await(ctx)
	fmt.Println("all functions done! result: ", res1, res2)
}
