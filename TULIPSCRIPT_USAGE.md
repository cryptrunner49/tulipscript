# TulipScript Usage Guide

Welcome to the TulipScript Usage Guide! TulipScript is a JavaScript-inspired scripting language with modern constructs, Unicode/emoji support, and built-in features to enhance productivity. This guide walks through language features with practical examples, themed around colors.

---

## Table of Contents

1. [Variables](#1-variables)  
2. [Types of Variables](#2-types-of-variables)  
3. [Constants](#3-constants)  
4. [Control Flow](#4-control-flow)  
5. [Loops](#5-loops)  
6. [Blocks](#6-blocks)  
7. [Closures](#7-closures)  
8. [Fibonacci Recursive](#8-fibonacci-recursive)  
9. [Fibonacci Iterative](#9-fibonacci-iterative)  
10. [Structs](#10-structs)  
11. [Arrays](#11-arrays)  
12. [Maps](#12-maps)  
13. [File Operations](#13-file-operations)  
14. [Modules](#14-modules)  
15. [Import](#15-import)  
16. [Additional Features](#16-additional-features)  
    - [16.1. Input Handling](#161-input-handling)  
    - [16.2. String Formatting](#162-string-formatting)  
    - [16.3. Operators](#163-operators)  
    - [16.4. Shadowing](#164-shadowing)  
17. [Unicode Support](#17-unicode-support)
18. [Native Functions](#18-native-functions)  

---

## 1. Variables

```javascript
let hue = "Red"
println("Color:", hue)

let saturation = 75
println("Saturation:", saturation)

let isVivid = true
println("Is Vivid?:", isVivid)

let shade = null
println("Shade:", shade)
```

---

## 2. Types of Variables

```javascript
let colorName = "Blue"          // String
let brightness = 100            // Number
let isPrimary = true            // Boolean
let noColor = null              // Null
println("Color:", colorName, "Brightness:", brightness)
println("Primary?:", isPrimary, "No Color:", noColor)
```

---

## 3. Constants

```javascript
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

## 4. Control Flow

```javascript
let color = "Green"
if (color == "Red") {
    println("Warm color")
} | (color == "Blue") {
    println("Cool color")
} else {
    println("Other color")
}
```

---

## 5. Loops

```javascript
let count = 0
while (count < 3) {
    println("Color shade #", to_str(count))
    count = count + 1
}

for (let i = 0; i < 3; i++) {
    println("Tone:", i)
}
```

---

## 6. Blocks

```javascript
let tint = "Light Blue"
{
    let tint = "Dark Blue"
    println("Inner tint:", tint)
}
println("Outer tint:", tint)
```

---

## 7. Closures

```javascript
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

## 8. Fibonacci Recursive

```javascript
function fib(n) {
    if (n < 2) return n
    return fib(n - 2) + fib(n - 1)
}

let start = clock()
println("Fibonacci(16):", fib(16))
printf("Time taken: %v seconds\n", clock() - start)
```

---

## 9. Fibonacci Iterative

```javascript
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

## 10. Structs

```javascript
struct Color {
    name = "Unknown",
    rgb = [255, 255, 255]
}

function describeColor(color) {
    return color.name + ": RGB(" + array_to_string(color.rgb) + ")"
}

let favorite = Color()
favorite.name = "Teal"
favorite.rgb = [0, 128, 128]
println("Description:", describeColor(favorite))
```

---

## 11. Arrays

```javascript
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

## 12. Maps

```javascript
let colorMap = { "Red": "#FF0000", "Blue": "#0000FF" }
println("Red Hex:", colorMap["Red"])
colorMap["Green"] = "#00FF00"
println("Updated map:", colorMap)

println("Contains Blue:", map_contains_key(colorMap, "Blue"))
println("Keys:", map_keys(colorMap))
```

---

## 13. File Operations

```javascript
let filename = "colors.txt"
let content = "Primary Colors:\nRed\nBlue\nYellow"
write_file(filename, content)
println("Wrote to file:", filename)

let readContent = read_file(filename)
println("Read from file:", readContent)
```

---

## 14. Modules

```javascript
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

## 15. Import

```javascript
// main.tlp
import "color_utils.tlp"

println("Shade:", ColorUtils.Tones.getShade("Blue"))
println("Default Hue:", ColorUtils.DEFAULT_HUE)
```

---

## 16. Additional Features

### 16.1. Input Handling

```javascript
println("Enter a color:")
let input = scanln()
println("You chose:", input)
```

### 16.2. String Formatting

```javascript
let color = "Purple"
let intensity = 80
let formatted = sprintf("Color: %v, Intensity: %v", color, intensity)
println("Formatted:", formatted)
```

### 16.3. Operators

```javascript
let r = 255
let g = 128
println("Add:", r + g)
println("Exponent:", r ** 2)
println("Integer Div:", r /_ g)
println("50 percent of 1000:", 50 %% 1000)
```

### 16.4. Shadowing

```javascript
let hue = "Orange"
{
    let hue = "Cyan"
    println("Inner hue:", hue)
}
println("Outer hue:", hue)
```

---

## 17. Unicode Support

```javascript
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

```javascript
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
while !iter_done(iter)
    println("Iterator value:", iter_value(iter))     // Get current value
    iter_next(iter)                                 // Move to next
end

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
println("Parsed time:", to_str(parsed_time))
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
