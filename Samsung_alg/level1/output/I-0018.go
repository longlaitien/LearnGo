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
	for i := 0; i < 12; i++ {

	}
}
func clockwise(num int, arr []int) int {
	count := 0
	for j := num; j < 12; j++ {
		if arr[j] == 1 {
			return count
		}
		if count == 11 {
			j = 0
		}
		if count == 12 {
			break
		}
		count++
	}
	return -1
}
