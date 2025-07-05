package main

    import "fmt"

    func mergeSort(array []int) []int {
        length := len(array)
        if length <= 1 {
            return array
        }
        mid := length / 2
        left := mergeSort(array[:mid])
        right := mergeSort(array[mid:])
        return merge(left, right)
    }

    func merge(left []int, right []int) []int {
        result := make([]int, 0, len(left)+len(right))
        i, j := 0, 0
        for i < len(left) && j < len(right) {
            if left[i] <= right[j] {
                result = append(result, left[i])
                i++
            } else {
                result = append(result, right[j])
                j++
            }
        }
        result = append(result, left[i:]...)
        result = append(result, right[j:]...)
        return result
    }

    func main() {
        numbers := []int{5, 2, 8, 12, 3}
        sortedNumbers := mergeSort(numbers)
        fmt.Println(sortedNumbers)
    }