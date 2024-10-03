# Golang begining

## 1. Concurrency: Goroutine & Channel
- the `go` keyword mark that function can call in new routine (multiple thread)
> let take a look at example below:

In Go, **goroutines** and **channels** are key concepts that facilitate concurrency. Let’s explore both.

#### 1. **Goroutines**
A **goroutine** is a lightweight thread managed by the Go runtime. You can think of it as a function that runs concurrently with other goroutines. They are much more efficient than system threads, allowing thousands of goroutines to run at the same time without overwhelming system resources.

To start a goroutine, you use the `go` keyword:

```go
package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello from goroutine")
}

func main() {
	go sayHello() // This runs in a separate goroutine
	fmt.Println("Hello from main")
	time.Sleep(1 * time.Second) // Give the goroutine time to finish
}
```

In this example:
- `go sayHello()` starts the `sayHello` function as a new goroutine.
- The main function runs in the main goroutine, and the `sayHello` function runs concurrently in a separate goroutine.

Goroutines don’t automatically wait for each other to finish, which is why `time.Sleep` is added to allow the `sayHello` function to complete. However, this is not the ideal way to manage goroutines. This is where **channels** come in.

#### 2. **Channels**
**Channels** are a Go construct that allow goroutines to communicate and synchronize by sending and receiving values. You can think of a channel as a pipe through which you can send data between goroutines.

##### Declaring a channel:
```go
ch := make(chan int)
```

- `chan int` declares a channel that transports `int` values.

##### Sending and receiving from a channel:
```go
package main

import "fmt"

func main() {
	ch := make(chan int)

	// Goroutine that sends data into the channel
	go func() {
		ch <- 42 // Send 42 to the channel
	}()

	// Receiving data from the channel
	val := <-ch
	fmt.Println(val) // Prints: 42
}
```

- `ch <- 42` sends the value `42` into the channel.
- `val := <-ch` receives the value from the channel.

This allows the main function and the anonymous goroutine to communicate through the channel.

#### 3. **Buffered vs. Unbuffered Channels**
Channels can be **buffered** or **unbuffered**:
- **Unbuffered** channels (default) block until the sender and receiver are both ready.
- **Buffered** channels allow for a certain number of values to be sent before blocking occurs.

Example of a **buffered channel**:
```go
ch := make(chan int, 2) // Channel with a buffer of 2
```

This allows up to two values to be sent before the sender is blocked.

#### 4. **Channel Closing**
Channels can be **closed** when no more data will be sent. Receivers can check if a channel is closed:

```go
package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 42
		close(ch)
	}()

	val, ok := <-ch
	if ok {
		fmt.Println("Received:", val)
	} else {
		fmt.Println("Channel closed")
	}
}
```

- `close(ch)` closes the channel.
- `val, ok := <-ch` returns `ok` as `false` if the channel is closed.

#### 5. **Example: Goroutines with Channels**
Here’s an example of a program using goroutines and channels to calculate the sum of two slices concurrently:

```go
package main

import "fmt"

// Function to sum slice elements
func sum(nums []int, ch chan int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	ch <- total // Send total to channel
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}

	// Create a channel
	ch := make(chan int)

	// Split the work into two goroutines
	go sum(numbers[:len(numbers)/2], ch) // First half
	go sum(numbers[len(numbers)/2:], ch) // Second half

	// Receive results from both goroutines
	result1, result2 := <-ch, <-ch

	// Sum up the results
	fmt.Println("Total:", result1+result2)
}
```

In this example:
- Two goroutines are created to sum halves of the slice.
- Each sends its result through the channel.
- The main function waits for both results and sums them.

#### Conclusion
- **Goroutines** provide a way to run functions concurrently.
- **Channels** enable communication between goroutines, allowing you to safely pass data.
- Together, they form Go’s powerful concurrency model.

Would you like to explore any specific aspect further?

**How about synchronous all goroutine**

Trong Go, mặc định thì goroutines hoạt động theo kiểu bất đồng bộ (asynchronous). Để thực hiện đồng bộ (synchronous) giữa các goroutines, bạn có thể sử dụng các kỹ thuật sau đây:

### 1. **Sử dụng `sync.WaitGroup`**
`sync.WaitGroup` cho phép bạn chờ cho một nhóm goroutines hoàn thành trước khi tiếp tục thực hiện chương trình.

#### Ví dụ:
```go
package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Đánh dấu rằng goroutine đã hoàn thành
	fmt.Printf("Worker %d đang làm việc...\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // Thêm vào WaitGroup
		go worker(i, &wg)
	}

	wg.Wait() // Chờ tất cả goroutines hoàn thành
	fmt.Println("Tất cả công việc đã hoàn thành.")
}
```
Trong ví dụ trên:
- `wg.Add(1)` thêm một goroutine vào WaitGroup.
- `defer wg.Done()` đánh dấu rằng goroutine đã hoàn thành.
- `wg.Wait()` chờ tất cả goroutines trong WaitGroup hoàn thành.

### 2. **Sử dụng Channels**
Channels cũng có thể được sử dụng để đồng bộ hóa giữa các goroutines.

#### Ví dụ:
```go
package main

import (
	"fmt"
)

func worker(id int, done chan bool) {
	fmt.Printf("Worker %d đang làm việc...\n", id)
	done <- true // Gửi tín hiệu hoàn thành
}

func main() {
	done := make(chan bool)

	for i := 1; i <= 3; i++ {
		go worker(i, done)
	}

	for i := 1; i <= 3; i++ {
		<-done // Chờ tín hiệu từ goroutines
	}

	fmt.Println("Tất cả công việc đã hoàn thành.")
}
```
Trong ví dụ này:
- Mỗi goroutine gửi một tín hiệu vào channel `done` khi nó hoàn thành công việc.
- Vòng lặp `for` trong hàm `main` chờ nhận tín hiệu từ tất cả các goroutines.

### 3. **Kết hợp cả hai**
Bạn cũng có thể kết hợp cả `sync.WaitGroup` và channels để có thêm tính linh hoạt trong việc đồng bộ hóa.

### Kết luận
Để kiểm soát sự đồng bộ hóa trong Go, bạn có thể sử dụng `sync.WaitGroup` hoặc channels, tùy thuộc vào nhu cầu cụ thể của bạn. `sync.WaitGroup` thường là phương pháp đơn giản và rõ ràng hơn cho nhiều tình huống. Hãy nhớ rằng việc đồng bộ hóa giúp đảm bảo rằng tất cả công việc đã hoàn thành trước khi chương trình kết thúc, đặc biệt là khi có nhiều goroutines chạy song song.

Nếu bạn cần thêm thông tin chi tiết về cách sử dụng goroutines và đồng bộ hóa trong Go, bạn có thể tham khảo tài liệu chính thức của Go hoặc các hướng dẫn trực tuyến.

