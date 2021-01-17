# Baseconv

Baseconv - is a number-base converter, which can convert numbers from one base to another.  
Supported:

* Base-2 (binary)
* Base-4 (quaternary)
* Base-8 (octal)
* Base-10 (decimal)
* Base-16 (hexadecimal)
* Base-36 (hexatrigesimal)
* Base-62 (duosexagesimal)

Also is able to convert numbers with custom alphabets.

## Installation

1. The first need [Go](https://golang.org/) installed, then you can use this command to install Baseconv.  
   `go get -u github.com/marksartdev/baseconv`

2. Import to your project:  
   `import github.com/marksartdev/baseconv`

## Documentation

| Function | Description |
|---|---|
| `conv.Convert` | Convert number from one base to another |
| `conv.ConvertFromCustomBase` | Convert number from custom base using custom alphabet |
| `conv.ConvertToCustomBase` | Convert number to custom base using custom alphabet |
| `search.IndexOfRune` | Search a rune in slice and returns index of it. Function uses a binary search algorithm. **Attention!** Slice must be **sorted** and contains only **unique** runes. |
