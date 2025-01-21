# Go Basics Summary


---

## **1. Hello, World!** 
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```
- `package main`: Defines an executable program.
- `import "fmt"`: Imports the `fmt` package for printing.
- `main()`: The entry point of the program.

---

## **2. Packages, Imports, and Exported Names** 
- Go programs are organized into **packages**.
- The **main** package is required for executables.
- Import standard libraries using `import "package_name"`.
- Exported names (public symbols) start with a **capital letter**.
```go
fmt.Println(math.Pi) // Pi is exported from the "math" package
```

---

## **3. Functions** 
```go
func add(x int, y int) int {
    return x + y
}
```
- If multiple parameters share a type, they can be grouped:
```go
func subtract(x, y int) int {
    return x - y
}
```

---

## **4. Multiple Return Values** 
```go
func swap(a, b string) (string, string) {
    return b, a
}
```
- Calling:
```go
x, y := swap("hello", "world")
fmt.Println(x, y) // Output: world hello
```

---

## **5. Named Return Values** 
```go
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return // Naked return
}
```
- **Naked returns** return named variables.

---

## **6. Variables and Constants** 
```go
var a, b int = 1, 2
```
- Short variable declaration:
```go
x := 42
```
- Constants:
```go
const Pi = 3.14
```

---

## **7. Basic Data Types** 
Common types in Go:
- `bool`, `string`
- `int, float64, uint`, etc.
- `byte` (alias for `uint8`), `rune` (alias for `int32`)
```go
var isTrue bool = true
var number int = 42
var name string = "GoLang"
```

---

## **8. Zero Values** 
- Variables **not explicitly initialized** get a **zero value**:
```go
var i int     // 0
var f float64 // 0.0
var b bool    // false
var s string  // ""
```

---

## **9. Type Conversions** 
- Go **does not** allow implicit type conversions.
- Use **explicit conversion**:
```go
var i int = 42
var f float64 = float64(i) // Convert int to float64
var u uint = uint(f)       // Convert float64 to uint
```

---

## **10. Type Inference** 
- Go **infers** types using `:=`:
```go
i := 42        // int
f := 3.14      // float64
g := "GoLang"  // string
```

---


