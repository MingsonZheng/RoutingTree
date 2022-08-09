package main

import (
	"fmt"
	"time"
)

type FilterBuilder func(next Filter) Filter

type Filter func(c *Context)

// 类型断言
var _ FilterBuilder = MetricsFilterBuilder

func MetricsFilterBuilder(next Filter) Filter {
	return func(c *Context) {
		start := time.Now().Nanosecond()
		next(c)
		end := time.Now().Nanosecond()
		fmt.Printf("用了 %d 纳秒", end-start)
	}
}
