package timeline_test

import (
	"fmt"

	"github.com/eihigh/timeline"
)

type TL = timeline.Timeline[int]

func Example() {
	// Span: Execute callbacks only within defined time ranges
	fmt.Println("=== Span Example ===")
	for t := range 8 {
		timeline.New(t).
			Span(3, func(tl TL) {
				fmt.Printf("t=%d: In span [0,3), elapsed=%d\n", t, tl.Elapsed())
			}).
			Span(3, func(tl TL) {
				fmt.Printf("t=%d: In span [3,6), elapsed=%d\n", t, tl.Elapsed())
			})
	}

	// LoopN: Execute callbacks n times at fixed intervals
	fmt.Println("=== LoopN Example ===")
	for t := range 10 {
		timeline.New(t).
			LoopN(4, 2, func(n int, tl TL) { // Every 4 units, repeat 2 times
				fmt.Printf("t=%d: Loop %d, elapsed=%d, ratio=%v\n", t, n, tl.Elapsed(), tl.Ratio())
			})
	}

	// Loop: Execute callbacks indefinitely at fixed intervals
	fmt.Println("=== Loop Example ===")
	for t := range 8 {
		timeline.New(t).
			Loop(3, func(n int, tl TL) { // Every 3 units, repeat forever
				fmt.Printf("t=%d: Infinite loop %d, elapsed=%d\n", t, n, tl.Elapsed())
			})
	}

	// Output:
	// === Span Example ===
	// t=0: In span [0,3), elapsed=0
	// t=1: In span [0,3), elapsed=1
	// t=2: In span [0,3), elapsed=2
	// t=3: In span [3,6), elapsed=0
	// t=4: In span [3,6), elapsed=1
	// t=5: In span [3,6), elapsed=2
	// === LoopN Example ===
	// t=0: Loop 0, elapsed=0, ratio=0
	// t=1: Loop 0, elapsed=1, ratio=0.25
	// t=2: Loop 0, elapsed=2, ratio=0.5
	// t=3: Loop 0, elapsed=3, ratio=0.75
	// t=4: Loop 1, elapsed=0, ratio=0
	// t=5: Loop 1, elapsed=1, ratio=0.25
	// t=6: Loop 1, elapsed=2, ratio=0.5
	// t=7: Loop 1, elapsed=3, ratio=0.75
	// === Loop Example ===
	// t=0: Infinite loop 0, elapsed=0
	// t=1: Infinite loop 0, elapsed=1
	// t=2: Infinite loop 0, elapsed=2
	// t=3: Infinite loop 1, elapsed=0
	// t=4: Infinite loop 1, elapsed=1
	// t=5: Infinite loop 1, elapsed=2
	// t=6: Infinite loop 2, elapsed=0
	// t=7: Infinite loop 2, elapsed=1
}

func ExampleTimeline_Ratio() {
	for t := range 15 {
		timeline.New(t).
			Span(5, func(tl TL) {
				fmt.Print(tl.Ratio(), " ")
			}).
			Span(10, func(tl TL) {
				fmt.Print(tl.Ratio(), " ")
			})
	}

	// Output:
	// 0 0.2 0.4 0.6 0.8 0 0.1 0.2 0.3 0.4 0.5 0.6 0.7 0.8 0.9
}

func Example_nested() {
	// Nested timelines allow complex timing patterns
	fmt.Println("=== Nested Timeline Example ===")
	for t := range 12 {
		tl := timeline.New(t)
		tl.
			// Outer timeline: 3 spans of 6 units each
			Span(6, func(tl TL) {
				fmt.Printf("t=%d: Outer span 1, ", t)
				// Inner timeline: 2 loops within the outer span
				tl.
					Loop(3, func(n int, tl TL) {
						fmt.Printf("inner loop %d (elapsed=%d)\n", n, tl.Elapsed())
					})
			}).
			Span(6, func(tl TL) {
				fmt.Printf("t=%d: Outer span 2, ", t)
				// Inner timeline: spans within the outer span
				tl.
					Span(2, func(tl TL) {
						fmt.Printf("inner span A (ratio=%.1f)\n", tl.Ratio())
					}).
					Span(2, func(tl TL) {
						fmt.Printf("inner span B (ratio=%.1f)\n", tl.Ratio())
					}).
					Span(2, func(tl TL) {
						fmt.Printf("inner span C (ratio=%.1f)\n", tl.Ratio())
					})
			})
	}

	// Output:
	// === Nested Timeline Example ===
	// t=0: Outer span 1, inner loop 0 (elapsed=0)
	// t=1: Outer span 1, inner loop 0 (elapsed=1)
	// t=2: Outer span 1, inner loop 0 (elapsed=2)
	// t=3: Outer span 1, inner loop 1 (elapsed=0)
	// t=4: Outer span 1, inner loop 1 (elapsed=1)
	// t=5: Outer span 1, inner loop 1 (elapsed=2)
	// t=6: Outer span 2, inner span A (ratio=0.0)
	// t=7: Outer span 2, inner span A (ratio=0.5)
	// t=8: Outer span 2, inner span B (ratio=0.0)
	// t=9: Outer span 2, inner span B (ratio=0.5)
	// t=10: Outer span 2, inner span C (ratio=0.0)
	// t=11: Outer span 2, inner span C (ratio=0.5)
}
