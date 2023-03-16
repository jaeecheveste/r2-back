package domain

import (
	"time"
)

type Spiral struct {
	Rows [][]int64 `json:"rows" bson:"rows,omitempty"`
	Ts   time.Time `json:"ts" bson:"ts,omitempty"`
}

type FibonacciService interface {
	GetSpiral(rows int32, cols int32) (*Spiral, error)
}
