package main

// Viết 1 hàm implement hàm map:  nhân đôi các phần tử trong mảng
// Viết thông thường
func mapX2(arr []int, f func(int) int) []int {
	result := make([]int, len(arr)) // tạo ra 1 mảng mới = mảng truyền vào

	for i, v := range arr {
		result[i] = f(v) // lặp qua từng phần tử và gán nó bằng giá trị trả về của hàm f
	}
	return result
}

// Viết bằng generic

func mapX2Generic[K, V any](arr []K, f func(K) V) []V {
	result := make([]V, len(arr)) // tạo ra 1 mảng mới = mảng truyền vào

	for i, v := range arr {
		result[i] = f(v) // lặp qua từng phần tử và gán nó bằng giá trị trả về của hàm f
	}
	return result
}

//func main() {
//	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//	rs := mapX2(arr, func(x int) int {
//		return x * 2
//	})
//	rsGeneric := mapX2(arr, func(x int) int {
//		return x * 2
//	})
//	fmt.Println(rs)
//	fmt.Println(rsGeneric)
//}
