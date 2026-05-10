# algo

Simple Go implementations of classic algorithms and data structures.

## Contents

- Insertion sort (`sort/InsertionSort`)
- Merge sort (`sort/MergeSort`)
- Selection sort (`sort/SelectionSort`)

## Usage

```go
package main

import (
    "fmt"

    "algo/sort"
)

func main() {
    data := []int{5, 2, 9, 1, 5, 6}

    sort.MergeSort(data)
    fmt.Println(data)
}
```

## Tests

```sh
go test ./...
```
