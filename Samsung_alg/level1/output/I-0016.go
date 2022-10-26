package level1

import "fmt"

// Tô màu:
//cho 1 ma trận NxM (5<=N, M<=50)
//trong mỗi ô sẽ được đánh 1 số, cần tô màu các ô thỏa mãn điều kiện sau:
//đk: bằng cách cộng hoặc trừ các số ở top, down,left, right, bạn có thể tính ra chữ số ở ô hiện tại
//in ra số ô được tô màu

func i0016() {

}
func set() {
	var n, m int32
	fmt.Println("N: ")
	_, err := fmt.Scanln(&n)
	if err != nil {
		return
	}
	fmt.Println("M: ")
	_, err = fmt.Scanln(&m)
	if err != nil {
		return
	}
	var array make([n][m]int32)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {

		}
	}
}
