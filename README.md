# timeline

[![Go Reference](https://pkg.go.dev/badge/github.com/eihigh/timeline.svg)](https://pkg.go.dev/github.com/eihigh/timeline)

A Go library for time-based sequencing of animations and events.

## Overview

Timeline provides a fluent API for scheduling callbacks based on elapsed time. It supports sequential spans, finite loops, and infinite loops with progress tracking.

## Usage

### Basic Timeline

Create a timeline with current time:

```go
tl := timeline.New(currentTime)
```

### Sequential Spans

Execute callbacks within specific time ranges:

```go
timeline.New(t).
    Span(3, func(tl Timeline[int]) {
        // Executes when t is in [0,3)
        fmt.Println("First span, elapsed:", tl.Elapsed())
    }).
    Span(5, func(tl Timeline[int]) {
        // Executes when t is in [3,8)
        fmt.Println("Second span, elapsed:", tl.Elapsed())
    })
```

### Loops

Repeat callbacks at fixed intervals:

```go
// Finite loop (2 iterations)
timeline.New(t).
    LoopN(4, 2, func(n int, tl Timeline[int]) {
        // Executes at t=0,1,2,3 (n=0) and t=4,5,6,7 (n=1)
        fmt.Printf("Loop %d, progress: %.0f%%\n", n, tl.Ratio()*100)
    })

// Infinite loop
timeline.New(t).
    Loop(3, func(n int, tl Timeline[int]) {
        // Executes every 3 time units forever
        fmt.Printf("Cycle %d\n", n)
    })
```

### Progress Tracking

Get elapsed time and progress ratio:

```go
elapsed := tl.Elapsed()        // Time since segment start
elapsedF := tl.Elapsedf()     // As float64
ratio := tl.Ratio()           // Progress (0.0 to 1.0)
```

### One-time Events

Execute callbacks at specific time points:

```go
timeline.New(t).
    Once(func() {
        // Executes when t=0 (initialization)
    }).
    Span(3, func(tl Timeline[int]) {
        // Executes during t=0,1,2
    }).
    Once(func() {
        // Executes when t=3 (transition)
    })
```

**Note**: `Once` only executes when time exactly matches the current position. It may not work reliably with variable time steps.

### Nested Timelines

Create complex timing patterns by nesting timelines:

```go
timeline.New(t).
    Span(6, func(tl Timeline[int]) {
        // Create inner timeline based on outer timeline's elapsed time
        tl.
            LoopN(3, 2, func(n int, tl Timeline[int]) {
                // Inner timeline executes within outer span
                fmt.Printf("Inner loop %d\n", n)
            })
    })
```

This allows for sophisticated animation sequences where inner timelines can have their own independent timing logic while being synchronized with outer timelines.

## Examples

See [example_test.go](example_test.go) for complete examples.

## License

MIT
