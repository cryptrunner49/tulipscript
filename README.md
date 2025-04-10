# 🌷 TulipScript

**TulipScript** is a lightweight, expressive, JavaScript-inspired scripting language designed for simplicity and flexibility. Drawing from JavaScript’s familiar syntax, it adds unique features like Unicode and emoji identifiers, native functions, and modern programming constructs. TulipScript is perfect for quick scripts or experimenting with creative coding. Whether you're shuffling arrays, writing to files, or defining structs with cat emojis (🐱), TulipScript makes programming fun and accessible.

---

## 👩‍💻 Hello World

```tulipscript
let hello = "Hello, World!"
println(hello)    // Outputs: Hello, World!
```

📖 Explore variables, structs, loops, and more in the [TulipScript Usage Guide](TULIPSCRIPT_USAGE.md).

---

## ✨ Features

- **🌍 Unicode & Emoji Identifiers** — Name variables like `🌸` or `π`.
- **🧠 JavaScript-Inspired Syntax** — Use familiar constructs like `let`, `const`, `function`, `if`, and `for`.
- **⚙️ Native Functions** — Built-ins like `clock()`, `shuffle()`, and `random_between()`.
- **🧱 Structs & Closures** — Define custom types and use functional constructs.
- **📁 File I/O** — Read and write files with `read_file()` and `write_file()`.
- **🖥 Cross-Platform** — Works on Linux and macOS with standard tooling.

📚 Learn more in the [TulipScript Usage Guide →](TULIPSCRIPT_USAGE.md)

---

## 📦 Installation

### ✅ Requirements

- Linux (Debian, Ubuntu, Fedora, Arch, etc.) or macOS
- [Go (Golang)](https://golang.org)
- Dependencies: `gcc`, `pkg-config`, `make`, `libffi`, `readline`

### 🧰 System Setup

#### Option 1: Install via Script

**System-wide installation (requires `sudo`)**:

```bash
curl -sL https://github.com/cryptrunner49/tulipscript/raw/refs/heads/main/install.sh | bash -s -- install --system
```

**User-only installation (`$HOME/.local/bin`)**:

```bash
curl -sL https://github.com/cryptrunner49/tulipscript/raw/refs/heads/main/install.sh | bash -s -- install --user
```

#### Option 2: Manual Download

```bash
curl -LO https://github.com/cryptrunner49/tulipscript/releases/latest/download/tulip
chmod +x tulip
```

---

### 🛠 Build From Source

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
   ./bin/tulip sample/rpg.tlp
   ```

---

## 🧪 Platform-Specific Setup

### Ubuntu/Debian

```bash
sudo apt update
sudo apt install gcc pkg-config make golang libffi-dev libreadline-dev
```

### macOS

```bash
brew install go pkg-config gcc make libffi readline
```

---

## 🗺 Roadmap

Coming soon to TulipScript:

- [ ] **Pattern Matching** — More expressive conditionals.
- [ ] **Switch Statement** — Cleaner multi-branch logic.
- [ ] **Elif Support** — Less nesting, more clarity.
- [ ] **Enums** — Organize data like a pro.
- [ ] **Error Handling** — Try-catch or similar constructs.
- [ ] **Standard Library** — More built-in power.

🎯 Track progress or suggest features via [Issues →](https://github.com/cryptrunner49/tulipscript/issues)

---

## 🤝 Contributing

We’d love your help! Whether it's fixing bugs, improving docs, or proposing features—your contributions matter.

📘 See the [Contributing Guide →](CONTRIBUTING.md) to get started.

---

## 📄 License

TulipScript is licensed under the **GNU General Public License v3.0 (GPL-3.0)**.  
See the [LICENSE](LICENSE) file for full details.
