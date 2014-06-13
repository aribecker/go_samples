// We can use channels to synchronize execution
// across goroutines. Here's an example of using a
// blocking receive to wait for a goroutine to finish.

package main

import "fmt"
import "time"

// This is the function we'll run in a goroutine. The
// `done` channel will be used to notify another
// goroutine that this function's work is done.
func worker1(done chan bool) {
    fmt.Print("working1...")
    time.Sleep(time.Second)
    fmt.Println("done1")

    // Send a value to notify that we're done.
    done <- true
}
func worker2(done chan bool) {
    fmt.Print("working2...")
    time.Sleep(time.Second*10)
    fmt.Println("done2")

    // Send a value to notify that we're done.
    done <- true
}
func worker3(done chan bool) {
    fmt.Print("working3...")
    time.Sleep(time.Second)
    fmt.Println("done3")

    // Send a value to notify that we're done.
    done <- true
}

func main() {

    // Start a worker goroutine, giving it the channel to
    // notify on.
    done1 := make(chan bool, 1)
    done2 := make(chan bool, 1)
    done3 := make(chan bool, 1)
    go worker1(done1)
    go worker2(done2)
    go worker3(done3)

    // Block until we receive a notification from the
    // worker on the channel.
    <-done1
    <-done2
    <-done3
}
