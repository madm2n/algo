# algo

Simple Go implementations of classic algorithms and data structures.

## Contents

### Sort

- Bubble sort ([`sort/bubble.go`](sort/bubble.go))
- Insertion sort ([`sort/insertion.go`](sort/insertion.go))
- Heap sort ([`sort/heapsort.go`](sort/heapsort.go))
- Merge sort ([`sort/merge.go`](sort/merge.go))
- Selection sort ([`sort/selection.go`](sort/selection.go))

| Algorithm      | Best       | Average    | Worst      | Space |
| -------------- | ---------- | ---------- | ---------- | ----- |
| Bubble sort    | O(n)       | O(n²)      | O(n²)      | O(1)  |
| Insertion sort | O(n)       | O(n²)      | O(n²)      | O(1)  |
| Heap sort      | O(n log n) | O(n log n) | O(n log n) | O(1)  |
| Merge sort     | O(n log n) | O(n log n) | O(n log n) | O(n)  |
| Selection sort | O(n²)      | O(n²)      | O(n²)      | O(1)  |

### Data Structures

- Max-priority queue ([`ds/priority_queue.go`](ds/priority_queue.go))

### Search

- Binary search ([`search/binary.go`](search/binary.go))

### Recurrences

- Naive matrix multiply ([`recurrences/matrix_multiply_naive.go`](recurrences/matrix_multiply_naive.go))

### Exercises

- HasIntSum ([`exercises/has_int_sum.go`](exercises/has_int_sum.go))

## Usage

```go
package main

import (
    "fmt"

    "algo/ds"
    "algo/search"
    "algo/sort"
)

func main() {
    data := []int{5, 2, 9, 1, 5, 6}

    sort.MergeSort(data)
    fmt.Println(data)

    idx := search.BinarySearch(data, 5)
    fmt.Println(idx)

    queue := ds.NewPriorityQueue[int]()
    queue.Insert(5)
    queue.Insert(2)
    queue.Insert(9)
    max, _ := queue.ExtractMax()
    fmt.Println(max)
}
```

## Tests

```sh
go test ./...
```
