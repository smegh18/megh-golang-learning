# Go Methods Summary


---

## **1. Introduction to Methods** 
- In Go, **methods** are functions with a special receiver argument.
- Receivers bind methods to specific types.

```go
type Vertex struct {
    X, Y float64
}

func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```
- The method `Abs` is associated with the `Vertex` type.

---

## **2. Pointer Receivers** 
- Methods can use **pointer receivers** to modify the value of the receiver.
- Pointer receivers are essential for methods that mutate the state of a struct.

```go
func (v *Vertex) Scale(factor float64) {
    v.X = v.X * factor
    v.Y = v.Y * factor
}
```
- Calling:
```go
v := Vertex{3, 4}
v.Scale(2)
fmt.Println(v) // {6, 8}
```

---

## **3. Value vs. Pointer Receivers** 
- Use **value receivers** when the method does not modify the receiver's value.
- Use **pointer receivers** to:
  - Modify the receiver's value.
  - Avoid copying large structs.

---

## **4. Methods and Functions** 
- Methods are functions with a receiver.
- Functions can perform similar tasks without being bound to a type.

```go
func Abs(v Vertex) float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```

---

## **5. Interfaces** 
- An **interface** defines a set of method signatures.
- A type satisfies an interface if it implements all its methods.

```go
type Abser interface {
    Abs() float64
}

func main() {
    var a Abser
    v := Vertex{3, 4}
    a = &v
    fmt.Println(a.Abs()) // 5
}
```

---

## **6. Interface Values** 
- Interface values can hold any type that implements the interface.

```go
type MyFloat float64

func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}

var a Abser
a = MyFloat(-7.5)
fmt.Println(a.Abs()) // 7.5
```

---

## **7. Practical Example: Geometry** 
- Combine structs, methods, and interfaces.

```go
type Shape interface {
    Area() float64
}

type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
```

- Use interface to work with multiple types:
```go
r := Rectangle{3, 4}
var s Shape = r
fmt.Println(s.Area()) // 12
```

---
