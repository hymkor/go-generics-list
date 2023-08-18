go-generics-list
================

`go-generics-list` is a thin wrapper around [container/list] using Generics

[container/list]: https://pkg.go.dev/container/list

``` example.go
package main

import (
    "fmt"

    "github.com/hymkor/go-generics-list"
)

func main() {
    L := list.New[rune]()

    for _, c := range "ABCDEFG" {
        L.PushBack(c)
    }

    fmt.Print("Ascend:")
    for p := L.Front(); p.IsSome(); p = p.Next() {
        // Use IsSome() or IsNone() because p can not be nil.
        fmt.Printf(" '%c'", p.Value())
    }
    fmt.Println()

    fmt.Print("Descend:")
    for p := L.Back(); p.IsSome(); p = p.Prev() {
        fmt.Printf(" '%c'", p.Value())
    }
    fmt.Println()
}
```

``` go run example.go |
Ascend: 'A' 'B' 'C' 'D' 'E' 'F' 'G'
Descend: 'G' 'F' 'E' 'D' 'C' 'B' 'A'
```

