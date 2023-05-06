package async

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var ctx = context.Background()

func TestExec(t *testing.T) {
	Exec(ctx, func() {
		fmt.Println("executes function")
	}).Await(ctx)
	fmt.Println("function done!")
}

func TestExecAll(t *testing.T) {
	var fns []func()
	for i := 0; i < 5; i++ {
		i := i
		fns = append(fns, func() {
			fmt.Println("executes function ", i)
			time.Sleep(time.Second * time.Duration(rand.Intn(10)))
		})
	}
	ExecAll(ctx, fns).Await(ctx)
	fmt.Println("all functions done!")
}
