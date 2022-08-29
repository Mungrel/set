# Set [![CI](https://github.com/Mungrel/set/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/Mungrel/set/actions/workflows/ci.yml)
Simple generic set implementation

## Usage
```go
// Empty set
s := set.New[string]()

s.Add("apple", "banana", "orange")

// Initialise with elements
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
