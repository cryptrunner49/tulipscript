# TulipScript Usage Guide

Welcome to the TulipScript Usage Guide! TulipScript is a JavaScript-like scripting language inspired by JavaScript, blending familiar syntax with unique features like Unicode and emoji identifiers. It supports modern programming constructs and offers a variety of built-in features to enhance developer productivity. This README provides practical examples and explanations of TulipScript's core features, making it an ideal starting point for developers. Each section includes a detailed explanation.

---

## Table of Contents

1. [Variables](#1-variables)
2. [Constants](#2-constants)
3. [Closures](#3-closures)
4. [Fibonacci Recursive](#4-fibonacci-recursive)
5. [Fibonacci Iterative](#5-fibonacci-iterative)
6. [Structs and Control Flow](#6-structs-and-control-flow)
7. [Arrays](#7-arrays)
8. [File Operations](#8-file-operations)
9. [Native Functions](#9-native-functions)
10. [Modules](#10-modules)
11. [Additional Features](#11-additional-features)

---

## 1. Variables

TulipScript supports mutable variables of various types, including numbers, strings, booleans, and null, with flexible naming using Unicode and emojis.

```tulipscript
// Number
let num = 42
println("Number:", num)  // Outputs: Number: 42

// String
let message = "Hello, TulipScript!"
println("Message:", message)  // Outputs: Message: Hello, TulipScript!

// Boolean
let isTrue = true
println("Boolean:", isTrue)  // Outputs: Boolean: true

// Null
let nothing = null
println("Null:", nothing)  // Outputs: Null: null

// Unicode and Emoji Names
let π = 3.14
println("Pi (π):", π)  // Outputs: Pi (π): 3.14

let 挨拶 = "こんにちは"
println("Greeting (挨拶):", 挨拶)  // Outputs: Greeting (挨拶): こんにちは

let 🔢 = 100
println("Number (🔢):", 🔢)  // Outputs: Number (🔢): 100
```

**Explanation**:  

- Use `let` to declare mutable variables with any value type: numbers, strings, booleans, or `null`.
- TulipScript supports Unicode (e.g., `π`, `挨拶`) and emoji (e.g., `🔢`) variable names, making it highly expressive.
- Semicolons are optional; statements are separated by line breaks or specific tokens.
- The `println` function outputs values with optional labels for clarity.

---

## 2. Constants

Constants in TulipScript are defined with `const` and cannot be modified after their initial definition, providing immutable values.

```tulipscript
// Constant Number
const maxScore = 100
println("Max Score:", maxScore)  // Outputs: Max Score: 100
// maxScore = 200  // Error: Cannot assign to constant variable 'maxScore'

// Constant String
const greeting = "Hello, World!"
println("Greeting:", greeting)  // Outputs: Greeting: Hello, World!

// Constant with Expression
const doublePi = 2 * 3.14
println("Double Pi:", doublePi)  // Outputs: Double Pi: 6.28
```

**Explanation**:  

- Use `const` to define immutable values that cannot be reassigned after initialization.
- Attempting to modify a constant (e.g., `maxScore = 200`) results in a runtime error.
- Constants can be initialized with expressions (e.g., `2 * 3.14`), computed at definition time.

---

## 3. Closures

Closures in TulipScript allow inner functions to access variables from outer scopes, enabling powerful functional programming patterns.

```tulipscript
function outer() {
    let a = 1
    let b = 2
    function middle() {
        let c = 3
        let d = 4
        function inner() {
            println("Sum:", a + c + b + d)  // Accesses variables from outer scopes
        }
        inner()
    }
    middle()
}
outer()  // Outputs: Sum: 10
```

**Explanation**:  

- Functions are defined with `function` (replacing `fn`) and can be nested.
- The `inner` function captures `a` and `b` from `outer` and `c` and `d` from `middle`, summing to `10` (1 + 3 + 2 + 4).
- Semicolons are omitted, relying on natural statement boundaries.

---

## 4. Fibonacci Recursive

This example computes the Fibonacci sequence recursively and measures execution time.

```tulipscript
function fib(n) {
    if (n < 2) return n
    return fib(n - 2) + fib(n - 1)
}

let start = clock()
println("Fibonacci(16):", fib(16))  // Outputs: Fibonacci(16): 987
printf("Time taken: %v seconds\n", clock() - start)  // Outputs: Time taken: <seconds>
```

**Explanation**:  

- `fib` uses recursion with a base case (`n < 2`).
- `clock()` returns the current time in seconds, used to measure performance.
- `printf` provides formatted output, showing the time difference.

---

## 5. Fibonacci Iterative

An iterative approach to Fibonacci calculation, optimized for performance.

```tulipscript
function fib(n) {
    if (n < 2) return n
    let a = 0
    let b = 1
    for (let i = 2; i <= n; i = i + 1) {
        let temp = a + b
        a = b
        b = temp
    }
    return b
}

let start = clock()
println("Fibonacci(16):", fib(16))  // Outputs: Fibonacci(16): 987
printf("Time taken: %v seconds\n", clock() - start)  // Outputs: Time taken: <seconds>
```

**Explanation**:  

- This version uses a `for` loop to iteratively compute Fibonacci numbers, updating `a` and `b` in each step.
- It’s more efficient than recursion for large `n`, as shown by the shorter execution time.
- The syntax `i = i + 1` increments the loop variable, though `i++` is also supported.

---

## 6. Structs and Control Flow

This script combines structs, functions, conditionals, and loops, with ASCII and Unicode examples.

```tulipscript
// ASCII Struct
struct Animal {
    species = "Unknown"
    length = 50  // Average cat length in cm
    height = 25  // Average cat height in cm
}

function describeAnimal(animal) {
    return animal.species + ": " + to_str(animal.length) + "x" + to_str(animal.height) + " cm"
}

let favorite = Animal()
favorite.species = "Cat"
if (favorite.length <= 50) {
    println("Animal is average or shorter")
} else {
    println("Animal is longer than average")
}
println("Description:", describeAnimal(favorite))  // Outputs: Description: Cat: 50x25 cm

let count = 0
while (count < 2) {
    println("Meow #", to_str(count))  // Outputs: Meow #0, Meow #1
    count = count + 1
}

// Unicode Struct
struct 猫 {
    種類 = "不明"
    年 = 2
}
let 子猫 = 猫()
子猫.種類 = "たまちゃん"
println("Cat name (子猫.種類):", 子猫.種類)  // Outputs: Cat name (子猫.種類): たまちゃん
```

**Explanation**:  

- Structs are defined with `struct`, allowing fields with default values (e.g., `species = "Unknown"`).
- The `if` statement checks conditions, and `while` loops iterate based on a condition.
- Unicode structs (e.g., `猫`) demonstrate multilingual support, with fields accessed via dot notation.

---

## 7. Arrays

TulipScript provides robust array support with built-in functions for manipulation.

```tulipscript
let arr = [1, 2, 3, 4, 5]
println("Original array:", array_to_string(arr))  // Outputs: Original array: [1, 2, 3, 4, 5]
push(arr, 6)
println("After push(6):", array_to_string(arr))  // Outputs: After push(6): [1, 2, 3, 4, 5, 6]
println("Popped:", pop(arr))  // Outputs: Popped: 6

for (let i = 0; i < len(arr); i = i + 1) {
    println("Element", i, ":", arr[i])  // Outputs: Element 0: 1, Element 1: 2, etc.
}
```

**Explanation**:  

- Arrays are created with square brackets (e.g., `[1, 2, 3]`).
- Functions like `push`, `pop`, `len`, and `array_to_string` manipulate and inspect arrays.
- The `for` loop iterates over indices, accessing elements with `arr[i]`.

---

## 8. File Operations

TulipScript supports basic file I/O operations for reading and writing text files.

```tulipscript
let filename = "test.txt"
let content = "Hello from TulipScript!\nWritten on April 09, 2025."
write_file(filename, content)
println("Wrote to file:", filename)

let readContent = read_file(filename)
println("Read from file:", readContent)  // Outputs: Read from file: Hello from TulipScript!...
```

**Explanation**:  

- `write_file(filename, content)` writes a string to a file, overwriting if it exists.
- `read_file(filename)` reads the entire file content as a string.
- Useful for simple file-based persistence or logging.

---

## 9. Native Functions

TulipScript includes native functions for tasks like timing, randomization, and I/O.

```tulipscript
let time = clock()
println("Current time:", time)  // Outputs: Current time: <timestamp>

let numbers = [1, 2, 3, 4, 5]
shuffle(numbers)
println("Shuffled array:", array_to_string(numbers))  // Outputs: Shuffled array: [e.g., 3, 1, 5, 2, 4]

let randNum = random_between(1, 10)
println("Random number (1-10):", randNum)  // Outputs: Random number (1-10): <random>
```

**Explanation**:  

- `clock()` returns the current time in seconds, useful for timing code.
- `shuffle(array)` randomizes array elements in place.
- `random_between(min, max)` generates a random integer between `min` and `max` (inclusive).

---

## 10. Modules

TulipScript supports modular programming with `import` and `mod` for code organization.

```tulipscript
// math.tlp
function add(a, b) {
    return a + b
}
const PI = 3.14159

// main.tlp
import "math.tlp"
println("Imported add(2, 3):", add(2, 3))  // Outputs: Imported add(2, 3): 5
println("Imported PI:", PI)  // Outputs: Imported PI: 3.14159
```

**Explanation**:  

- Use `import "filename.tlp"` to include external scripts.
- Functions and constants from the imported module (e.g., `add`, `PI`) are directly accessible.
- Modules help organize code into reusable units.

---

## 11. Additional Features

### 11.1. Input Handling

```tulipscript
println("Enter a sentence:")
let input = scanln()
println("You entered:", input)  // Outputs: You entered: <user input>
```

**Explanation**:  

- `scanln()` reads a line of user input from the console, returning it as a string.

### 11.2. String Formatting

```tulipscript
let name = "Alice"
let age = 25
let formatted = sprintf("Name: %v, Age: %v", name, age)
println("Formatted:", formatted)  // Outputs: Formatted: Name: Alice, Age: 25
```

**Explanation**:  

- `sprintf(format, ...args)` formats strings with placeholders (e.g., `%v` for values).

### 11.3. Control Flow Enhancements

```tulipscript
let x = 5
if (x > 0) {
    println("Positive")
} else if (x < 0) {
    println("Negative")
} else {
    println("Zero")
}

for (let i = 0; i < 3; i = i + 1) {
    if (i == 1) continue  // Skip 1
    println("i:", i)  // Outputs: i: 0, i: 2
}
```

**Explanation**:  

- Supports `if`, `else if`, `else` for conditionals, and `continue`/`break` in loops.
- Logical operators `&&` (and) and `||` (or) can be used in conditions (e.g., `if (x > 0 && x < 10)`).

### 11.4. Operators

```tulipscript
let a = 10
let b = 3
println("Add:", a + b)  // Outputs: Add: 13
println("Exponent:", a ** 2)  // Outputs: Exponent: 100
println("Integer Div:", a /_ b)  // Outputs: Integer Div: 3
println("25 percent of 1000:", 25 %% 1000)  // Outputs: 250
```

**Explanation**:  

- Includes arithmetic (`+`, `-`, `*`, `/`), exponentiation (`**`), integer division (`/_`), and percentage (`%%`).
- Logical operators `&&` and `||` replace `and` and `or` for boolean operations.
