package main

    import "fmt"

    func bubbleSort(array []int, size int) {
        for i := 0; i < size-1; i++ {
            for j := 0; j < size-i-1; j++ {
                if array[j] > array[j+1] {
                    array[j], array[j+1] = array[j+1], array[j]
                }
            }
        }
    }

    func main() {
        numbers := []int{5, 2, 8, 12, 3}
        bubbleSort(numbers, len(numbers))
        fmt.Println(numbers)
    }