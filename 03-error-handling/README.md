# Task 3: Error Handling - File-Based Contact Manager

## Concepts Covered
- Multiple return values with errors
- The `error` interface
- Creating errors with `errors.New()` and `fmt.Errorf()`
- Error wrapping with `%w`
- Error unwrapping with `errors.Is()` and `errors.As()`
- Custom error types
- Error handling patterns

## What I'm Building
A contact management system with file persistence that demonstrates comprehensive error handling in Go.

## Features
- Add, view, update, delete contacts
- Persist contacts to JSON file
- Validate contact data
- Custom error types for different failure scenarios
- Error wrapping for context
- User-friendly error messages

## Key Design Decisions
- Always Check Errors Immediately
- Keep error messages simple and lowercase
- return error as last value

## What I Learned
- Any type that implements `Error() string` is an error
- Sentinel Errors - package level errors with var
- `%w` Preserves original error, creates error chain
- `error.Is()` used to check error matches
- `error.As()` used to extract error types, must pass pointer

## Challenges Faced
- Mannual handling or error

## How to Run
```bash
cd 03-error-handling
go run main.go
```

## Sample Output
[ ]