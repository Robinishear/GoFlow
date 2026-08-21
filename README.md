# 🚀 GoFlow — Go Control Flow & Fundamentals Masterclass

[![Go Version](https://img.shields.io/badge/Go-1.20+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![Learning Series](https://img.shields.io/badge/Level--2-Go%20Fundamentals-orange?style=for-the-badge)]()
[![Status](https://img.shields.io/badge/Status-Completed-success?style=for-the-badge)]()

Welcome to **GoFlow**! A hands-on repository dedicated to mastering control flow, user input handling, loops, conditionals, and functional programming concepts in **Golang**.

---

## 📑 Table of Contents

- [📁 Project Structure](#-project-structure)
- [🧠 Topics & Code Explanations](#-topics--code-explanations)
  - [1. User Input (`Scan/`)](#1-user-input-scan)
  - [2. Conditional Logic (`if-else/`)](#2-conditional-logic-if-else)
  - [3. Switch Statements (`Switch/`)](#3-switch-statements-switch)
  - [4. Loops & Iterations (`for-loop/`)](#4-loops--iterations-for-loop)
  - [5. Anonymous Functions & IIFE (`main.go`)](#5-anonymous-functions--iife-maingo)
- [⚡ How to Run](#-how-to-run)
- [💡 Key Go Language Rules Learned](#-key-go-language-rules-learned)

---

## 📁 Project Structure

```bash
GoFlow/
├── Scan/
│   └── main.go       # Terminal user input with fmt.Scan & pointers
├── if-else/
│   └── main.go       # Boolean checks & multi-branch grading system
├── Switch/
│   └── main.go       # Tagged vs tagless (expressionless) switch cases
├── for-loop/
│   └── main.go       # Standard loops, while-style loops, break & continue
├── main.go           # Anonymous functions, closures & IIFE
└── README.md         # Documentation & learning notes
```

---

## 🧠 Topics & Code Explanations

### 1. User Input (`Scan/`)
📁 **File:** [`Scan/main.go`](file:///home/mrrobinahmed/Projects/code/Level-2/GoFlow/Scan/main.go)

Go provides the `fmt.Scan`, `fmt.Scanf`, and `fmt.Scanln` functions to read input from standard input (`os.Stdin`).

```go
package main

import "fmt"

func main() {
    var choice int

    fmt.Print("Enter your number:\n")
    fmt.Scan(&choice) // '&' passes the memory address (pointer) of variable choice

    fmt.Println("Your choice is:", choice)
}
```

#### 🔍 Explanation:
- `var choice int`: Declares an integer variable.
- `fmt.Scan(&choice)`: We pass `&choice` (the memory address of `choice`) so that `fmt.Scan` can store the user's input directly into the variable.

---

### 2. Conditional Logic (`if-else/`)
📁 **File:** [`if-else/main.go`](file:///home/mrrobinahmed/Projects/code/Level-2/GoFlow/if-else/main.go)

Conditional branching allows code execution based on boolean conditions.

```go
package main

import "fmt"

func main() {
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
        fmt.Println("Grade F", score)
    }
}
```

#### 🔍 Explanation:
- Evaluates conditions from top to bottom.
- As soon as one condition evaluates to `true` (e.g. `score >= 80`), its block executes and the rest of the chain is skipped.
- If none match, the `else` block executes.

---

### 3. Switch Statements (`Switch/`)
📁 **File:** [`Switch/main.go`](file:///home/mrrobinahmed/Projects/code/Level-2/GoFlow/Switch/main.go)

Go switch statements are flexible and don't require manual `break` statements.

#### Option A: Tagged Switch
```go
switch day {
case "sunday":
    fmt.Println("Today is sunday")
case "friday":
    fmt.Println("Today is friday")
default:
    fmt.Println("Today is Not sunday")
}
```

#### Option B: Tagless / Expressionless Switch (Clean alternative to `if-else if`)
```go
switch {
case day == "sunday":
    fmt.Println("Today is sunday")
case day == "friday":
    fmt.Println("Today is friday")
case day == "weekend":
    fmt.Println("Today is weekend")
default:
    fmt.Println("Today is Not sunday")
}
```

#### 🔍 Explanation:
- **Automatic Break:** Go automatically breaks after executing a matching case (no fallthrough unless explicitly specified with `fallthrough`).
- **Tagless Switch:** Writing `switch { ... }` is equivalent to `switch true { ... }`, making it clean for complex conditional comparisons.

---

### 4. Loops & Iterations (`for-loop/`)
📁 **File:** [`for-loop/main.go`](file:///home/mrrobinahmed/Projects/code/Level-2/GoFlow/for-loop/main.go)

> [!NOTE]
> Go has **only one** looping construct: the `for` loop! It handles standard loops, while-style loops, and infinite loops.

```go
package main

import "fmt"

func makeCoffee(x int) {
    fmt.Println("making coffee............!!!", x)
}

func main() {
    // 1. Standard 3-component loop
    // for i := 0; i <= 10; i++ {
    //     makeCoffee(i)
    // }

    // 2. While-style loop
    // i := 1
    // for i <= 5 {
    //     makeCoffee(i)
    //     i++
    // }

    // 3. Continue & Skip odd numbers (Prints only even numbers)
    for i := 0; i <= 10; i++ {
        if i%2 != 0 {
            continue // Skip rest of current iteration
        }
        makeCoffee(i)
    }
}
```

#### 🔍 Key Loop Controls:
- `break`: Terminate the loop immediately.
- `continue`: Skip the current iteration and move to the next step.

---

### 5. Anonymous Functions & IIFE (`main.go`)
📁 **File:** [`main.go`](file:///home/mrrobinahmed/Projects/code/Level-2/GoFlow/main.go)

Functions in Go are first-class citizens. You can assign functions to variables or execute them immediately.

```go
package main

import "fmt"

func main() {
    // 1. Anonymous Function assigned to a variable
    coffeeOrder := func() {
        fmt.Println("Coffee order placed")
    }
    coffeeOrder()

    // 2. Immediately Invoked Function Expression (IIFE)
    func(CoffeeType string) {
        fmt.Printf("Coffee order placed %s.........\n", CoffeeType)
    }("Latte")

    // 3. Anonymous Function with multiple variables (Closure)
    mackCoffee := func() {
        coffee := "Black Coffee"
        price := "150"
        fmt.Printf("Make Coffee Price %s TK %s\n", coffee, price)
    }
    mackCoffee()
}
```

#### 🔍 Explanation:
- **Anonymous Functions:** Functions without names, defined inline.
- **IIFE:** Defined and invoked immediately using `func(...) { ... }(arguments)`. Useful for scoped execution and goroutines.

---

## ⚡ How to Run

To run any of the examples directly in your terminal:

```bash
# Run User Input module
go run ./Scan/main.go

# Run If-Else module
go run ./if-else/main.go

# Run Switch module
go run ./Switch/main.go

# Run For-Loop module
go run ./for-loop/main.go

# Run Root main.go (Anonymous Functions / IIFE)
go run ./main.go
```

---

## 💡 Key Go Language Rules Learned

1. **No parentheses around conditions:** Go conditions (`if score >= 80`, `for i <= 5`) do not require parentheses `()`.
2. **Curly braces `{}` are mandatory:** Even for single-line blocks, `{}` must be used.
3. **Pointers for Input:** `fmt.Scan` requires pointer references (`&varName`) to mutate the variable's value.
4. **Single Loop Keyword:** The `for` keyword does everything (standard loop, while loop, infinite loop, range loop).
5. **No implicit fallthrough in switch:** Cases break automatically unless `fallthrough` is explicitly written.
6. **First-class Functions:** Functions can be stored in variables, passed around, or executed immediately (IIFE).
