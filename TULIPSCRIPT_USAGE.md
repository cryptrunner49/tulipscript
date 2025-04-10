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
    for (let i = 2; i <= n; i = i + 1) {
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

for (let i = 0; i < len(arr); i = i + 1) {
    println("Element", i, ":", arr[i])
}
```

---

## 7. File Operations

```tulipscript
let filename = "test.txt"
let content = "Hello from TulipScript!\nWritten on April 09, 2025."
write_file(filename, content)
println("Wrote to file:", filename)

let readContent = read_file(filename)
println("Read from file:", readContent)
```

---

## 8. Native Functions

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

## 9. Modules

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

## 10. Import

```tulipscript
// main.tlp
import "geometry.tlp"

println("Circle Area:", Geometry.Shapes.area_circle(5))
println("Circle Perimeter:", Geometry.Shapes.perimeter_circle(5))
```

---

## 11. Additional Features

### 11.1. Input Handling

```tulipscript
println("Enter a sentence:")
let input = scanln()
println("You entered:", input)
```

### 11.2. String Formatting

```tulipscript
let name = "Alice"
let age = 25
let formatted = sprintf("Name: %v, Age: %v", name, age)
println("Formatted:", formatted)
```

### 11.3. Advanced Control Flow

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
    if (i == 1) continue
    println("i:", i)
}
```

### 11.4. Operators

```tulipscript
let a = 10
let b = 3
println("Add:", a + b)
println("Exponent:", a ** 2)
println("Integer Div:", a /_ b)
println("25 percent of 1000:", 25 %% 1000)
```

---

## 12. Unicode Support

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
