# 🌷 TulipScript

**TulipScript** is a lightweight, expressive scripting language inspired by JavaScript — but more fun. It supports Unicode and emoji identifiers, functional programming, file I/O, and is embeddable as a shared C library. Whether you're scripting a game or experimenting with creative code, TulipScript makes it joyful.

---

## 🚀 Quick Start

```tulipscript
let hello = "Hello, World!"
println(hello)  // Outputs: Hello, World!
```

📖 Explore language features in the [Usage Guide →](TULIPSCRIPT_USAGE.md)

---

## ✨ Features

- 🌍 **Unicode & Emoji Identifiers** — Name your variables `🌸`, `π`, or anything you love.
- 🧠 **Familiar JavaScript-like Syntax** — Use `let`, `const`, `if`, `for`, and more.
- ⚙️ **Native Functions** — `clock()`, `shuffle()`, `random_between()` and others built-in.
- 📁 **File I/O** — `read_file()` and `write_file()` to handle files natively.
- 🧱 **Structs & Closures** — Define types, encapsulate behavior.
- 🐧 **Linux-Only** — Works out of the box on most Linux distributions.
- 🧬 **Embeddable VM** — Use it in C, Go, Rust, or C++ apps.

---

## 🔽 Download

Get the latest prebuilt binaries and development files from the [Releases Page →](https://github.com/cryptrunner49/tulipscript/releases/latest):

- **🌷 VM Executable**: [Download `tulip`](https://github.com/cryptrunner49/tulipscript/releases/latest/download/tulip)
- **🔧 Development Files**:
  - [libtulip.h](https://github.com/cryptrunner49/tulipscript/releases/latest/download/libtulip.h)
  - [libtulip.so](https://github.com/cryptrunner49/tulipscript/releases/latest/download/libtulip.so)
- **📦 Full Release Bundle** (VM + Libs + Headers + Examples):  
  [tulip-release.zip](https://github.com/cryptrunner49/tulipscript/releases/latest/download/tulip-release.zip)

---

## ⚙️ Installation

### ✅ Requirements

- OS: Linux
- Tools: `gcc`, `make`, `pkg-config`, `libffi`, `readline`
- [Go (Golang)](https://golang.org)

### 🧰 Install with Script

**System-wide:**

```bash
curl -sL https://github.com/cryptrunner49/tulipscript/raw/refs/heads/main/install.sh | bash -s -- install --system
```

**User-only:**

```bash
curl -sL https://github.com/cryptrunner49/tulipscript/raw/refs/heads/main/install.sh | bash -s -- install --user
```

### 🏗 Build from Source

```bash
git clone https://github.com/cryptrunner49/tulipscript.git
cd tulipscript
make vm
./bin/tulip samples/scripts/rpg_game.tlp
```

---

## 🧠 Embedding TulipScript

TulipScript is easy to embed in other languages like **C, Go, C++**, and **Rust**.

### ✅ Example in C

```c
#include "libtulip.h"
#include <stdio.h>
#include <stdlib.h>

int main(int argc, char** argv) {
    Tulip_Init(argc, argv);

    if (argc > 1) {
        Tulip_RunFile(argv[1]);
    } else {
        int exitCode;
        char* result = Tulip_InterpretWithResult("1 + 2;", "<repl>", &exitCode);
        if (exitCode == 0) printf("Last value: %s\n", result);
        else printf("Execution failed with code %d\n", exitCode);
        free(result);
    }

    Tulip_Free();
    return 0;
}
```

### 🛠 Build & Run

```bash
make lib
gcc -o run_sample samples/lib/c/sample.c -Lbin -ltulip -Ibin
LD_LIBRARY_PATH=bin ./run_sample
```

---

## 🔍 More Embedding Examples

Find ready-to-run embedding samples in:

- [📄 C](samples/lib/c/sample.c)
- [📄 C++](samples/lib/c/sample.cpp)
- [📄 Go](samples/lib/c/sample.go)
- [📄 Rust](samples/lib/c/sample.rust)

These show how to use TulipScript with FFI across different ecosystems.

---

## 🧪 OS Setup Instructions

### Ubuntu / Debian

```bash
sudo apt update
sudo apt install gcc pkg-config make golang libffi-dev libreadline-dev
```

---

## 🗺 Roadmap

What’s next for TulipScript?

- [x] **Elif Support**
- [ ] **Pattern Matching**
- [ ] **Switch Statement**
- [ ] **Enums**
- [ ] **Error Handling**
- [ ] **Standard Library**

✨ Track progress or suggest features via [Issues →](https://github.com/cryptrunner49/tulipscript/issues)

---

## 🤝 Contributing

We’d love your help! Bug fixes, documentation, or feature proposals — it all counts.

👉 See the [Contributing Guide →](CONTRIBUTING.md)

---

## 📄 License

**TulipScript** is licensed under the **GNU GPL-3.0**.  
See the [LICENSE](LICENSE) for details.
