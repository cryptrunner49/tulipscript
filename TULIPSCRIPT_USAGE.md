# TulipScript Usage Guide

Welcome to the TulipScript Usage Guide! TulipScript is a JavaScript-inspired scripting language with modern constructs, Unicode/emoji support, and built-in features to enhance productivity. This guide walks through language features with practical examples.

---

## Table of Contents

1. [Variables](#1-variables)  
2. [Closures](#2-closures)  
3. [Fibonacci Recursive](#3-fibonacci-recursive)  
4. [Fibonacci Iterative](#4-fibonacci-iterative)  
5. [Structs and Control Flow](#5-structs-and-control-flow)  
6. [Arrays](#6-arrays)  
7. [File Operations](#7-file-operations)  
8. [Native Functions](#8-native-functions)  
9. [Modules](#9-modules)  
10. [Import](#10-import)  
11. [Additional Features](#11-additional-features)  
12. [Unicode Support](#12-unicode-support)

---

## 1. Variables

```tulipscript
let num = 42
println("Number:", num)

let message = "Hello, TulipScript!"
println("Message:", message)

let isTrue = true
println("Boolean:", isTrue)

let nothing = null
println("Null:", nothing)
```

---

## 2. Closures

```tulipscript
function outer() {
    let a = 1
    let b = 2
    function middle() {
        let c = 3
        let d = 4
        function inner() {
            println("Sum:", a + c + b + d)
        }
        inner()
    }
    middle()
}
outer()
```

---

## 3. Fibonacci Recursive

```tulipscript
function fib(n) {
    if (n < 2) return n
    return fib(n - 2) + fib(n - 1)
}

let start = clock()
println("Fibonacci(16):", fib(16))
printf("Time taken: %v seconds\n", clock() - start)
```

---

## 4. Fibonacci Iterative

```tulipscript
function fib(n) {
    if (n < 2) return n
    let a = 0
    let b = 1
    for (let i = 2; i <= n; i++) {
        let temp = a + b
        a = b
        b = temp
    }
    return b
}

let start = clock()
println("Fibonacci(16):", fib(16))
printf("Time taken: %v seconds\n", clock() - start)
```

---

## 5. Structs and Control Flow

```tulipscript
struct Animal {
    species = "Unknown",
    length = 50,
    height = 25
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
println("Description:", describeAnimal(favorite))

let count = 0
while (count < 2) {
    println("Meow #", to_str(count))
    count = count + 1
}
```

---

## 6. Arrays

```tulipscript
let arr = [1, 2, 3, 4, 5]
println("Original array:", array_to_string(arr))
push(arr, 6)
println("After push(6):", array_to_string(arr))
println("Popped:", pop(arr))

for (let i = 0; i < len(arr); i++) {
    println("Element", i, ":", arr[i])
}
```

---

## 7. Maps

```tulipscript
// Basic Map Operations
println("--- Map Demo ---")
let map = { "name": "Alice", "age": 30 }
println("Name:", map["name"])  // Outputs: Alice
println("Age:", map["age"])    // Outputs: 30
map["age"] = 31
println("Updated age:", map["age"])  // Outputs: 31

// Map Functions
println("--- Map Functions ---")
let m = {"a": 1, "b": 2}
println("Initial map:", m)
map_remove(m, "a")
println("After removing 'a':", m)
println("Contains key 'b':", map_contains_key(m, "b"))
println("Contains value 2:", map_contains_value(m, 2))
println("Map size:", map_size(m))
println("Keys:", map_keys(m))
println("Values:", map_values(m))
map_clear(m)
println("After clear:", m)

// Map Addition
println("--- Map Addition ---")
let a = {"x": 1, "y": 2}
let b = {"y": 3, "z": 4}
println("a + b:", a + b)

// Map Subtraction
println("--- Map Subtraction ---")
let a = {"x": 1, "y": 2, "z": 3}
let b = {"y": null, "w": null}
println("a - b:", a - b)
```

---

## 8. File Operations

```tulipscript
let filename = "test.txt"
let content = "This is a file handling example.\nDemonstrating how to read and write files."
write_file(filename, content)
println("Wrote to file:", filename)

let readContent = read_file(filename)
println("Read from file:", readContent)
```

---

## 9. Native Functions

```tulipscript
let time = clock()
println("Current time:", time)

let numbers = [1, 2, 3, 4, 5]
shuffle(numbers)
println("Shuffled array:", array_to_string(numbers))

let randNum = random_between(1, 10)
println("Random number (1-10):", randNum)
```

---

## 10. Modules

```tulipscript
// geometry.tlp

mod Geometry {
    mod Shapes {
        function area_circle(radius) {
            return Geometry.PI * radius * radius
        }

        function perimeter_circle(radius) {
            return 2 * Geometry.PI * radius
        }
    }

    const PI = 3.14159
}
```

---

## 11. Import

```tulipscript
// main.tlp
import "geometry.tlp"

println("Circle Area:", Geometry.Shapes.area_circle(5))
println("Circle Perimeter:", Geometry.Shapes.perimeter_circle(5))
```

---

## 12. Additional Features

### 12.1. Input Handling

```tulipscript
println("Enter a sentence:")
let input = scanln()
println("You entered:", input)
```

### 12.2. String Formatting

```tulipscript
let name = "Alice"
let age = 25
let formatted = sprintf("Name: %v, Age: %v", name, age)
println("Formatted:", formatted)
```

### 12.3. Advanced Control Flow

```tulipscript
let x = 5
if (x > 0) {
    println("Positive")
} else if (x < 0) {
    println("Negative")
} else {
    println("Zero")
}

for (let i = 0; i < 3; i++) {
    if (i == 1) continue // Skip the iteration when i is 1
    if (i == 2) break    // Exit the loop when i is 2
    println("i:", i)
}
```

### 12.4. Operators

```tulipscript
let a = 10
let b = 3
println("Add:", a + b)
println("Exponent:", a ** 2)
println("Integer Div:", a /_ b)
println("25 percent of 1000:", 25 %% 1000)
```

---

## 13. Unicode Support

```tulipscript
let π = 3.14159
println("Value of π:", π)

let 🌷 = "Tulip emoji"
println("Emoji variable 🌷:", 🌷)

let greeting = "こんにちは世界"
println("Greeting in Japanese:", greeting)

struct 数学 {
    半径 = 5
}
let 円 = 数学()
println("半径 (radius):", 円.半径)
```
