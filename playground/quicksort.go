package main

    import "fmt"

    func quickSort(arr []int, low, high int) []int {
        if low < high {
            arr, p := partition(arr, low, high)
            arr = quickSort(arr, low, p-1)
            arr = quickSort(arr, p+1, high)
        }
        return arr
    }

    func partition(arr []int, low, high int) ([]int, int) {
        pivot := arr[high]
        i := low
        for j := low; j < high; j++ {
            if arr[j] < pivot {
                arr[i], arr[j] = arr[j], arr[i]
                i++
            }
        }
        arr[i], arr[high] = arr[high], arr[i]
        return arr, i
    }

    func main() {
        numbers := []int{5, 2, 8, 12, 3}
        numbers = quickSort(numbers, 0, len(numbers)-1)
        fmt.Println(numbers)
    }