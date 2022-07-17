# 115

### wall clock and monotonic clock

Operating systems provide both a “wall clock”, which is subject to changes for clock synchronization, and a “monotonic clock”, which is not. The general rule is that the wall clock is for telling time and the monotonic clock is for measuring time. Rather than split the API, in this package the Time returned by time.Now contains both a wall clock reading and a monotonic clock reading; later time-telling operations use the wall clock reading, but later time-measuring operations, specifically comparisons and subtractions, use the monotonic clock reading.

### Timer and Ticker

The Timer type represents a single event.

A Ticker holds a channel that delivers “ticks” of a clock at intervals.
