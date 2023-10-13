### Simple linked list in go

#### How to install
``` go get github.com/sgatu/go-simple-linked-list@v1.0.6 ```

#### How to use
```go
  import linkedlist "github.com/sgatu/go-simple-linked-list"
  func main() {
    list := linkedlist.LinkedList[int]{}

    list.Push(1)
    list.Push(2)
    list.Push(3)
    for list.Len() > 0 {
      if elem, err := list.Pop(); err == nil {
        fmt.Printf("Got element %d\n", elem)
      }
    }
  }
```
