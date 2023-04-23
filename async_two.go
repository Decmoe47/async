package async

import (
	"context"
	"time"
)

type resultTwo[T1, T2 any] struct {
	value1 T1
	value2 T2
}

// ExecReturnTwo executes the function which returns two values in a separate goroutine and returns a future to await
func ExecReturnTwo[T1, T2 any](ctx context.Context, fn func() (T1, T2)) *FutureTwo[T1, T2] {
	result := make(chan resultTwo[T1, T2])
	go func() {
		v1, v2 := fn()
		select {
		case <-ctx.Done():
			return
		default:
			result <- resultTwo[T1, T2]{value1: v1, value2: v2}
		}
	}()

	return &FutureTwo[T1, T2]{
		Await: func(ctx context.Context) (res1 T1, res2 T2) {
			for {
				select {
				case <-ctx.Done():
					return
				case res := <-result:
					return res.value1, res.value2
				default:
					time.Sleep(time.Millisecond)
					continue
				}
			}
		},
	}
}

// ExecTow is the alias of ExecReturnTwo
func ExecTwo[T1, T2 any](ctx context.Context, fn func() (T1, T2)) *FutureTwo[T1, T2] {
	return ExecReturnTwo(ctx, fn)
}

// ExecAllReturnTwo executes all functions which return two values in separate goroutines and returns a future to await
func ExecAllReturnTwo[T1, T2 any](ctx context.Context, fns []func() (T1, T2)) *FutureTwo[[]T1, []T2] {
	result := make(chan resultTwo[T1, T2], len(fns))
	for _, fn := range fns {
		fn := fn
		go func() {
			v1, v2 := fn()
			select {
			case <-ctx.Done():
				return
			default:
				result <- resultTwo[T1, T2]{value1: v1, value2: v2}
			}
		}()
	}

	return &FutureTwo[[]T1, []T2]{
		Await: func(ctx context.Context) (res1 []T1, res2 []T2) {
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

// ExecAllTwo is the alias of ExecAllReturnTwo
func ExecAllTwo[T1, T2 any](ctx context.Context, fns []func() (T1, T2)) *FutureTwo[[]T1, []T2] {
	return ExecAllReturnTwo(ctx, fns)
}
