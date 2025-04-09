# TulipScript

TulipScript is a lightweight, expressive scripting language designed for simplicity and flexibility. With support for Unicode and emoji identifiers, native functions, and modern programming constructs, TulipScript is perfect for quick scripts, educational purposes, or experimenting with creative coding. Whether you're shuffling arrays, writing to files, or defining structs with cat emojis (🐱), TulipScript makes programming fun and accessible.

---

## Features

- **Unicode & Emoji Support**: Use `π`, `挨拶`, or `🔢` as variable names.
- **Simple Syntax**: Easy-to-learn constructs like `let`, `const`, `function`, `if`, and `for`.
- **Native Functions**: Built-ins like `clock()`, `shuffle()`, and `random_between()`.
- **Structs & Closures**: Define custom types and leverage functional programming.
- **File I/O**: Read and write files with `read_file()` and `write_file()`.
- **Cross-Platform**: Runs on any system with the required dependencies.

Explore more in the [TulipScript Usage Guide](TULIPSCRIPT_USAGE.md)!

---

## Installation

### Requirements

To build and run the TulipScript VM, you’ll need:

```text
- Go (golang)
- libffi
- readline
- gcc
- pkg-config
```

### Install on Ubuntu

```bash
sudo apt update
sudo apt install golang libffi-dev libreadline-dev gcc pkg-config
```

### Install on macOS

```bash
brew install go libffi readline pkg-config
```

### Building TulipScript

1. Clone the repository:

   ```bash
   git clone https://github.com/cryptrunner49/tulipscript.git
   cd tulipscript
   ```

2. Build the interpreter:

   ```bash
   make build
   ```

3. Run a script:

   ```bash
   ./bin/seed sample/rpg_game.tlp
   ```

---

## Usage

Try this simple example:

```tulipscript
var hello = "Hello, TulipScript!";
println(hello)  // Outputs: Hello, TulipScript!

for (var i = 0; i < 3; i = i + 1) {
    println("Count:", i)  // Outputs: Count: 0, Count: 1, Count: 2
}
```

For detailed examples of variables, structs, loops, and more, check out the [TulipScript Usage Guide](TULIPSCRIPT_USAGE.md).

---

## Roadmap

Here’s what’s planned for TulipScript’s future:

- [ ] **Pattern Matching**: Add expressive pattern matching for conditionals.
- [ ] **Switch Case**: Implement a `switch` statement for multi-branch logic.
- [ ] **Elif**: Extend `if` with `elif` for cleaner conditional chains.
- [ ] **Enums**: Introduce enumerated types for structured data.
- [ ] **Error Handling**: Add try-catch or similar mechanisms.
- [ ] **Standard Library**: Expand with more utility functions.

See the [Issues](https://github.com/cryptrunner49/tulipscript/issues) tab for progress and to suggest features!

---

## Contributing

We’d love your help to make TulipScript better! Whether it’s adding examples, fixing bugs, or suggesting features, your contributions are welcome. Read our [Contributing Guide](CONTRIBUTING.md) for details on how to get started.

---

## License

TulipScript is licensed under the MIT License. See the [LICENSE](LICENSE) file for full details.
