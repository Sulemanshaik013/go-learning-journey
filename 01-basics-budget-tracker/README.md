# Task 1: Go Basics - Budget Tracker

## Concepts Covered
- Variables and basic types
- Slices and maps
- Loops and conditionals
- Functions

## What I Built
A command-line budget tracker that manages income and expenses.

## Features
- Add transactions
- Calculate totals
- Categorize spending
- Show summary

## What I Learned
- Every Go file starts with a package declaration, Go is statically typed.
- `main` package is special, it's the entry point for executable programs.
- `fmt` is the formatting package for printing output.
- Structs group related data together (similar to structures in C programming).
- Field names starting with `lowercase` are unexported (private to the package), `capitalcase` are exported (public).
- Slices are dynamic arrays that can grow, `append()` adds items and returns a new slice.
- Maps are key-value pairs, `map[string]float64` means "keys are strings, values are float64".
- `make()` initializes the map, `+=` to accumulate values.
- `range` iterates over slices and maps, returns index and value (we ignore index with `_`).
- Standard if/else syntax, no parentheses needed around conditions and
`==` for equality comparison.
- Functions can return `multiple` values, return types are specified after `parameters`.
    ```go
    func calculateTotals(transactions []Transaction) (float64, float64, float64) {
        // ...
        return totalIncome, totalExpenses, netBalance
    }
    ```
- `var` for explicit declaration, `:=` for short declaration with type inference.

## Challenges Faced
- bit complex in writing, more lines of code
- no built methods
- adapting to the syntax
- declaration and initialization
## How to Run

```bash
go run main.go
```