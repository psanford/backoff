package backoff

import (
	"testing"
	"time"
)

func TestBackoffNoJitter(t *testing.T) {
	boff := New(1*time.Second, 10*time.Second, WithNoJitter())

	expect := []time.Duration{
		1 * time.Second,
		2 * time.Second,
		4 * time.Second,
		8 * time.Second,
		10 * time.Second,
		10 * time.Second,
	}

	for i, expectD := range expect {
		got := boff.Next()
		if got != expectD {
			t.Errorf("%d: expected=%s got=%s", i, expectD, got)
		}
	}

	boff.Reset()

	for i, expectD := range expect {
		got := boff.Next()
		if got != expectD {
			t.Errorf("%d: expected=%s got=%s", i, expectD, got)
		}
	}
}

func TestBackoffWithJitter(t *testing.T) {
	boff := New(1*time.Second, 10*time.Second, WithNoJitter())

	expect := []struct {
		min time.Duration
		max time.Duration
	}{
		{
			500 * time.Millisecond,
			1500 * time.Millisecond,
		},
		{
			1000 * time.Millisecond,
			3000 * time.Millisecond,
		},
		{
			2000 * time.Millisecond,
			6000 * time.Millisecond,
		},
		{
			4000 * time.Millisecond,
			10000 * time.Millisecond,
		},
		{
			5000 * time.Millisecond,
			10000 * time.Millisecond,
		},
		{
			5000 * time.Millisecond,
			10000 * time.Millisecond,
		},
	}

	for i, expectRange := range expect {
		got := boff.Next()
		if got < expectRange.min || got > expectRange.max {
			t.Errorf("%d: expected=%s-%s got=%s", i, expectRange.min, expectRange.max, got)
		}
	}

	boff.Reset()

	for i, expectRange := range expect {
		got := boff.Next()
		if got < expectRange.min || got > expectRange.max {
			t.Errorf("%d: expected=%s-%s got=%s", i, expectRange.min, expectRange.max, got)
		}
	}
}
