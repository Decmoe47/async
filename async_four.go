package async

import (
	"context"
	"time"
)

type resultFour[T1, T2, T3, T4 any] struct {
	value1 T1
	value2 T2
	value3 T3
	value4 T4
}

// ExecReturnFour executes the function which returns four values in a separate goroutine and returns a future to await
func ExecReturnFour[T1, T2, T3, T4 any](ctx context.Context, fn func() (T1, T2, T3, T4)) *FutureFour[T1, T2, T3, T4] {
	result := make(chan resultFour[T1, T2, T3, T4])
	go func() {
		v1, v2, v3, v4 := fn()
		select {
		case <-ctx.Done():
			return
		default:
			result <- resultFour[T1, T2, T3, T4]{value1: v1, value2: v2, value3: v3, value4: v4}
		}
	}()

	return &FutureFour[T1, T2, T3, T4]{
		Await: func(ctx context.Context) (res1 T1, res2 T2, res3 T3, res4 T4) {
			for {
				select {
				case <-ctx.Done():
					return
				case res := <-result:
					return res.value1, res.value2, res.value3, res.value4
				default:
					time.Sleep(time.Millisecond)
					continue
				}
			}
		},
	}
}

// ExecFour is the alias of ExecReturnFour
func ExecFour[T1, T2, T3, T4 any](ctx context.Context, fn func() (T1, T2, T3, T4)) *FutureFour[T1, T2, T3, T4] {
	return ExecReturnFour(ctx, fn)
}

// ExecAllReturnFour executes all functions which return four values in separate goroutines and returns a future to await
func ExecAllReturnFour[T1, T2, T3, T4 any](
	ctx context.Context,
	fns []func() (T1, T2, T3, T4),
) *FutureFour[[]T1, []T2, []T3, []T4] {
	result := make(chan resultFour[T1, T2, T3, T4], len(fns))
	for _, fn := range fns {
		fn := fn
		go func() {
			v1, v2, v3, v4 := fn()
			select {
			case <-ctx.Done():
				return
			default:
				result <- resultFour[T1, T2, T3, T4]{value1: v1, value2: v2, value3: v3, value4: v4}
			}
		}()
	}

	return &FutureFour[[]T1, []T2, []T3, []T4]{
		Await: func(ctx context.Context) (res1 []T1, res2 []T2, res3 []T3, res4 []T4) {
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
					res4 = append(res4, r.value4)
					doneCount++
				default:
					if doneCount == fnsLen {
						return
					}
					time.Sleep(time.Microsecond * 10)
				}
			}
		},
	}
}

// ExecAllFour is the alias of ExecAllReturnFour
func ExecAllFour[T1, T2, T3, T4 any](
	ctx context.Context,
	fns []func() (T1, T2, T3, T4),
) *FutureFour[[]T1, []T2, []T3, []T4] {
	return ExecAllReturnFour(ctx, fns)
}
