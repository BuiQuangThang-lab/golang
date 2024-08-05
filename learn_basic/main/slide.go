package main

//func main() {
//primes := [6]int{2, 3, 5, 7, 11, 13}
//a := primes[0:2]
//b := primes[1:3]
//fmt.Println(a, b)
//
//b[0] = 10000
//
//fmt.Println(a, b)
//fmt.Println(primes)

// slide không tạo ra 1 vùng nhớ mới mà nó vẫn thuộc vùng nhớ của thằng cha => khi thay đổi giá trị của phần tử nào từ slide thì tất cả bị thay đổi theo

// length: sức chứa thực sự của 1 slide
// capacity: sức chứa tối đa của 1 slide có thể chứa

// bình thường thêm 1 thằng vào arr: copy arr mới = độ dài arr cũ + 1 rồi đẩy phần tử qua và add thằng mới vào; capacity giúp thêm phần tử mà k cần copy

//slide := make([]int, 0, 5) // length = 0 cap = 5
////fmt.Println(slide)
//
//slide = append(slide, 1, 5, 9, 12, 23)

//fmt.Println(slide)

// for của slide

//for i := 0; i < len(slide); i++ {
//	fmt.Println(slide[i])
//}

//for i, v := range slide {
//	fmt.Println(i, v)
//}
//
//for _, v := range slide {
//	fmt.Println(v)
//}
//
//for i := range slide {
//	fmt.Println(i)
//}

//}
