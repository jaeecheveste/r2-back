package app

import (
	"fmt"
	"time"

	"github.com/jaeecheveste/r2-back/pkg/domain"

	"go.uber.org/zap"
)

type FibonacciService struct {
	log *zap.SugaredLogger
}

func NewFibonacciService(log *zap.SugaredLogger) domain.FibonacciService {
	return FibonacciService{
		log: log,
	}
}

func (fs FibonacciService) GetSpiral(rows int32, cols int32) (*domain.Spiral, error) {
	sp := make([][]int64, rows)
	for i := range sp {
		sp[i] = make([]int64, cols)
	}

	c := make(chan int64)
	n := int(rows * cols)
	fmt.Println(n)

	go fibonacciGenerator(c)

	x := 0
	y := 0

	direction := "right"
	right := int(cols) - 1
	bottom := int(rows) - 1
	left := 0
	top := 0
	for i := 0; i < n; i++ {
		val := <-c
		sp[x][y] = val

		if direction == "right" {
			if y < right {
				y++
			} else {
				direction = "bottom"
				right--
				top++
				x++
			}

			continue
		}

		if direction == "bottom" {
			if x < bottom {
				x++
			} else {
				direction = "left"
				bottom--
				y--
			}

			continue
		}

		if direction == "left" {
			if y > left {
				y--
			} else {
				direction = "top"
				x--
				left++
			}

			continue
		}

		if direction == "top" {
			if x > top {
				x--
			} else {
				direction = "right"
				top++
				y++
			}
		}

	}

	spiral := &domain.Spiral{
		Ts:   time.Now(),
		Rows: sp,
	}

	return spiral, nil
}

func fibonacciGenerator(c chan<- int64) {
	a, b := int64(0), int64(1)

	for {
		c <- a

		a, b = b, a
		a = a + b
	}
}
