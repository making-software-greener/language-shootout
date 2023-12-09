package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func bubbleSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    
    fmt.Println("Enter the number of elements:")
    scanner.Scan()
    n, err := strconv.Atoi(scanner.Text())
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error reading the size: %v\n", err)
        return
    }

    arr := make([]int, n)
    
    fmt.Println("Enter the elements (one per line):")
    for i := 0; i < n; i++ {
        scanner.Scan()
        arr[i], err = strconv.Atoi(scanner.Text())
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error reading an element: %v\n", err)
            return
        }
    }

    bubbleSort(arr)

    fmt.Println("Sorted array:")
    for _, v := range arr {
        fmt.Printf("%d\n", v)
    }
}

