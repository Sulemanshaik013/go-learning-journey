# Task 2: Structs, Methods, Receivers & Interfaces - Library Management System

## Concepts Covered
- Defining structs with multiple fields
- Value receivers vs pointer receivers
- Methods attached to types
- Interfaces and polymorphism
- Type assertions
- Composition

## What I'm Building
A library management system that tracks books, members, and borrowing operations using structs, methods, and interfaces.

## Features
- Add books and members to library
- Borrow and return books
- Track which member has which book
- Enforce borrowing rules (max 3 books per member)
- List available books
- Display member information
- Use interfaces for polymorphic behavior

## Key Design Decisions
- Group releated data together logically
- If any method need pointer recceiver, use pointer recceiver for all to maintain consitency.
- small interfaces and behavior based naming for interfaces

## What I Learned
- structs and interfaces
- receivers (pointer & value)
- when to use pointer receiver and value receiver
- no implement keyword, struct has to have all the methods defined in interface,
- items in range loop are copy 
## Challenges Faced
- delete opertion is slice is mannual
- understanding pointer copy while iterating list
- implict implementation of interfaces
## How to Run
```bash
cd 02-structs-methods-interfaces
go run main.go
```

## Sample Output
[ ]
