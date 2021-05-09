package congroup_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/zhjp0/goes/congroup"
	"golang.org/x/time/rate"
)

func TestWithContext(t *testing.T) {
	g, ctx := congroup.WithContext(context.Background())
	g.SetLimiter(rate.NewLimiter(rate.Limit(2), 1))

	n := 10
	for i := 0; i < n; i++ {
		i := i
		g.Go(func() error {
			fmt.Printf("%v i = %d start\n", time.Now(), i)
			// time.Sleep(time.Second)
			// time.Sleep(time.Duration(float64(time.Second*3) * (float64(i) / float64(n))))
			fmt.Printf("%v i = %d end\n", time.Now(), i)
			return fmt.Errorf("mock error")
		})
	}
	if err := g.Wait(ctx); err != nil {
		fmt.Println(err)
	}
}
