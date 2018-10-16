package ai_util

import (
	"strconv"
	"sync/atomic"
)

type AtomicInt struct {
	i int64
}

func (ai *AtomicInt) Inc() {
	atomic.AddInt64(&ai.i, 1)
}

func (ai *AtomicInt) Dec() {
	atomic.AddInt64(&ai.i, -1)
}

func (ai *AtomicInt) Set(i int64) {
	atomic.StoreInt64(&ai.i, i)
}

func (ai *AtomicInt) Get() int64 {
	return atomic.LoadInt64(&ai.i)
}

func (ai *AtomicInt) Reset(i int64) int64 {
	return atomic.SwapInt64(&ai.i, i)
}

func (ai *AtomicInt) Add(delta int64) {
	atomic.AddInt64(&ai.i, delta)
}

func (ai AtomicInt) String() string {
	return strconv.FormatInt(ai.i, 10)
}
