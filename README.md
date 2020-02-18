# go-todotxt

Simple todo.txt library.

## Usage

```go
import "github.com/kitagry/go-todotxt"

f, err := os.Open("todo.txt")
r := todotxt.NewReader(f)
tasks, err := r.ReadAll()

for _, task := range tasks {
  task.SetDescription(task.Description() + " +Project")
}

f, err := os.Write("todo.txt")
w := todotxt.NewWriter(f)
w.WriteAll(tasks)
```

## Tools

### Install

```
$ cd cmd/todotxt
$ go install
```

![go-todotxt](https://user-images.githubusercontent.com/21323222/74706397-3cdef380-525a-11ea-9877-458ae3b6cebd.gif)

## License

MIT

## Author

Ryo Kitagawa

