package main

import "fmt"

func main() {
    var m, n int
    fmt.Scan(&m, &n)

    a := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&a[i])
    }

    queries := make([][2]int, m)
    for i := 0; i < m; i++ {
        fmt.Scan(&queries[i][0], &queries[i][1])
    }

    fmt.Println(m, n)

    for i := 0; i < n; i++ {
        if i > 0 {
            fmt.Print(" ")
        }
        fmt.Print(a[i])
    }
    fmt.Println()

    for i := 0; i < m; i++ {
        fmt.Println(queries[i][0], queries[i][1])
    }
}