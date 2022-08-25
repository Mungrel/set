# Set
Simple generic set implementation

## Usage
```go
s := set.New(1, 2, 3)

if s.Contains(1) {
    ...
}

if s.ContainsAll(1, 2, 3) {
    ...
}

data, _ := json.Marshal(s)
fmt.Println(string(data)) // [1,2,3]
```
