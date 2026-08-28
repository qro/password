# 🔐 Password

A cross-platform desktop application for generating secure passwords and evaluating password strength, built in Go with Fyne.

---

## Features

- **Password Generator** — generate cryptographically secure passwords with configurable character sets (uppercase, lowercase, numbers, symbols) and length (8–128)
- **Strength Checker** — evaluate password entropy in bits with a visual strength rating
- **Breach Detection** — checks passwords against the [Have I Been Pwned](https://haveibeenpwned.com/) database using k-anonymity (your full password never leaves your machine)

---

## Installation

### Prerequisites

- [Go 1.21+](https://go.dev/dl/)
- A C compiler (required by Fyne for CGo):
  - **Windows:** [MSYS2](https://www.msys2.org/) with the MinGW GCC toolchain
    ```bash
    pacman -S mingw-w64-ucrt-x86_64-gcc
    ```
    Then add `C:\msys64\ucrt64\bin` to your system `PATH`.
  - **macOS:** Xcode Command Line Tools (`xcode-select --install`)
  - **Linux:** `gcc` via your package manager (e.g. `sudo apt install gcc`)

### Clone & Run

```bash
git clone https://github.com/qro/password.git
cd password
go run ./cmd/password
```

---

## Dependencies

| Module | Purpose |
|--------|---------|
| [fyne.io/fyne/v2](https://fyne.io/) | Cross-platform GUI framework |
| Go standard library `crypto/rand` | Cryptographically secure random number generation |
| Go standard library `crypto/sha1` | SHA-1 hashing for HIBP k-anonymity lookup |
| Go standard library `math` | Entropy calculation via `math.Log2` |
| Go standard library `net/http` | HTTP requests to the HIBP Pwned Passwords API |
| Go standard library `io` | Reading HTTP response bodies |
| Go standard library `unicode` | Detecting character classes in passwords |
| Go standard library `embed` | Embedding icon asset into the binary |

---

## External Services

| Service | Usage |
|---------|-------|
| [Have I Been Pwned — Pwned Passwords API](https://haveibeenpwned.com/API/v3#SearchingPwnedPasswordsByRange) | Breach detection using k-anonymity range queries. Only the first 5 characters of the password's SHA-1 hash are sent. |

---

## Information

- This was a half finished project that was lying around on my computer, so I decided to finish it up and give it a lovely UI
- Some other ideas I have in mind for this project:
    - Password manager tab
    - Breach lookup tab
- Please create an [issue](https://github.com/qro/password/issues/new) if you spot anything. If you need any support, join my [discord server](https://discord.gg/QsyXxynFEp) / send me an email at qro.gh@pm.me