package async

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestExecReturnThree(t *testing.T) {
	res1, res2, res3 := ExecThree(ctx, func() (int, string, bool) {
		fmt.Println("executes function")
		return 1, "a", false
	}).Await(ctx)
	fmt.Println("function done! result: ", res1, res2, res3)
}

func TestExecAllThree(t *testing.T) {
	var fns []func() (int, string, bool)
	for i := 0; i < 5; i++ {
		i := i
		fns = append(fns, func() (int, string, bool) {
			fmt.Println("executes function ", i)
			time.Sleep(time.Second * time.Duration(rand.Intn(10)))
			return i, strconv.Itoa(i + 5), i%2 == 0
		})
	}
	res1, res2, res3 := ExecAllThree(ctx, fns).Await(ctx)
	fmt.Println("all functions done! result: ", res1, res2, res3)
}
