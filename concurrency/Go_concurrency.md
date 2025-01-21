# Go Concurrency Summary


---

## **1. Introduction to Concurrency** 
- Go makes **concurrent programming** simple with **goroutines** and **channels**.
- Concurrency in Go is about dealing with many tasks simultaneously, not parallelism (though Go can handle parallelism too).

---

## **2. Goroutines** 
- A **goroutine** is a lightweight thread managed by the Go runtime.
- Use the `go` keyword to start a goroutine.

```go
func say(message string) {
    for i := 0; i < 5; i++ {
        fmt.Println(message)
    }
}

func main() {
    go say("Hello")
    say("World")
}
```
- Output may vary due to concurrency.

---

## **3. Channels** 
- **Channels** are used to communicate between goroutines.
- Use `chan` to declare a channel.

```go
ch := make(chan int)
go func() { ch <- 42 }() // Send value to channel
fmt.Println(<-ch)        // Receive value from channel
```
- Channels are **typed**; they carry values of a specific type.

---

## **4. Buffered Channels** 
- Channels can be **buffered** to store multiple values.
- Create a buffered channel using `make(chan type, capacity)`.

```go
ch := make(chan int, 2)
ch <- 1
ch <- 2
fmt.Println(<-ch) // 1
fmt.Println(<-ch) // 2
```

---

## **5. Channel Direction** 
- Channels can be **directional**, restricting their use to send-only or receive-only.

```go
func send(ch chan<- int, value int) {
    ch <- value // Send value
}

func receive(ch <-chan int) {
    fmt.Println(<-ch) // Receive value
}
```

---

## **6. Select Statement** 
- The **select** statement lets a goroutine wait on multiple communication operations.

```go
ch1 := make(chan int)
ch2 := make(chan int)

go func() { ch1 <- 1 }()
go func() { ch2 <- 2 }()

select {
case msg1 := <-ch1:
    fmt.Println("Received from ch1:", msg1)
case msg2 := <-ch2:
    fmt.Println("Received from ch2:", msg2)
default:
    fmt.Println("No communication")
}
```

---

## **7. Closing Channels** 
- Close a channel to signal no more values will be sent.
- Use `close(channel)`.

```go
ch := make(chan int, 2)
ch <- 1
close(ch)

for value := range ch {
    fmt.Println(value) // Iterates over remaining values
}
```

---

## **8. Deadlock and Best Practices** 
- Sending or receiving from a channel without a corresponding receiver or sender causes a **deadlock**.
- Best practices:
  - Always ensure goroutines are not blocked indefinitely.
  - Avoid sharing memory between goroutines; use channels for communication.

---

## **9. Worker Pool Example** 
- Implementing a worker pool demonstrates how to manage multiple goroutines effectively.

```go
func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job)
        results <- job * 2
    }
}

func main() {
    jobs := make(chan int, 5)
    results := make(chan int, 5)

    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)

    for r := 1; r <= 5; r++ {
        fmt.Println(<-results)
    }
}
```

---

## **10. Advantages of Concurrency in Go** 
- Goroutines are lightweight and efficient.
- Channels simplify communication and synchronization between goroutines.
- The **select** statement allows for responsive and scalable concurrent programs.

---

