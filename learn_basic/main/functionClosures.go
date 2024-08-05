package main

// Hàm tạo bộ lọc với tiêu chí cụ thể
func createFilter(minValue int) func(int) bool {
	return func(n int) bool {
		return n > minValue
	}
}

//func main() {
//	// Tạo một bộ lọc để giữ lại các số lớn hơn 10
//	filter := createFilter(10)
//
//	numbers := []int{5, 12, 8, 20, 15, 3}
//	filteredNumbers := []int{}
//
//	// Sử dụng bộ lọc để lọc các số
//	for _, number := range numbers {
//		if filter(number) {
//			filteredNumbers = append(filteredNumbers, number)
//		}
//	}
//
//	fmt.Println("Filtered Numbers:", filteredNumbers)
//}
