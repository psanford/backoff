package backoff

import (
	"math/rand"
	"time"
)

type Backoff struct {
	initial time.Duration
	max     time.Duration

	cur        time.Duration
	jitterFunc func(time.Duration) time.Duration
}

// Returns a new Backoff with an initial and maximum backoff
// values.
func New(initial, max time.Duration, opts ...Option) *Backoff {
	if max < initial {
		panic("backoff: max must be >= initial")
	}

	var backoffOptions options
	for _, opt := range opts {
		opt.setOption(&backoffOptions)
	}

	b := Backoff{
		initial: initial,
		cur:     initial,
		max:     max,

		jitterFunc: boundHalfJitter,
	}

	if backoffOptions.disableJitter {
		b.jitterFunc = noJitter
	}

	return &b
}

func boundHalfJitter(d time.Duration) time.Duration {
	half := int64(d / 2)
	return time.Duration(half + rand.Int63n(half))
}

func noJitter(d time.Duration) time.Duration {
	return d
}

// Returns the next delay to apply
func (b *Backoff) Next() time.Duration {
	backoff := b.jitterFunc(b.cur)

	next := b.cur * 2

	if next > b.max {
		b.cur = b.max
	} else {
		b.cur = next
	}

	return backoff
}

// Resets back to the initial value
func (b *Backoff) Reset() {
	b.cur = b.initial
}

type options struct {
	disableJitter bool
}

type Option interface {
	setOption(*options)
}

type noJitterOpt struct {
}

func (o noJitterOpt) setOption(opts *options) {
	opts.disableJitter = true
}

// Disables jitter
func WithNoJitter() Option {
	return &noJitterOpt{}
}
