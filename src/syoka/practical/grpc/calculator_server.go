package grpc

import (
	"context"
	"fmt"
)

type Calculator struct {
}

func (s *Calculator) Add(ctx context.Context, ele *Element) (*Result, error) {
	total := ele.GetA() + ele.GetB()
	fmt.Printf("result:%d", total)
	return &Result{Result: int64(total)}, nil
}
