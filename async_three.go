package async

import (
	"context"
	"time"
)

type resultThree[T1, T2, T3 any] struct {
	value1 T1
	value2 T2
	value3 T3
}

// ExecReturnThree executes the function which returns three values in a separate goroutine and returns a future to await
func ExecReturnThree[T1, T2, T3 any](ctx context.Context, fn func() (T1, T2, T3)) *FutureThree[T1, T2, T3] {
	result := make(chan resultThree[T1, T2, T3])
	go func() {
		v1, v2, v3 := fn()
		select {
		case <-ctx.Done():
			return
		default:
			result <- resultThree[T1, T2, T3]{value1: v1, value2: v2, value3: v3}
		}
	}()

	return &FutureThree[T1, T2, T3]{
		Await: func(ctx context.Context) (res1 T1, res2 T2, res3 T3) {
			for {
				select {
				case <-ctx.Done():
					return
				case res := <-result:
					return res.value1, res.value2, res.value3
				default:
					time.Sleep(time.Millisecond)
					continue
				}
			}
		},
	}
}

// ExecThree is the alias of ExecReturnThree
func ExecThree[T1, T2, T3 any](ctx context.Context, fn func() (T1, T2, T3)) *FutureThree[T1, T2, T3] {
	return ExecReturnThree(ctx, fn)
}

// ExecAllReturnThree executes all functions which return three values in separate goroutines and returns a future to await
func ExecAllReturnThree[T1, T2, T3 any](ctx context.Context, fns []func() (T1, T2, T3)) *FutureThree[[]T1, []T2, []T3] {
	result := make(chan resultThree[T1, T2, T3], len(fns))
	for _, fn := range fns {
		fn := fn
		go func() {
			v1, v2, v3 := fn()
			select {
			case <-ctx.Done():
				return
			default:
				result <- resultThree[T1, T2, T3]{value1: v1, value2: v2, value3: v3}
			}
		}()
	}

	return &FutureThree[[]T1, []T2, []T3]{
		Await: func(ctx context.Context) (res1 []T1, res2 []T2, res3 []T3) {
			var (
				doneCount int
				fnsLen    = len(fns)
			)
			for {
				select {
				case <-ctx.Done():
					return
				case r := <-result:
					res1 = append(res1, r.value1)
					res2 = append(res2, r.value2)
					res3 = append(res3, r.value3)
					doneCount++
				default:
					if doneCount == fnsLen {
						return
					}
					time.Sleep(time.Millisecond)
					continue
				}
			}
		},
	}
}

// ExecAllThree is the alias of ExecAllReturnThree
func ExecAllThree[T1, T2, T3 any](ctx context.Context, fns []func() (T1, T2, T3)) *FutureThree[[]T1, []T2, []T3] {
	return ExecAllReturnThree(ctx, fns)
}
