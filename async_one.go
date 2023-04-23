package async

import (
	"context"
	"time"
)

// ExecReturnOne executes the function which returns one value in a separate goroutine and returns a future to await
func ExecReturnOne[T any](ctx context.Context, fn func() T) *FutureOne[T] {
	result := make(chan T)
	go func() {
		v := fn()
		select {
		case <-ctx.Done():
			return
		default:
			result <- v
		}
	}()

	return &FutureOne[T]{
		Await: func(ctx context.Context) (res T) {
			for {
				select {
				case <-ctx.Done():
					return
				case res = <-result:
					return
				default:
					time.Sleep(time.Millisecond)
					continue
				}
			}
		},
	}
}

// ExecOne is the alias of ExecReturnOne
func ExecOne[T any](ctx context.Context, fn func() T) *FutureOne[T] {
	return ExecReturnOne(ctx, fn)
}

// ExecAllReturnOne executes all functions which return one value in separate goroutines and returns a future to await
func ExecAllReturnOne[T any](ctx context.Context, fns []func() T) *FutureOne[[]T] {
	result := make(chan T, len(fns))
	for _, fn := range fns {
		fn := fn
		go func() {
			v := fn()
			select {
			case <-ctx.Done():
				return
			default:
				result <- v
			}
		}()
	}

	return &FutureOne[[]T]{
		Await: func(ctx context.Context) (res []T) {
			var (
				doneCount int
				fnsLen    = len(fns)
			)
			for {
				select {
				case <-ctx.Done():
					return
				case r := <-result:
					res = append(res, r)
					doneCount++
				default:
					time.Sleep(time.Millisecond)
					continue
				}
				if doneCount == fnsLen {
					return
				}
			}
		},
	}
}

// ExecAllOne is the alias of ExecAllReturnOne
func ExecAllOne[T any](ctx context.Context, fns []func() T) *FutureOne[[]T] {
	return ExecAllReturnOne(ctx, fns)
}
