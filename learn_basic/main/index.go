package main

// Tìm index trong mảng truyền vào value

func Index[T comparable](arr []T, value T) int {
	for i, v := range arr {
		if v == value {
			return i
		}
	}
	return -1
}

//func main() {
//	si := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//	fmt.Println(Index(si, 3))
//}
