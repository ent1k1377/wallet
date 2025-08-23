package app

import (
	"context"
	"fmt"
	"strings"
	"sync"
)

type Func func(ctx context.Context) error

type Closer struct {
	mu    sync.Mutex
	funcs []Func
}

func (c *Closer) Add(f Func) {
	c.funcs = append(c.funcs, f)
}

func (c *Closer) Close(ctx context.Context) error {
	msgs := make([]string, 0, len(c.funcs))
	
	var wg sync.WaitGroup
	for _, f := range c.funcs {
		wg.Add(1)
		go func(f Func) {
			if err := f(ctx); err != nil {
				c.mu.Lock()
				msgs = append(msgs, fmt.Sprintf("[!] %v", err.Error()))
				c.mu.Unlock()
			}
			wg.Done()
		}(f)
	}
	
	wg.Wait()
	
	if len(msgs) > 0 {
		return fmt.Errorf("shutdown finished with error(s): \n%s", strings.Join(msgs, "\n"))
	}
	
	return nil
}
