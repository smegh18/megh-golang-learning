# Go Generics Summary


---

## **1. Introduction to Generics** 
- Generics enable writing functions and data structures that work with any data type.
- They use **type parameters** to define type flexibility.

```go
func Print[T any](v T) {
    fmt.Println(v)
}
```
- `T` is the **type parameter**, and `any` is a constraint that allows any type.

---

## **2. Type Parameters** 
- A **type parameter** is defined in square brackets `[]` after a function or type name.
- It allows the function or type to operate on different types.

```go
func Sum[T int | float64](a, b T) T {
    return a + b
}
```

---

## **3. Constraints** 
- **Constraints** restrict the types that can be used with a generic function or type.
- Example constraints:
  - `any`: Accepts all types.
  - Specific types: e.g., `int`, `float64`.
  - Interfaces: Define a set of methods the type must implement.

```go
type Number interface {
    int | float64
}

func Add[T Number](a, b T) T {
    return a + b
}
```

---

## **4. Using Generics with Structs** 
- Generics can be applied to structs for flexibility in field types.

```go
type Pair[T any] struct {
    First, Second T
}

p := Pair[int]{First: 10, Second: 20}
fmt.Println(p) // {10 20}
```

---

## **5. Generic Slices and Maps** 
- Generics simplify working with collections like slices and maps.

```go
func Filter[T any](s []T, fn func(T) bool) []T {
    var result []T
    for _, v := range s {
        if fn(v) {
            result = append(result, v)
        }
    }
    return result
}
```
- Example usage:
```go
nums := []int{1, 2, 3, 4}
even := Filter(nums, func(n int) bool { return n%2 == 0 })
fmt.Println(even) // [2 4]
```

---

## **6. Type Inference** 
- Go can often infer type parameters automatically.
- Explicit type arguments are optional when the compiler can deduce them.

```go
fmt.Println(Sum(3, 5)) // Type inference: T is int
```

---

## **7. Generics with Interfaces** 
- Combine interfaces and generics to define reusable types and methods.

```go
type Comparable interface {
    CompareTo(other Comparable) int
}

func Max[T Comparable](a, b T) T {
    if a.CompareTo(b) > 0 {
        return a
    }
    return b
}
```

---

## **8. Advantages of Generics** 
- Reduce code duplication by creating reusable components.
- Increase flexibility while maintaining type safety.

---

