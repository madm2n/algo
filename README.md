# algo

Simple Go implementations of classic algorithms and data structures.

## Contents

### Sort

- Insertion sort (`sort/InsertionSort`)
- Merge sort (`sort/MergeSort`)
- Selection sort (`sort/SelectionSort`)

### Search

- Binary search (`search/BinarySearch`)

## Usage

```go
package main

import (
    "fmt"

    "algo/search"
    "algo/sort"
)

func main() {
    data := []int{5, 2, 9, 1, 5, 6}

    sort.MergeSort(data)
    fmt.Println(data)

    idx := search.BinarySearch(data, 5)
    fmt.Println(idx)
}
```

## Tests

```sh
go test ./...
```
