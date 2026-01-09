# Pointers Learning & Exercises

A dedicated folder for mastering pointer manipulation in Go.

## Exercises Included
1. **Basic Swap**: Write a function that swaps two integers using pointers.
2. **Array Modification**: Pass an array (or slice) pointer to a function and modify its elements.
3. **Struct Updates**: Use pointer receivers to update struct fields.

## Key Concepts to Remember
- `&x` is the address.
- `*p` is the value.
- Pointers are `nil` by default.
- Go has **no pointer arithmetic** (you can't do `p++` to move to the next memory address like in C).

## Why this is important
Understanding pointers is the difference between a Go beginner and a Go professional. It allows you to write high-performance code and understand how Go's "under the hood" memory management works.
