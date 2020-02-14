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

## License

MIT

## Author

Ryo Kitagawa
