package async

import "context"

type Future struct {
	Await func(ctx context.Context)
}

type FutureOne[T any] struct {
	Await func(ctx context.Context) T
}

type FutureTwo[T1, T2 any] struct {
	Await func(ctx context.Context) (T1, T2)
}

type FutureThree[T1, T2, T3 any] struct {
	Await func(ctx context.Context) (T1, T2, T3)
}

type FutureFour[T1, T2, T3, T4 any] struct {
	Await func(ctx context.Context) (T1, T2, T3, T4)
}
