package ouput

import "fmt"

// 1 chiếc két sắt với 3 vòng số
//mỗi vòng số có 12 chữ số được đánh số 0 hoặc 1
//mỗi vòng số có thể được xoay theo chiều kim đồng hồ hoặc ngược chiều kim đồng hồ
//để mở két sắt, 3 vòng số phải xếp thẳng hàng số 1 bằng cách xoay theo chiều kim đồng hồ hoặc ngược lại
//với số lần quay ít nhất

func I0018() {
	d1 := []int{1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0}
	d2 := []int{0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0}
	d3 := []int{0, 1, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0}
	var arr [12]int
	min := 10
	for h := 0; h < 12; h++ {
		value := turn(h, d1) + turn(h, d2) + turn(h, d3)
		arr[h] = value
		if value < min {
			min = value
		}
	}
	fmt.Println(min)
}

func turn(idx int, arr []int) int {
	count := 0
	i := idx
	j := idx
	for count < 12 {
		if arr[j] == 1 || arr[i] == 1 {
			return count
		}
		if i == 11 {
			i = 0
		} else {
			i++
		}
		if j == 0 {
			j = 11
		} else {
			j--
		}
		if count == 12 {
			break
		}
		count++
	}
	return -1
}
