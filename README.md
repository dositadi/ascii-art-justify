# ASCII-Art-Justify

A powerful command-line tool written in Go that generates stunning ASCII art banners with customizable alignment. Whether you need your text centered, right-aligned, or justified, this tool handles the heavy lifting of processing banner files and calculating terminal dimensions to provide a premium visual experience.

## ✨ Features

- **Custom Alignment**: Support for `left`, `right`, `center`, and `justify`.
- **Multiple Banner Styles**: Choose between `standard`, `shadow`, and `tinkertoy`.
- **Automatic Terminal Detection**: Dynamically calculates terminal width for perfect alignment.
- **Handling Newlines**: Support for `\n` in input strings to create multi-line ASCII art.

## 🚀 Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) (1.20+ recommended)

### Installation

Clone the repository to your local machine:

```bash
git clone https://github.com/your-username/ascii-art-justify.git
cd ascii-art-justify
```

## 🛠️ Usage

The basic syntax for running the program is:

```bash
go run . [OPTION] [STRING] [BANNER]
```

### Options

| Option | Description |
| :--- | :--- |
| `--align=left` | Aligns the ASCII art to the left (Default). |
| `--align=right` | Aligns the ASCII art to the right based on terminal width. |
| `--align=center` | Centers the ASCII art in the terminal. |
| `--align=justify` | Spreads words evenly across the terminal width. |

### Banners

- `standard`
- `shadow`
- `tinkertoy`

## 💡 Examples

**Right Alignment:**
```bash
go run . --align=right "Hello World" standard
```

**Center Alignment:**
```bash
go run . --align=center "ASCII" shadow
```

**Multi-line Input:**
```bash
go run . --align=left "Hello\nWorld" tinkertoy
```

## 📂 Project Structure

- `main.go`: Entry point for the application.
- `utils/`: Contains core logic for processing input and generating the ASCII output.
- `assets/`: Banner template files (e.g., `standard.txt`, `shadow.txt`).

## 🤝 Contributing

Contributions are welcome! If you have suggestions for new features or improvements, feel free to open an issue or submit a pull request.

---
*Created with ❤️ by John Falana*
