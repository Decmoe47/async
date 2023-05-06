package async

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestExecReturnFour(t *testing.T) {
	res1, res2, res3, err := ExecFour(ctx, func() (int, string, bool, error) {
		fmt.Println("executes function")
		return 1, "a", false, errors.New("err")
	}).Await(ctx)
	fmt.Println("function done! result: ", res1, res2, res3, err)
}

func TestExecAllFour(t *testing.T) {
	var fns []func() (int, string, bool, error)
	for i := 0; i < 5; i++ {
		i := i
		fns = append(fns, func() (int, string, bool, error) {
			fmt.Println("executes function ", i)
			time.Sleep(time.Second * time.Duration(rand.Intn(10)))
			return i, strconv.Itoa(i + 5), i%2 == 0, errors.New("err")
		})
	}
	res1, res2, res3, errs := ExecAllFour(ctx, fns).Await(ctx)
	fmt.Println("all functions done! result: ", res1, res2, res3, errs)
}
