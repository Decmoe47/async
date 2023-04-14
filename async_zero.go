package async

import "context"

// Exec executes the function in a separate goroutine and returns a future to await
func Exec(ctx context.Context, fn func()) *Future {
	done := make(chan struct{})
	go func() {
		fn()
		select {
		case <-ctx.Done():
			return
		default:
			done <- struct{}{}
		}
	}()

	return &Future{
		Await: func(ctx context.Context) {
			for {
				select {
				case <-ctx.Done():
					return
				case <-done:
					return
				default:
					continue
				}
			}
		},
	}
}

// ExecAll executes all functions in separate goroutines and returns a future to await
func ExecAll(ctx context.Context, fns []func()) *Future {
	done := make(chan struct{}, len(fns))
	for _, fn := range fns {
		fn := fn
		go func() {
			fn()
			select {
			case <-ctx.Done():
				return
			default:
				done <- struct{}{}
			}
		}()
	}

	return &Future{
		Await: func(ctx context.Context) {
			var (
				doneCount int
				fnsLen    = len(fns)
			)
			for {
				select {
				case <-ctx.Done():
					return
				case <-done:
					doneCount++
				default:
					continue
				}
				if doneCount == fnsLen {
					return
				}
			}
		},
	}
}
