// Package timeline provides time-based sequencing for animations and events.
package timeline

import (
	"github.com/eihigh/ng"
)

// Timeline is a builder for creating time-based sequences of animations and events.
type Timeline[T ng.Int] struct {
	from, to, now T
}

// New creates a timeline starting at 0 with the given current time.
func New[T ng.Int](now T) Timeline[T] {
	return Timeline[T]{from: 0, now: now}
}

// Elapsed returns time passed since start.
func (t Timeline[T]) Elapsed() T {
	return t.now - t.from
}

// ElapsedF returns elapsed time as float64.
func (t Timeline[T]) ElapsedF() float64 {
	return float64(t.now - t.from)
}

// Ratio returns progress within the current span (0.0 to 1.0).
func (t Timeline[T]) Ratio() float64 {
	denom := t.to - t.from
	if denom == 0 {
		return 0
	}
	return float64(t.now-t.from) / float64(denom)
}

// Span defines a time segment and executes callbacks within it.
func (t Timeline[T]) Span(duration T, f ...func(Timeline[T])) Timeline[T] {
	t.to = t.from + duration
	if t.from <= t.now && t.now < t.to {
		for _, fn := range f {
			fn(t)
		}
	}
	return Timeline[T]{
		from: t.to,
		to:   t.to,
		now:  t.now,
	}
}

// Loop repeats callbacks indefinitely at fixed intervals.
func (t Timeline[T]) Loop(duration T, f ...func(int, Timeline[T])) {
	if t.now < t.from {
		return
	}
	n := (t.now - t.from) / duration
	for _, fn := range f {
		fn(int(n), Timeline[T]{
			from: t.from + n*duration,
			to:   t.from + (n+1)*duration,
			now:  t.now,
		})
	}
}

// LoopN repeats callbacks n times at fixed intervals.
func (t Timeline[T]) LoopN(duration T, n int, f ...func(int, Timeline[T])) Timeline[T] {
	m := (t.now - t.from) / duration
	if 0 <= m && m < T(n) {
		for _, fn := range f {
			fn(int(m), Timeline[T]{
				from: t.from + m*duration,
				to:   t.from + (m+1)*duration,
				now:  t.now,
			})
		}
	}

	return Timeline[T]{
		from: t.from + T(n)*duration,
		to:   t.from + T(n)*duration,
		now:  t.now,
	}
}

// Once executes callbacks at the current time point.
// This is useful for one-time events like playing sounds or initializing states.
//
// Note: This method only executes when t.now exactly equals t.from.
// It may not work reliably with variable time steps or when frames are skipped.
// For reliable execution, ensure time increments by consistent steps (e.g., 1 per frame).
func (t Timeline[T]) Once(f ...func()) Timeline[T] {
	if t.now == t.from {
		for _, fn := range f {
			fn()
		}
	}
	return t
}
