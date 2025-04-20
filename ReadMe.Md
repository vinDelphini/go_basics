# 📘 About Go

Go (also known as Golang) is:

1. **Statically typed**
2. **Strongly typed**
3. **Compiled**
4. **Fast to compile**
5. Built-in **concurrency** support via **Goroutines**
6. Emphasizes **simplicity**, with features like **automatic garbage collection**

---

## 🛠️ Core Concepts

1. ✅ **Variables** and basic **data types**
2. ✅ **Functions** and **control structures** (`if`, `switch`, `for`)
3. ✅ **Arrays**, **Slices**, **Maps**, and **Loops**
4. ✅ **Strings**, **Runes**, **Bytes**
5. ✅ **Structs** and **Interfaces**

---

## 🚀 Main Concepts

1. ⭐ **Pointers**
2. ⭐ **Goroutines**
3. ⭐ **Channels**
4. ⭐ **Generics** (introduced in Go 1.18)

---

## 📚 `strings` Package Cheatsheet

```go
import "strings"
```

| Function                                  | Description                             | Example Output         |
|-------------------------------------------|-----------------------------------------|-------------------------|
| `strings.ToUpper("go")`                  | Converts to uppercase                   | `"GO"`                 |
| `strings.ToLower("GO")`                  | Converts to lowercase                   | `"go"`                 |
| `strings.Contains("hello", "he")`        | Checks for substring                    | `true`                 |
| `strings.HasPrefix("hello", "he")`       | Checks if string starts with prefix     | `true`                 |
| `strings.HasSuffix("hello", "lo")`       | Checks if string ends with suffix       | `true`                 |
| `strings.Split("a,b,c", ",")`            | Splits string into slice                | `[]string{"a","b","c"}`|
| `strings.Join([]string{"a", "b"}, "-")`  | Joins slice into string                 | `"a-b"`                |
| `strings.ReplaceAll("go go go", "go", "run")` | Replaces all instances of a substring | `"run run run"`        |

---
