# TulipScript Usage Guide

Welcome to the TulipScript Usage Guide! TulipScript is a JavaScript-inspired scripting language with modern constructs, Unicode/emoji support, and built-in features to enhance productivity. This guide walks through language features with practical examples, themed around colors.

---

## Table of Contents

1. [Variables](#1-variables)
2. [Constants](#2-constants)  
3. [Control Flow](#3-control-flow)
4. [Loops](#4-loops)
5. [Blocks](#5-blocks)
6. [Closures](#6-closures)
7. [Fibonacci Recursive](#7-fibonacci-recursive)
8. [Fibonacci Iterative](#8-fibonacci-iterative)
9. [Structs](#9-structs)
10. [Arrays](#10-arrays)
11. [Maps](#11-maps)
12. [File Operations](#12-file-operations)
13. [Modules](#13-modules)
14. [Import](#14-import)
15. [Additional Features](#15-additional-features)
    - [15.1. Input Handling](#151-input-handling)
    - [15.2. String Formatting](#152-string-formatting)
    - [15.3. Shadowing](#153-shadowing)
16. [Operators](#16-operators)
    - [16.1. Arithmetic Operators](#161-arithmetic-operators)
    - [16.2. Assignment Operators](#162-assignment-operators)
    - [16.3. Comparison Operators](#163-comparison-operators)
    - [16.4. Logical Operators](#164-logical-operators)
    - [16.5. Unary Operators](#165-unary-operators)
    - [16.6. Force Operator](#166-force-operator)
    - [16.7. Operator Precedence](#167-operator-precedence)
17. [Unicode Support](#17-unicode-support)
18. [Native Functions](#18-native-functions)

---

## 1. Variables

The `let` keyword is used to declare variables. A variable can store values like numbers, strings, booleans, or null. Variables are mutable, which means their values can be changed later. They are scoped to the block where they are declared and are also available in inner blocks, unless a variable with the same name is declared there. TulipScript figures out the variable type automatically based on the value you assign.

```tlp
// Number
let num = 42
println("Number:", num)

// String
let message = "Hello, Tulip!"
println("Message:", message)

// Boolean
let isTrue = true
println("Boolean:", isTrue)

// Null
let nothing = null
println("Nothing:", nothing)
```

---

## 2. Constants

The `const` keyword is used to declare constants. Constants are immutable, so their values cannot be changed once they are set. Like variables, constants can store values such as numbers, strings, booleans, or null and must be assigned a value when declared. Constants are block-scoped and are also accessible in inner blocks unless redefined there.

```tlp
const PRIMARY_COLOR = "Red"
println("Primary Color:", PRIMARY_COLOR)

const MAX_BRIGHTNESS = 255
println("Max Brightness:", MAX_BRIGHTNESS)

const IS_DEFAULT = false
println("Is Default?:", IS_DEFAULT)

const EMPTY_SHADE = null
println("Empty Shade:", EMPTY_SHADE)
```

---

## 3. Control Flow

The `if`, `else`, and `else if` (written as `|`) keywords allow for conditional logic. Conditions must be inside parentheses. You can write single-statement blocks directly after the condition without curly braces. For multiple statements, curly braces `{}` are required to group the block.

```tlp
let color = "Green"

if (color == "Green")
    println("Green")

if (color == "Red") {
    println("Warm color")
} | (color == "Blue") {
    println("Cool color")
} else {
    println("Other color")
}
```

---

## 4. Loops

The `while` keyword creates loops that run while a condition is true, using parentheses and curly braces. The `for` keyword iterates with an initializer, condition, and increment expression in parentheses, followed by a block. The `iter` keyword provides a concise way to iterate over arrays, using `let` to declare the loop variable and `in` to specify the array.

```tlp
// While Loop Iteration
println("--- While Loop ---")
let index = 0
while (index < len(arr)) {
    println("Element at index", index, ":", arr[index])
    index = index + 1
}

// For Loop Iteration
println("--- For Loop ---")
let arr = [10, 20, 30]
for (let i = 0; i < len(arr); i++) {
    println("Element at index", i, ":", arr[i])
}

// Iterator Loop
println("--- Iterator Loop ---")
iter (let item in [10, 20, 30]) {
    println("Item:", item)
}
```

---

## 5. Blocks

Curly braces `{}` define blocks, creating local scopes for variables declared with `let` or constants with `const`. Variables and Constants in inner blocks do not affect outer scopes, supporting isolated execution.

```tlp
let tint = "Light Blue"
{
    let tint = "Dark Blue"
    println("Inner tint:", tint)
}
println("Outer tint:", tint)
```

---

## 6. Closures

The `function` keyword defines functions, which can capture outer variables, forming closures. Closures retain access to their lexical scope, enabling stateful behavior across calls.

```tlp
function createColorMixer() {
    let baseColor = "White"
    function mix(color) {
        println("Mixing", color, "with", baseColor)
    }
    return mix
}
let mixer = createColorMixer()
mixer("Red")
```

---

## 7. Fibonacci Recursive

The `function` keyword, combined with `if` and `return`, supports recursive function definitions. Recursion involves a function calling itself with modified arguments until a base case is reached.

```tlp
function fib(n) {
    if (n < 2) return n
    return fib(n - 2) + fib(n - 1)
}

let start = clock()
println("Fibonacci(16):", fib(16))
printf("Time taken: %v seconds\n", clock() - start)
```

---

## 8. Fibonacci Iterative

The `for` loop, `let` declarations, and assignment (`=`) enable iterative algorithms. Variables are updated within the loop, avoiding recursive overhead for efficiency.

```tlp
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

## 9. Structs

The `struct` keyword defines custom data types with fields initialized using `=`. Fields are accessed with dot notation (`.`), and instances are created with the struct name and `{}`. The force operator `!{}` allows initializing structs with fields not defined in the struct, overriding defaults.

```tlp
struct Point {
    x = 10,
    y = 22
}

let p = Point{}  // Creates a Point instance with defaults x=10, y=22
println(p)

let p2 = Point{x = 1, y = 2}  // Creates a Point instance with x=1, y=2
println(p2)

struct Vec3
let v3 = Vec3!{x = 1, y = 2, z = 3}  // Forces creation with fields x, y, z
println(v3)
```

---

## 10. Arrays

Arrays are defined with square brackets `[]`, and elements are accessed via indices (e.g., `arr[i]`). Built-in functions like `push` and `pop` manipulate arrays dynamically.

```tlp
let colors = ["Red", "Blue", "Green"]
println("Colors:", array_to_string(colors))
push(colors, "Yellow")
println("After push:", array_to_string(colors))
println("Popped:", pop(colors))

for (let i = 0; i < len(colors); i++) {
    println("Color", i, ":", colors[i])
}
```

---

## 11. Maps

Maps are defined with curly braces `{}`, using key-value pairs (e.g., `"key": value`). Keys are accessed with square brackets `[]`, and built-in functions manage entries.

```tlp
let colorMap = { "Red": "#FF0000", "Blue": "#0000FF" }
println("Red Hex:", colorMap["Red"])
colorMap["Green"] = "#00FF00"
println("Updated map:", colorMap)

println("Contains Blue:", map_contains_key(colorMap, "Blue"))
println("Keys:", map_keys(colorMap))
```

---

## 12. File Operations

Built-in functions `write_file` and `read_file` handle file operations, taking filename and content as arguments. They enable persistent storage and retrieval of data.

```tlp
let filename = "colors.txt"
let content = "Primary Colors:\nRed\nBlue\nYellow"
write_file(filename, content)
println("Wrote to file:", filename)

let readContent = read_file(filename)
println("Read from file:", readContent)
```

---

## 13. Modules

The `mod` keyword defines modules, organizing code into namespaces. Nested modules and constants are supported, accessed via dot notation (e.g., `Module.Submodule`).

```tlp
// color_utils.tlp
mod ColorUtils {
    mod Tones {
        function getShade(color) {
            return color + " Shade"
        }
    }
    const DEFAULT_HUE = "Gray"
}
```

---

## 14. Import

The `import` keyword loads external modules by filename (e.g., `"file.tlp"`). Imported modules are accessed via their namespace, enabling modular code reuse.

```tlp
// main.tlp
import "color_utils.tlp"

println("Shade:", ColorUtils.Tones.getShade("Blue"))
println("Default Hue:", ColorUtils.DEFAULT_HUE)
```

---

## 15. Additional Features

### 15.1. Input Handling

The `scanln` function reads a line of user input, returning a string. It facilitates interactive programs by capturing input dynamically.

```tlp
println("Enter a color:")
let input = scanln()
println("You chose:", input)
```

### 15.2. String Formatting

The `sprintf` function formats strings using placeholders (e.g., `%v`), combining values into a single string for output or storage.

```tlp
let color = "Purple"
let intensity = 80
let formatted = sprintf("Color: %v, Intensity: %v", color, intensity)
println("Formatted:", formatted)
```

### 15.3. Shadowing

Shadowing in TulipScript allows a variable declared with `let` or `const` to override a previous variable with the same name, either in the same scope or in an inner scope. Both mutable (`let`) and immutable (`const`) variables can be shadowed. A `const` variable, while immutable (cannot be reassigned), can be shadowed by a new `let` or `const` declaration, creating a distinct variable that takes precedence. Shadowing in the same scope replaces the earlier declaration, with the last one being used. In different scopes, an inner scope variable shadows the outer scope variable without affecting it, preserving the outer variable's value outside the inner scope.

```tlp
// Shadowing in the same scope
let hue = "Red"
let hue = "Blue"  // Shadows previous 'hue' with a new mutable variable
println("Same scope hue:", hue)  // Outputs: Blue

const MAX_VALUE = 100
const MAX_VALUE = 200  // Shadows previous 'const' with a new immutable variable
println("Shadowed const:", MAX_VALUE)  // Outputs: 200

let shade = "Light"
const shade = "Dark"  // Shadows 'let' with 'const'
println("Shadowed let with const:", shade)  // Outputs: Dark

// Shadowing in different scopes
let color = "Yellow"
{
    let color = "Green"  // Shadows outer 'color' in inner scope
    println("Inner color:", color)  // Outputs: Green
    const color = "Purple"  // Shadows again in the same inner scope
    println("Inner const color:", color)  // Outputs: Purple
}
println("Outer color:", color)  // Outputs: Yellow, unaffected by inner scope
```

---

## 16. Operators

TulipScript supports a variety of operators for arithmetic, assignment, comparison, logical operations, and more. Below are the supported operators, organized by category, followed by their precedence rules.

### 16.1. Arithmetic Operators

- `+` (Addition)
- `-` (Subtraction)
- `*` (Multiplication)
- `/` (Division)
- `%` (Modulus)
- `**` (Exponentiation)
- `/_` (Integer division)
- `%%` (Percentage)

**Example**:

```tlp
let a = 10
let b = 3
println("Addition:", a + b)           // 13
println("Subtraction:", a - b)        // 7
println("Multiplication:", a * b)     // 30
println("Division:", a / b)           // 3.333...
println("Modulus:", a % b)            // 1
println("Exponentiation:", a ** 2)    // 100
println("Integer Division:", a /_ b)  // 3
println("Percentage:", 25 %% 1000)    // 250
```

### 16.2. Assignment Operators

- `=` (Assignment)

**Example**:

```tlp
let x = 5
x = x + 1
println("x:", x)  // 6
```

### 16.3. Comparison Operators

- `==` (Equal to)
- `!=` (Not equal to)
- `>` (Greater than)
- `<` (Less than)
- `>=` (Greater than or equal to)
- `<=` (Less than or equal to)

**Example**:

```tlp
let a = 5
let b = 3
println("Equal:", a == b)          // false
println("Not Equal:", a != b)      // true
println("Greater Than:", a > b)    // true
println("Less Than:", a < b)       // false
println("Greater or Equal:", a >= b) // true
println("Less or Equal:", a <= b)    // false
```

### 16.4. Logical Operators

- `&&` (Logical AND)
- `||` (Logical OR)
- `!` (Logical NOT)

**Example**:

```tlp
let t = true
let f = false
println("AND:", t && f)  // false
println("OR:", t || f)   // true
println("NOT:", !t)      // false
```

### 16.5. Unary Operators

- `++` (Increment)
- `--` (Decrement)
- `-` (Unary negation)

**Example**:

```tlp
let x = 5
println("Increment:", ++x)  // 6
println("Decrement:", --x)  // 5
println("Negation:", -x)    // -5
```

### 16.6. Force Operator

- `!{}` (Force struct instantiation with custom fields)

**Example**:

```tlp
struct Vec3
let v = Vec3!{x = 1, y = 2, z = 3}  // Forces creation with fields x, y, z
println(v)
```

### 16.7. Operator Precedence

Operator precedence determines the order in which operators are evaluated. The table below lists the precedence levels from highest to lowest.

| Precedence Level | Operators |
|------------------|-----------|
| Literals         | `number`, `string`, `boolean`, `null`, `( )` (grouped expressions) |
| Calls            | `.` (field access), `[]` (subscripting), `()` function calls |
| Unary            | `++`, `--`, `-` (negation), `!` (not) |
| Multiplicative   | `*`, `/`, `%`, `**`, `/_`, `%%` |
| Additive         | `+`, `-` |
| Comparison       | `>`, `<`, `>=`, `<=` |
| Equality         | `==`, `!=` |
| LogicalAnd       | `&&` |
| LogicalOr        | `\|\|` |
| Assignment       | `=` |

---

## 17. Unicode Support

TulipScript supports Unicode and emoji characters in identifiers and strings, using UTF-8 encoding. This allows multilingual and expressive variable names.

```tlp
let 🎨 = "Rainbow"
println("Palette 🎨:", 🎨)

let 色 = "青"
println("Color 色:", 色)

struct 色彩 {
    名前 = "不明"
}
let 好きな色 = 色彩()
好きな色.名前 = "緑"
println("Favorite color name:", 好きな色.名前)
```

---

## 18. Native Functions

Built-in functions, organized by category (e.g., String, Array), provide system-level operations. They are called directly, taking arguments to perform tasks like string manipulation or date handling.

```tlp
// === String Functions ===
let str = "  Hello, Tulip!  "
println("Original string:", str)
println("To string:", to_str(42))                    // Convert number to string
let chars = to_chars(str)                           // Convert string to array of characters
println("Characters:", array_to_string(chars))
println("Char at 2:", char_at(str, 2))              // Get character at index
println("Substring:", substring(str, 2, 7))         // Get substring
println("Index of 'Tulip':", str_index_of(str, "Tulip")) // First occurrence
println("Last index of 'l':", str_last_index_of(str, "l")) // Last occurrence
println("Contains 'Tulip':", str_contains(str, "Tulip")) // Check if substring exists
println("Starts with 'Hello':", starts_with(str, "Hello")) // Check prefix
println("Ends with '!':", ends_with(str, "!"))       // Check suffix
println("Uppercase:", to_upper(str))                 // Convert to uppercase
println("Lowercase:", to_lower(str))                 // Convert to lowercase
println("Trimmed:", trim(str))                      // Remove whitespace
let split = split(str, ",")                         // Split by delimiter
println("Split by ',':", array_to_string(split))
println("Replace 'Tulip' with 'World':", replace(str, "Tulip", "World")) // Replace substring
println("String length:", str_length(str))           // Get string length

// === Array Functions ===
let arr = [1, 2, 2, 3]
println("Original array:", array_to_string(arr))
println("Array length:", len(arr))                   // Get array length
push(arr, 4)                                        // Push element
println("After push(4):", array_to_string(arr))
let popped = pop(arr)                               // Pop element
println("Popped:", popped, "Array:", array_to_string(arr))
array_sort(arr)                                     // Sort array
println("Sorted:", array_to_string(arr))
let split_arr = array_split(arr, 2)                  // Split by separator
println("Split by 2:", array_to_string(split_arr))
let arr2 = [5, 6]
let joined = array_join(arr, arr2)                   // Join arrays
println("Joined:", array_to_string(joined))
array_sorted_push(arr, 3)                           // Insert into sorted array
println("Sorted push(3):", array_to_string(arr))
println("Linear search 2:", array_linear_search(arr, 2)) // Linear search
println("Binary search 3:", array_binary_search(arr, 3)) // Binary search
println("Index of 2:", index_of(arr, 2))            // First occurrence
println("Last index of 2:", last_index_of(arr, 2))  // Last occurrence
println("Contains 3:", array_contains(arr, 3))      // Check if element exists
array_reverse(arr)                                  // Reverse array
println("Reversed:", array_to_string(arr))
array_remove(arr, 2)                                // Remove first occurrence
println("After remove(2):", array_to_string(arr))
array_clear(arr)                                    // Clear array
println("Cleared:", array_to_string(arr))

// === Iterator Functions ===
let iter_arr = [10, 20, 30]
let iter = array_iter(iter_arr)                     // Create iterator
while !iter_done(iter) {
    println("Iterator value:", iter_value(iter))     // Get current value
    iter_next(iter)                                 // Move to next
}

// === Map Functions ===
let map = {"a": 1, "b": 2}
println("Map:", to_str(map))
map_remove(map, "a")                                // Remove key
println("After remove 'a':", to_str(map))
println("Contains key 'b':", map_contains_key(map, "b")) // Check key
println("Contains value 2:", map_contains_value(map, 2)) // Check value
println("Map size:", map_size(map))                 // Get map size
let keys = map_keys(map)                            // Get all keys
println("Keys:", array_to_string(keys))
let values = map_values(map)                        // Get all values
println("Values:", array_to_string(values))
map_clear(map)                                      // Clear map
println("Cleared map:", to_str(map))

// === Date Functions ===
let date = Date(2023, 10, 15)                       // Create date
println("Date:", to_str(date))
println("Current date:", to_str(date_now()))        // Current date
let parsed_date = date_parse_datetime("2024-01-01") // Parse date
println("Parsed date:", to_str(parsed_date))
println("Formatted date:", date_format_datetime(date, "2006-Jan-02")) // Format
let added_date = date_add_datetime(date, 1, 2, 3)   // Add time
println("Added date:", to_str(added_date))
let sub_date = date_subtract_datetime(date, 1, 2, 3) // Subtract time
println("Subtracted date:", to_str(sub_date))
println("Year:", date_get_component(date, "year"))  // Get component
let set_date = date_set_component(date, "year", 2025) // Set component
println("Set year:", to_str(set_date))
println("Add 5 days:", to_str(date_add_days(date, 5))) // Add days
println("Subtract 5 days:", to_str(date_subtract_days(date, 5))) // Subtract days

// === Time Functions ===
let time = Time(14, 30, 0)                          // Create time
println("Time:", to_str(time))
println("Current time:", to_str(time_now()))        // Current time
let parsed_time = time_parse("15:04:05")            // Parse time
println("Formatted time:", time_format(time, "15:04")) // Format
let added_time = time_add(time, 1, 30, 0)           // Add time
println("Added time:", to_str(added_time))
let sub_time = time_subtract(time, 1, 30, 0)        // Subtract time
println("Subtracted time:", to_str(sub_time))
println("Timezone:", time_get_timezone(time))       // Get timezone
let converted = time_convert_timezone(time, "America/New_York") // Convert timezone
println("Converted timezone:", to_str(converted))

// === DateTime Functions ===
let dt = DateTime(2023, 10, 15, 14, 30, 0)          // Create datetime
println("DateTime:", to_str(dt))
println("Current datetime:", to_str(datetime_now())) // Current datetime
let parsed_dt = datetime_parse("2023-10-15 14:30:00") // Parse datetime
println("Parsed datetime:", to_str(parsed_dt))
println("Formatted datetime:", datetime_format(dt, "2006-Jan-02 15:04")) // Format
let added_dt = datetime_add(dt, 1, 2, 3, 4, 5, 6)   // Add datetime
println("Added datetime:", to_str(added_dt))
let sub_dt = datetime_subtract(dt, 1, 2, 3, 4, 5, 6) // Subtract datetime
println("Subtracted datetime:", to_str(sub_dt))
println("Hour:", datetime_get_component(dt, "hour")) // Get component
let set_dt = datetime_set_component(dt, "hour", 16) // Set component
println("Set hour:", to_str(set_dt))
println("Add 5 days:", to_str(datetime_add_days(dt, 5))) // Add days
println("Subtract 5 days:", to_str(datetime_subtract_days(dt, 5))) // Subtract days

// === Random Functions ===
let colors = ["Red", "Blue", "Green"]
shuffle(colors)                                     // Shuffle array
println("Shuffled colors:", array_to_string(colors))
let rand_num = random_between(1, 10)                // Random number
println("Random number:", rand_num)
let rand_str = random_string(8)                     // Random string
println("Random string:", rand_str)

// === Output Functions ===
print("Hello, ")                                    // Print without newline
println("Tulip!")                                   // Print with newline
printf("Number: %d, String: %s\n", 42, "Tulip")     // Formatted print

// === Input Functions ===
let scanned = scan()                             // Read input as array
println("Scanned:", array_to_string(scanned))
let line = scanln()                             // Read line
println("Line:", line)
let formatted = scanf("%s")                     // Read formatted input
println("Formatted input:", formatted)

// === Formatting Functions ===
let formatted = sprintf("Value: %d", 42)            // Format string
println("Formatted:", formatted)
let err_msg = errorf("Error: %s", "Invalid")        // Format error
println("Error message:", err_msg)

// === File Operations ===
write_file("test.txt", "Hello, Tulip!")             // Write to file
let content = read_file("test.txt")                 // Read from file
println("File content:", content)

// === Utility Functions ===
let num = parse_int("123")                          // Parse string to int
println("Parsed int:", num)

// === Type Functions ===
println("Type of 42:", get_runtype(42))             // Get runtime type

// === Other Functions ===
let time = clock()                                  // Get current time in seconds
println("Current time (seconds):", time)
```
