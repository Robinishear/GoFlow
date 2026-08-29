# 🚀 GoFlow — The Ultimate Go (Golang) Control Flow & Fundamentals Guide

[![Go Version](https://img.shields.io/badge/Go-1.20+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![Learning Series](https://img.shields.io/badge/Level--2-Go%20Fundamentals-orange?style=for-the-badge)]()
[![Status](https://img.shields.io/badge/Status-Completed-success?style=for-the-badge)]()

Welcome to **GoFlow**! This repository serves as a highly detailed reference and learning guide for mastering core **Golang** concepts. From basic inputs and control flow to advanced structures like slices, pointers, custom structs, and anonymous functions, this guide explains it all.

---

## 📑 Table of Contents

- [📁 Project Architecture & Directory Structure](#-project-architecture--directory-structure)
- [🧠 Core Topics & Detailed Code Explanations](#-core-topics--detailed-code-explanations)
  - [1. Anonymous Functions & IIFEs (`/`)](#1-anonymous-functions--iifes-)
  - [2. User Input & Scan (`Scan/`)](#2-user-input--scan-scan)
  - [3. Conditional Logic (`if-else/`)](#3-conditional-logic-if-else)
  - [4. Switch Statements (`Switch/`)](#4-switch-statements-switch)
  - [5. Loops & Iterations (`for-loop/`)](#5-loops--iterations-for-loop)
  - [6. Arrays (`Array/`)](#6-arrays-array)
  - [7. Slices & Dynamic Arrays (`Slice/`)](#7-slices--dynamic-arrays-slice)
  - [8. Pointers & Memory Referencing (`Pointers/`)](#8-pointers--memory-referencing-pointers)
  - [9. Structs & Custom Types (`Struct/`)](#9-structs--custom-types-struct)
- [📊 Quick Comparison: Array vs. Slice](#-quick-comparison-array-vs-slice)
- [⚡ How to Run the Code](#-how-to-run-the-code)
- [💡 Key Go Rules & Gotchas to Remember](#-key-go-rules--gotchas-to-remember)

---

## 📁 Project Architecture & Directory Structure

Here is an overview of the workspace structure. Each folder focuses on a single, isolated topic with a self-contained executable `main.go`.

```bash
GoFlow/
├── Array/
│   └── main.go       # Fixed-size arrays, indexing, and iteration
├── Pointers/
│   └── main.go       # Pointer variables, dereferencing, and pass-by-reference
├── Scan/
│   └── main.go       # Handling CLI input using pointers
├── Slice/
│   └── main.go       # Dynamic slices, slicing arrays, append(), len(), and cap()
├── Struct/
│   └── main.go       # Custom types, nested structs, and constructor functions
├── Switch/
│   └── main.go       # Tagged vs tagless switch statements
├── for-loop/
│   └── main.go       # Loop patterns: standard, while-style, break, and continue
├── if-else/
│   └── main.go       # Grade calculator and boolean condition evaluation
├── main.go           # Root file demonstrating Anonymous functions and IIFEs
└── README.md         # Comprehensive documentation
```

---

## 🧠 Core Topics & Detailed Code Explanations

---

### 1. Anonymous Functions & IIFEs (`/`)
📁 **File:** [main.go](file:///home/mrrobinahmed/Projects/code/Level-2/GoFlow/main.go)

In Go, functions are **first-class citizens**. This means they can be assigned to variables, passed as arguments, returned from other functions, and executed dynamically.

```go
package main

import "fmt"

func main() {
	// 1. Anonymous Function assigned to a variable
	coffeeOrder := func() {
		fmt.Println("Coffee order placed")
	}
	coffeeOrder() // Execution

	// 2. Immediately Invoked Function Expression (IIFE)
	func(CoffeeType string) {
		fmt.Printf("Coffee order placed  %s.........", CoffeeType)
	}("Latte") // Invoked immediately with argument "Latte"

	// 3. Another Anonymous Function (Closure properties)
	mackCoffee := func() {
		Coffee := "Black Coffee"
		price := "150"
		fmt.Printf("Mack Coffee Price %s TK %s", Coffee, price)
	}
	mackCoffee()
}
```

#### 🔍 Detailed Explanation:
- **Anonymous Function (`coffeeOrder`):** A function defined without a name identifier. We bind it to `coffeeOrder`, which acts as a callable variable.
- **Immediately Invoked Function Expression (IIFE):** On line 17, the function is declared and appended with `("Latte")`. This executes the function block immediately without assigning it to any variable.
- **Local Scope:** Variables inside `mackCoffee` (like `Coffee` and `price`) are isolated to that scope and cannot be accessed outside it.

---

### 2. User Input & Scan (`Scan/`)
📁 **File:** [Scan/main.go](file:///home/mrrobinahmed/Projects/code/Level-2/GoFlow/Scan/main.go)

Go uses the `fmt` package to handle input operations. To write data back to a variable from input, Go requires passing the **memory address** of that variable.

```go
package main

import "fmt"

func main() {
	var choice int

	fmt.Print("Enter your number:\n")
	fmt.Scan(&choice) // '&' denotes passing by reference (Pointer)

	fmt.Println("Your choice is:", choice)
}
```

#### 🔍 Detailed Explanation:
- `var choice int`: Allocates space in memory for an integer variable, defaulting to `0`.
- `&choice`: The ampersand operator `&` extracts the memory address of `choice`.
- `fmt.Scan(&choice)`: Reads text from standard input (stdin), parses it into an integer, and writes it directly to the memory address of standard variable `choice`.

---

### 3. Conditional Logic (`if-else/`)
📁 **File:** [if-else/main.go](file:///home/mrrobinahmed/Projects/code/Level-2/GoFlow/if-else/main.go)

Conditional branching directs code execution paths based on evaluated boolean expressions (`true` or `false`).

```go
package main

import "fmt"

func main() {
	age := 20
	isAdult := age >= 18 // yields true
	fmt.Print(isAdult)
	
	if age >= 18 {
		fmt.Println("You are eligible for voting")
	}

	// Grading System Multi-branch Conditional Check
	score := 100

	if score >= 80 {
		fmt.Println("Grade A", score)
	} else if score >= 70 {
		fmt.Println("Grade B", score)
	} else if score >= 60 {
		fmt.Println("Grade C", score)
	} else if score >= 50 {
		fmt.Println("Grade D", score)
	} else {
		fmt.Println("Grade F sala fa fa", score)
	}
}
```

#### 🔍 Detailed Explanation:
- **No Parentheses:** In Go, conditions do not need parentheses `()` unless they dictate logical evaluation precedence.
- **Strict Code Blocks:** Curly braces `{}` are **mandatory** even if the block contains only one line of code.
- **Grading Evaluation:** The `if-else` chain checks conditions sequentially. Once a condition yields `true`, its block runs, and the rest of the chain is skipped.

---

### 4. Switch Statements (`Switch/`)
📁 **File:** [Switch/main.go](file:///home/mrrobinahmed/Projects/code/Level-2/GoFlow/Switch/main.go)

Switch statements provide a clean, readable alternative to nested `if-else` statements. Go switches differ from other languages because they have an **automatic break** mechanism.

```go
package main

import "fmt"

func main() {
	day := "sunday"

	// 1. Tagged Switch (Matching against 'day' value)
	switch day {
	case "sunday":
		fmt.Println("Today is sunday")
	case "friday":
		fmt.Println("Today is friday ")
	case "weekend":
		fmt.Println("Today is weekend")
	default:
		fmt.Println("Today is Not sunday")
	}

	// 2. Tagless Switch (Acts like clean if-else blocks)
	switch {
	case day == "sunday":
		fmt.Println("Today is sunday")
	case day == "friday":
		fmt.Println("Today is friday ")
	case day == "weekend":
		fmt.Println("Today is weekend")
	default:
		fmt.Println("Today is Not sunday")
	}
}
```

#### 🔍 Detailed Explanation:
- **Implicit Break:** Go automatically exits the switch block after executing the matched case. You do not need to write `break` at the end of each case.
- **Fallthrough:** If you want execution to move to the next case automatically, you must use the `fallthrough` keyword explicitly.
- **Tagless Switch:** `switch { ... }` acts exactly like `switch true { ... }`. It is highly useful when checking different variables or compound conditions in each case.

---

### 5. Loops & Iterations (`for-loop/`)
📁 **File:** [for-loop/main.go](file:///home/mrrobinahmed/Projects/code/Level-2/GoFlow/for-loop/main.go)

Go has **only one** looping keyword: `for`. However, it can represent three-component loops, while-style loops, and infinite loops.

```go
package main

import "fmt"

func makeCoffee(x int) {
	fmt.Println("making coffee............!!!", x)
}

func main() {
	// 1. Standard 3-component loop (initialization; condition; post-execution)
	for i := 0; i <= 10; i++ {
		makeCoffee(i)
	}

	// 2. While-style loop (condition only)
	i := 1
	for i <= 5 {
		makeCoffee(i)
		i++
	}

	// 3. Loop with Break (Stops loop early when condition met)
	for i := 0; i <= 10; i++ {
		makeCoffee(i)
		if i == 6 {
			break // Exits loop
		}
	}

	// 4. Loop with Continue (Skips current iteration for odd numbers)
	for i := 0; i <= 10; i++ {
		if i%2 != 0 {
			continue // Skip odd indexes, proceed to next iteration
		}
		makeCoffee(i)
	}
}
```

#### 🔍 Detailed Explanation:
- **While Loop emulation:** Done simply by omitting initialization and post-execution parts (`for condition {}`).
- **`break`:** Terminates the loop's execution completely.
- **`continue`:** Skip the remaining code inside the loop body for the current iteration and jumps directly to the next iteration step.

---

### 6. Arrays (`Array/`)
📁 **File:** [Array/main.go](file:///home/mrrobinahmed/Projects/code/Level-2/GoFlow/Array/main.go)

An array is a fixed-size sequence of elements of a single type. Its size is part of its type definition (e.g., `[5]int` and `[10]int` are different types).

```go
package main

import "fmt"

func main() {
	// Declaration of an array of size 5
	var numbers [5]int
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50

	fmt.Println(numbers)

	// Fetch length of the array
	fmt.Println("Length", len(numbers))

	// Direct Index Access
	fmt.Println("Index is: ", numbers[4], "Type is: ", numbers[3])

	// Iterating through an array
	for i := 0; i < len(numbers); i++ {
		fmt.Println("Index is: ", i, "Value is: ", numbers[i])
	}
}
```

#### 🔍 Detailed Explanation:
- **Value Type:** Arrays in Go are **value types**, not reference types. When you assign or pass an array to a function, Go copies the *entire contents* of the array.
- **Fixed Size:** The length of an array is determined at declaration and cannot grow or shrink.
- **Default/Zero Value:** If elements are not initialized, they default to the type's zero value (e.g., `0` for integers, `""` for strings).

---

### 7. Slices & Dynamic Arrays (`Slice/`)
📁 **File:** [Slice/main.go](file:///home/mrrobinahmed/Projects/code/Level-2/GoFlow/Slice/main.go)

A slice is a wrapper around a Go array, providing a dynamic, flexible, and windowed view of standard sequence arrays.

```go
package main

import "fmt"

func main() {
	// 1. Slice Literal
	var slice = []int{
		100, 200, 300, 400, 500, 600, 700, 800, 900, 1000,
	}

	// Mutating indexes
	slice[0] = 50
	slice[1] = 150
	slice[2] = 250
	slice[3] = 350
	slice[4] = 450
	slice[5] = 550
	slice[6] = 650
	slice[7] = 750
	slice[8] = 850
	slice[9] = 950

	// 2. Dynamic growth using append()
	slice = append(slice, 22)

	fmt.Println("The slice:", slice)
	fmt.Println("Length of the slice:", len(slice))     // Current elements
	fmt.Println("Capacity of the slice:", cap(slice))   // Total allocated underlying memory space
}
```

#### 🔍 Detailed Explanation:
- **Internal Structure:** A slice consists of three components:
  1. A pointer to the underlying array.
  2. The **length** (`len()`): number of elements present in the slice.
  3. The **capacity** (`cap()`): maximum capacity limits the slice can reach before resizing.
- **Dynamic Growth (`append`):** If a slice's capacity is exceeded, Go allocates a new, larger underlying array, copies the existing items, updates the slice reference pointer, and doubles the capacity.
- **Sharing Arrays:** Modifying slice elements modifies the backing array. Any other slice pointing to the same backing array will see the change.

---

### 8. Pointers & Memory Referencing (`Pointers/`)
📁 **File:** [Pointers/main.go](file:///home/mrrobinahmed/Projects/code/Level-2/GoFlow/Pointers/main.go)

Pointers store the memory address of another variable. They allow you to share and modify data without creating expensive memory copies.

```go
package main

import "fmt"

// Mutates value by reference
func change(x *int) {
	*x = 100 // Dereferencing and writing 100 to standard address
}

func main() {
	num := 10

	// Pass the address of 'num' using the '&' operator
	change(&num)

	fmt.Println(num) // Output is 100!
}
```

#### 🔍 Detailed Explanation:
- **`&` (Address-of Operator):** Finds the exact memory address location of a variable.
- **`*` (Dereferencing / Pointer Type):**
  - In a type declaration (like `x *int`), it means `x` is a pointer to an integer.
  - In an expression (like `*x = 100`), it reads or writes the value stored at the memory address `x` points to.
- **Pass by Reference simulation:** Go uses pass-by-value strictly. By passing a copy of the pointer address, we can manipulate the original variable value outside the function's stack frame.

---

### 9. Structs & Custom Types (`Struct/`)
📁 **File:** [Struct/main.go](file:///home/mrrobinahmed/Projects/code/Level-2/GoFlow/Struct/main.go)

Structs are user-defined composite types that group together zero or more named fields. They are used to model real-world concepts and data objects.

```go
package main

import "fmt"

// Define a struct type representation
type user struct {
	name string
	age  int
	role string
}

func main() {
	// A helper constructor-like function
	newUser := func(name string, age int, role string) user {
		return user{
			name: name,
			age:  age,
			role: role,
		}
	}

	// Initialize the struct using the helper constructor
	users := newUser("jon", 25, "admin")
	
	// Print struct details with field tags (+v)
	fmt.Printf("%+v", users)
}
```

#### 🔍 Detailed Explanation:
- **Declaration:** Struct fields are declared inside `type StructName struct` block.
- **Instantiating:** Instantiation can be done using a struct literal (e.g., `user{name: "jon"}`) or constructor functions.
- **Formatting `%+v`:** The `%+v` flag in `fmt.Printf` prints the struct along with its field names, which is very helpful for debugging.

---

## 📊 Quick Comparison: Array vs. Slice

| Feature | Array (`[5]int`) | Slice (`[]int`) |
| :--- | :--- | :--- |
| **Size** | Fixed size at compile-time | Dynamic (can grow at runtime) |
| **Declaration Syntax** | `var arr [5]int` | `var sl []int` |
| **Type Category** | Value type (copied on assignment) | Reference type (refers to backing array) |
| **Resizing** | Not allowed | Supported via `append(slice, val)` |
| **Capacity** | Always equal to length | Can be greater than length (`cap(s)`) |

---

## ⚡ How to Run the Code

To execute any program, navigate to the project directory and use standard `go run` commands.

```bash
# Run Anonymous Functions & IIFEs
go run main.go

# Run Arrays Module
go run Array/main.go

# Run Slices Module
go run Slice/main.go

# Run Pointers Module
go run Pointers/main.go

# Run Structs Module
go run Struct/main.go

# Run Switch Statements Module
go run Switch/main.go

# Run Loop Iterations Module
go run for-loop/main.go

# Run Conditionals Module
go run if-else/main.go

# Run Input Scan Module
go run Scan/main.go
```

---

## 💡 Key Go Rules & Gotchas to Remember

1. **Unused Imports are Errors:** If you import a library (like `"fmt"`) and don't use it, Go will refuse to compile.
2. **Unused Variables are Errors:** Any local variable declared must be used, or the compiler will throw an error.
3. **No Implicit Type Casting:** You cannot implicitly cast types (e.g., adding `int` and `int64` directly requires explicit type conversion like `int(myInt64) + myInt`).
4. **Cap Resizing Double Overhead:** When slice elements exceed capacity, Go allocates a new array of double the size. Be mindful of large loops appending to un-preallocated slices.
5. **Formatting Structs:** Use `%+v` to print structural fields clearly instead of just `%v` which only prints values.
