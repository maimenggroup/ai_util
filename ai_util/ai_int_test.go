package ai_util

import (
	"testing"
)

func TestAtomicInt_Inc(t *testing.T) {
	ai := &AtomicInt{i: 0}

	for i := int64(0); i < 100; i++ {
		if ai.i != i {
			t.Errorf("AtomicInt Inc: expected %v, actually %v", i, ai.i)
		}
		ai.Inc()
	}
}

func TestAtomicInt_Dec(t *testing.T) {
	ai := &AtomicInt{i: 100}

	for i := int64(100); i >= 0; i-- {
		if ai.i != i {
			t.Errorf("AtomicInt Dec: expected %v, actually %v", i, ai.i)
		}
		ai.Dec()
	}
}

func TestAtomicInt_Add(t *testing.T) {
	ai := &AtomicInt{i: 10}

	for i := int64(0); i < 100; i++ {
		if ai.i != 10+i*2 {
			t.Errorf("AtomicInt Add: expected %v, actually %v", 10+i*2, ai.i)
		}
		ai.Add(2)
	}
}

func TestAtomicInt_Get(t *testing.T) {
	ai := &AtomicInt{i: 0}

	for i := int64(0); i < 100; i++ {
		if ai.Get() != i {
			t.Errorf("AtomicInt Get: expected %v, actually %v", i, ai.i)
		}
		ai.Inc()
	}
}

func TestAtomicInt_Set(t *testing.T) {
	ai := &AtomicInt{i: 0}

	for i := int64(0); i < 100; i++ {
		ai.Set(i * 3)
		if ai.i != i*3 {
			t.Errorf("AtomicInt Set: expected %v, actually %v", i*3, ai.i)
		}
	}
}

func TestAtomicInt_Reset(t *testing.T) {
	ai := &AtomicInt{i: -3}

	for i := int64(0); i < 100; i++ {
		if val := ai.Reset(i * 3); val != i*3-3 {
			t.Errorf("AtomicInt Reset: expected %v, actually %v", i*3-3, val)
		}
		if ai.i != i*3 {
			t.Errorf("AtomicInt Reset: expected %v, actually %v", i*3, ai.i)
		}
	}
}

func TestAtomicInt_String(t *testing.T) {
	ai := &AtomicInt{i: 100}
	t.Logf("%v", ai)
}

func TestAtomicInt_Concurrent(t *testing.T) {
	ai := &AtomicInt{i: 0}

	done := make(chan bool, 70)
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				ai.Set(int64(i))
			}
			done <- true
		}()
		go func() {
			for j := 0; j < 100; j++ {
				ai.Get()
			}
			done <- true
		}()
		go func() {
			for j := 0; j < 100; j++ {
				ai.Inc()
			}
			done <- true
		}()
		go func() {
			for j := 0; j < 100; j++ {
				ai.Dec()
			}
			done <- true
		}()
		go func() {
			for j := 0; j < 100; j++ {
				ai.Add(10)
			}
			done <- true
		}()
		go func() {
			for j := 0; j < 100; j++ {
				ai.Reset(int64(10 * j))
			}
			done <- true
		}()
		go func() {
			for j := 0; j < 100; j++ {
				ai.String()
			}
			done <- true
		}()
	}

	for i := 0; i < 70; i++ {
		<-done
	}
}
