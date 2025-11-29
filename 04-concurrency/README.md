# Task 4: Concurrency - Parallel Web Scraper

## Concepts Covered
- Goroutines and concurrent execution
- Channels for communication
- Buffered vs unbuffered channels
- Channel directions and closing
- sync.WaitGroup for coordination
- sync.Mutex for shared data
- Select statement for multiplexing
- Worker pool pattern
- Fan-out/fan-in pattern
- Timeouts and cancellation

## What I'm Building
A concurrent web scraper that fetches multiple URLs in parallel using goroutines, channels, and synchronization primitives.

## Features
- Concurrent URL fetching with worker pool
- Limited concurrency (configurable workers)
- Error handling across goroutines
- Statistics collection with mutex
- Timeout handling with select
- Graceful shutdown
- Result aggregation

## Key Design Decisions
[ ]

## What I Learned
[ ]

## Challenges Faced
[ ]

## How to Run
```bash
cd 04-concurrency
go run main.go

# To check for race conditions:
go run -race main.go
```